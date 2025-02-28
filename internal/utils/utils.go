package utils

import (
	"net"
	"strconv"
	"strings"
)

func ParsePortRange(portRange string) ([]int, error) {
	// Parse the port range
	var ports []int
	ranges := strings.Split(portRange, ",")
	for _, r := range ranges {
		if strings.Contains(r, "-") {
			startEnd := strings.Split(r, "-")
			start, err := strconv.Atoi(startEnd[0])
			if err != nil {
				return nil, err
			}
			end, err := strconv.Atoi(startEnd[1])
			if err != nil {
				return nil, err
			}
			for i := start; i <= end; i++ {
				ports = append(ports, i)
			}
		} else {
			port, err := strconv.Atoi(r)
			if err != nil {
				return nil, err
			}
			ports = append(ports, port)
		}
	}
	return ports, nil
}

// IsValidHost checks if the input is a valid IP address or domain name.
func IsValidHost(host string) bool {
	// Check if it's a valid IP address
	if ip := net.ParseIP(host); ip != nil {
		return true
	}

	// Check if it's a valid domain name
	if _, err := net.LookupHost(host); err == nil {
		return true
	}

	return false
}
