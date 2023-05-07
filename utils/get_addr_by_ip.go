package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"net"
)

func GetAddrByGin(c *gin.Context) (ip, addr string) {
	ip = c.ClientIP()
	addr = GetAddr(ip)
	return ip, addr
}

func GetAddr(ip string) string {
	parseIP := net.ParseIP(ip)
	if IsIntranetIP(parseIP) {
		return "内网地址"
	}

	record, err := global.AddrDB.City(net.ParseIP(ip))
	if err != nil {
		return "错误的地址"
	}
	var province string
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}

	city := record.City.Names["zh-CN"]
	return fmt.Sprintf("%s-%s", province, city)
}

// IsIntranetIP 内网地址判断
func IsIntranetIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}
	fmt.Println(ip4)
	// 192.168
	// 172.16 -172.31
	// 10
	// 169.254
	return ip4[0] == 192 && ip4[1] == 168 ||
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) ||
		(ip4[0] == 10) ||
		(ip4[0] == 169 && ip4[1] == 254)
}
