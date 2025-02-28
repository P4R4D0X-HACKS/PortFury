package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"portfury/internal/scanner"
	"portfury/internal/utils"
)

func showHelp() {
	fmt.Println("PortFury - A fast, concurrent port scanner")
	fmt.Println("Usage: portfury [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  -host string    Target host (default: scanme.nmap.org)")
	fmt.Println("  -ports string   Port range or list (e.g., 1-1000, 80,443) (default: 1-1024)")
	fmt.Println("  -timeout int    Timeout per scan in seconds (default: 2)")
	fmt.Println("  -workers int    Number of concurrent workers (default: 100)")
	fmt.Println("  -json           Output results in JSON format")
	fmt.Println("  -h, --help      Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  portfury -host=example.com -ports=80,443")
	fmt.Println("  portfury -host=scanme.nmap.org -ports=1-1000 -workers=200 -json")
}

func main() {
	host := flag.String("host", "scanme.nmap.org", "Target host")
	portRange := flag.String("ports", "1-1024", "Port range (e.g., 1-100, 80,443)")
	timeout := flag.Int("timeout", 2, "Scan timeout in seconds")
	workers := flag.Int("workers", 100, "Number of concurrent workers")
	jsonOutput := flag.Bool("json", false, "Output results in JSON format")
	help := flag.Bool("h", false, "Show help message")
	flag.BoolVar(help, "help", false, "Show help message")
	
	flag.Parse()

	// Show help if -h or --help is used, or if no arguments are provided
	if *help || len(os.Args) == 1 {
		showHelp()
		return
	}

	// Validate host
    if !utils.IsValidHost(*host) {
        log.Fatalf("Invalid host: %s (must be a valid IP address or domain name)", *host)
    }
	
	ports, err := utils.ParsePortRange(*portRange)
	if err != nil {
		log.Fatalf("Invalid port range: %v", err)
	}

	results := scanner.ScanPorts(*host, ports, *timeout, *workers)
	if *jsonOutput {
		fmt.Println(results.ToJSON())
	} else {
		fmt.Println(results.ToString())
	}
}
