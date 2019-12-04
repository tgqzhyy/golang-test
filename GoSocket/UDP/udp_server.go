package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6666")

	if err != nil {
		fmt.Println("resolveudpaddr err:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("listen err", err)
		return
	}

	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, udpaddr, err := conn.ReadFromUDP(buf)
		if n == 0 {
			fmt.Println("用户%s退出\n", conn.RemoteAddr().String())
			break
		}
		if err != nil {
			fmt.Println("read err", err)
			return
		}

		_, err = conn.WriteToUDP((buf[:n]), udpaddr)
		if err != nil {
			fmt.Println("writeudp err:", err)
			return
		}
	}

}
