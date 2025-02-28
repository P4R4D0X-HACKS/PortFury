# PortFury 🚀  
**A fast, concurrent port scanner written in Go.**

---

## Description  
PortFury is a lightweight and efficient port scanner designed for offensive security professionals and enthusiasts. It uses **goroutines** for concurrent scanning and supports **banner grabbing** to identify services running on open ports.  

---

## Features  
- **Concurrent Scanning**: Scan thousands of ports in seconds using worker pools.  
- **Banner Grabbing**: Identify services (e.g., SSH, HTTP) running on open ports.  
- **Customizable**: Specify target hosts, port ranges, timeouts, and more.  
- **JSON Output**: Export results in JSON format for integration with other tools.  

---

## Installation  

### Prerequisites  
- Go 1.20 or higher.  

### Steps  
1. Clone the repository:  
   ```bash
   git clone https://github.com/P4R4D0X-HACKS/PortFury
   cd portfury
   ```
2. Build the project:
    ```bash
    make build
    ```
    Or Manually:
    ```bash
    go build -o bin/portfury ./cmd/portfury
    ```

## Usage

### Basic Scan

**Scan a target with default settings (ports 1-1024, 100 workers, 2s timeout)**

```bash
./bin/portfury -host=scanme.nmap.org
```
### Custom Port Range

Scan specific ports (e.g., 80, 443, 8080):
```bash
./bin/portfury -host=example.com -ports=80,443,8080
```

### Advanced Scan

Scan a large range of ports with 200 workers and a 5s timeout:
```bash
./bin/portfury -host=example.com -ports=1-10000 -workers=200 -timeout=5
```

### JSON Output

Export results in JSON format:
```bash
./bin/portfury -host=example.com -ports=1-1000 -json
```


| Flag | Description | Default Value |
| --- | --- | --- |
| `-host` | Target host (e.g., `example.com`) | `scanme.nmap.org` |
| `-ports` | Port range or list (e.g., `1-1000`) | `1-1024` |
| `-timeout` | Timeout per scan in seconds | `2` |
| `-workers` | Number of concurrent workers | `100` |
| `-json` | Output results in JSON format | `false` |

Contributing
------------

Contributions are welcome! Here's how you can help:

1.  Fork the repository.

2.  Create a new branch (`git checkout -b feature/your-feature`).

3.  Commit your changes (`git commit -m 'Add some feature'`).

4.  Push to the branch (`git push origin feature/your-feature`).

5.  Open a pull request.

Acknowledgments
---------------

-   Inspired by tools like `nmap` and `masscan`.

-   Built with ❤️ using Go.

> NOTE : YOU CAN ALSO YOU IP ADDRESS IN HOST.

---

> This project is still under development and below is a consise roadmap of the features. You can find an extensive PortFury development sheet in the file called ToDo.txt . If you really like this project feel free to support my journey of developing this tool by contributing to it. 

## Roadmap  🚀

### Phase 1: Core Features
- [ ] Implement SYN Scan (Stealth Scan).
- [ ] Add UDP Scan support.
- [ ] Add Rate Limiting.

### Phase 2: Advanced Scanning
- [ ] Implement OS Fingerprinting.
- [ ] Add Distributed Scanning.
- [ ] Add Asynchronous DNS Resolution.

### Phase 3: Usability Enhancements
- [ ] Add Interactive Mode.
- [ ] Add Configuration File Support.
- [ ] Add Progress Bar.

### Phase 4: Advanced Output and Reporting
- [ ] Generate HTML Reports.
- [ ] Add Tool Integration (CSV, XML, JSON).
- [ ] Add Real-Time Alerts (Email, Slack).

### Phase 5: Security and Evasion
- [ ] Add Evasion Techniques (Packet Fragmentation, Random Delays).
- [ ] Add Proxy Support (SOCKS, HTTP).
- [ ] Add Encrypted Communication (TLS, AES).

### Phase 6: AI/ML Integration
- [ ] Add Vulnerability Prediction using Machine Learning.
- [ ] Add Anomaly Detection.

### Phase 7: Community and Collaboration
- [ ] Add Plugin System.
- [ ] Add API for Automation (REST, gRPC).

### Phase 8: Documentation and Support
- [ ] Write Comprehensive Documentation.
- [ ] Create Community Forum or Discord Server.
