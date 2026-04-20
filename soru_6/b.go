package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type DHCPRecord struct {
	Time     string
	ClientID string
	ReqIP    string
	VendorID string
	Hostname string
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	file, err := os.Create("ayristirilmis_dhcp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	url := "https://raw.githubusercontent.com/lowgame/rakort_stajyer_dhcp_data/refs/heads/main/dhcp.txt"
	cmd := exec.CommandContext(ctx, "curl", "-s", url)
	
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DHCP akisi dinleniyor (300s)...")

	scanner := bufio.NewScanner(stdout)
	var currentBlock []string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "---") {
			processBlock(currentBlock, file)
			currentBlock = nil
		} else {
			currentBlock = append(currentBlock, line)
		}
	}

	if len(currentBlock) > 0 {
		processBlock(currentBlock, file)
	}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Sure doldu, islem tamamlandi.")
	}
}

func processBlock(lines []string, file *os.File) {
	isRequest := false
	record := DHCPRecord{
		Time:     "N/A",
		ClientID: "N/A",
		ReqIP:    "N/A",
		VendorID: "N/A",
		Hostname: "N/A",
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "TIME:") {
			record.Time = strings.TrimPrefix(line, "TIME: ")
		}

		if strings.Contains(line, "OPTION: 53") && strings.Contains(line, "type 3") {
			isRequest = true
		}

		if strings.Contains(line, "OPTION: 61") {
			record.ClientID = extractValue(line, "Client-identifier")
		}
		if strings.Contains(line, "OPTION: 50") {
			record.ReqIP = extractValue(line, "Request IP address")
		}
		if strings.Contains(line, "OPTION: 60") {
			record.VendorID = extractValue(line, "Vendor class identifier")
		}
		if strings.Contains(line, "OPTION: 12") {
			record.Hostname = extractValue(line, "Host name")
		}
	}

	if isRequest {
		output := fmt.Sprintf("[%s] IP: %-15s | Host: %-15s | ClientID: %-20s | Vendor: %s\n",
			record.Time, record.ReqIP, record.Hostname, record.ClientID, record.VendorID)
		
		fmt.Print(output)
		file.WriteString(output)
	}
}

func extractValue(line, label string) string {
	parts := strings.Split(line, label)
	if len(parts) > 1 {
		val := strings.TrimSpace(parts[1])
		val = strings.TrimPrefix(val, ":")
		return strings.TrimSpace(val)
	}
	return "N/A"
}
