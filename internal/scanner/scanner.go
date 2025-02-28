package scanner

import (
	"fmt"
	"net"
	"portfury/internal/models"
	"portfury/pkg/banner"
	"sync"
	"time"
)

func ScanPort(host string, port int, timeout int) models.ScanResult {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
	if err != nil {
		return models.ScanResult{
			Port:   port,
			Status: "closed",
		}
	}
	defer conn.Close()
	return models.ScanResult{
		Port:   port,
		Status: "open",
		Banner: banner.GetBanner(conn),
	}
}

func ScanPorts(host string, ports []int, timeout int, workers int) models.ScanResults {
	var wg sync.WaitGroup
	results  := make(chan models.ScanResult, len(ports))
	portsChan := make(chan int , workers)

	// Start workers
	for i := 0; i < workers; i++ {
		go func ()  {
			for ports := range portsChan {
				results <- ScanPort(host, ports, timeout)
				wg.Done()
			}
		}()
	}

	// Send port to workers

	for _, port := range ports {
		wg.Add(1)
		portsChan <- port
	}

	close(portsChan)
	wg.Wait()
	close(results)

	// Collect results

	var scanResults models.ScanResults
	for result := range results {
		scanResults = append(scanResults, result)
	}
	return scanResults
}
