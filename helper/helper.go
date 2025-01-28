package helper

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
)

var Layout_date string = "2006-01-02 15:04:05"
var Sysdate time.Time = GetSysdate()

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

func ConvertDate(dateStr string, layout string) (time.Time, error) {
	var date time.Time
	convert, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Erro ao converter a string para time.Time:", err)
		return date, err
	}
	return convert, nil
}

func GetSysdate() time.Time {
	return time.Now()
}

func AddDays(date *time.Time, qtyDays int) time.Time {
	return date.Add(time.Duration(qtyDays) * 24 * time.Hour)
}
