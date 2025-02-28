package banner

import (
	"bufio"
	"net"
	"time"
)

func GetBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	reader := bufio.NewReader(conn)
	banner, _ := reader.ReadString('\n')
	return banner
}
