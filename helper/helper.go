package helper

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
)

type Helper struct {
}

func NewHelper() *Helper {
	return &Helper{}
}

var Layout_date string
var Sysdate time.Time

func Init() {
	Sysdate = time.Now()
	Layout_date = "2006-01-02 15:04:05"
}

func (h *Helper) GetIp() string {
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
func (h *Helper) GetHost() string {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("Erro ao obter o nome do host:", err)
		return ""
	}
	return host
}

func (h *Helper) GetOperationSystem() string {
	return runtime.GOOS
}

func (h *Helper) ConvertDate(dateStr string, layout string) (time.Time, error) {
	var date time.Time
	convert, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Erro ao converter a string para time.Time:", err)
		return date, err
	}
	return convert, nil
}

//func AddDays(date *time.Time, qtyDays int) time.Time {
//	return date.Add(time.Duration(qtyDays) * 24 * time.Hour)
//}
