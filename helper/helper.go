package helper

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Erro ao obter endereços:", err)
		return ""
	}
	var ip string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ip = ipNet.IP.String()
			break
		}
	}
	return ip
}
func GetHost() string {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("Erro ao obter o nome do host:", err)
		return ""
	}
	return host
}

func GetOperationSystem() string {
	return runtime.GOOS
}
