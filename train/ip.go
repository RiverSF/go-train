package train

import (
	"errors"
	"net"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

func ParseIP(s string) int {
	ip := net.ParseIP(s)
	if ip == nil {
		return 0
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return 4
		case ':':
			return 6
		}
	}
	return 0
}

func GetClientIP(ctx *context.Context) string {
	clientIP := ctx.Request.Header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(ctx.Request.Header.Get("X-Real-Ip"))
	}
	if clientIP != "" {
		return clientIP
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(ctx.Request.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func GetLocalIp() (string, error) {
	// 获取本机所有的网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	// 遍历所有网络接口
	for _, iface := range interfaces {
		// 忽略down状态的接口
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		// 遍历接口上的所有地址
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 过滤IPv4和IPv6地址
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4() // 只保留IPv4地址，如果需要IPv6则注释掉这一行
			if ip == nil {
				continue
			} else {
				return ip.String(), nil
			}
		}
	}

	return "", errors.New("fail to getLocalIp")
}
