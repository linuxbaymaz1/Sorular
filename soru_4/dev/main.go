package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Result struct {
	IP   string
	Live bool
}

func ping(ip string) bool {
	return exec.Command("ping", "-c", "1", "-W", "1", ip).Run() == nil
}

func getIPs(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	if len(ips) > 2 {
		return ips[1 : len(ips)-1], nil
	}
	return ips, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func getDefaultSubnet() string {
	out, err := exec.Command("docker", "network", "inspect", "bridge").Output()
	if err != nil {
		return "172.17.0.0/16"
	}

	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, `"Subnet"`) {
			parts := strings.Split(line, `"`)
			if len(parts) >= 4 {
				return parts[3]
			}
		}
	}
	return "172.17.0.0/16"
}

func main() {
	subnet := flag.String("subnet", "", "CIDR network")
	workers := flag.Int("workers", 512, "Concurrency")
	outFile := flag.String("out", "down_hosts.log", "Log file")
	interval := flag.Duration("interval", 10*time.Second, "Scan interval")
	flag.Parse()

	target := *subnet
	if target == "" {
		target = getDefaultSubnet()
	}

	ips, err := getIPs(target)
	if err != nil {
		log.Fatalf("Subnet error: %v\n", err)
	}

	fmt.Printf("🚀 Scan started on %s | IPs: %d | Workers: %d\n\n", target, len(ips), *workers)

	cycle := 1

	for {
		start := time.Now()
		
		jobs := make(chan string, len(ips))
		results := make(chan Result, len(ips))
		var wg sync.WaitGroup

		for i := 0; i < *workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for ip := range jobs {
					results <- Result{IP: ip, Live: ping(ip)}
				}
			}()
		}

		go func() {
			for _, ip := range ips {
				jobs <- ip
			}
			close(jobs)
		}()

		go func() {
			wg.Wait()
			close(results)
		}()

		var down []string
		active := 0

		for res := range results {
			if res.Live {
				active++
			} else {
				down = append(down, res.IP)
			}
		}

		file, err := os.OpenFile(*outFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			writer := bufio.NewWriter(file)
			fmt.Fprintf(writer, "\n=== CYCLE %d | %s ===\n", cycle, time.Now().Format("2006-01-02 15:04:05"))
			for _, ip := range down {
				fmt.Fprintf(writer, "[DOWN] %s\n", ip)
			}
			writer.Flush()
			file.Close()
		}

		fmt.Printf("🔄 [Cycle %d] Done in %v | Active: %d | Down: %d\n", cycle, time.Since(start).Round(time.Millisecond), active, len(down))
		
		if len(down) > 0 {
			limit := 5
			if len(down) < 5 {
				limit = len(down)
			}
			for i := 0; i < limit; i++ {
				fmt.Printf("      - %s\n", down[i])
			}
			if len(down) > limit {
				fmt.Printf("      ... +%d more\n", len(down)-limit)
			}
		}
		fmt.Println(strings.Repeat("-", 40))

		cycle++
		time.Sleep(*interval)
	}
}
