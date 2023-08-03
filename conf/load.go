package conf

import (
	"fmt"
	"net"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	global *Config
)

// C 全局配置对象
func C() *Config {
	if global == nil {
		panic("Load Config first")
	}
	return global
}

// LoadConfigFromToml 从toml中添加配置文件, 并初始化全局对象
func LoadConfigFromToml(filePath string) error {
	cfg := newConfig()
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return err
	}

	// hostip := getLocalIP()
	// if hostip != "" {
	// 	cfg.App.HTTP.Host = hostip
	// }
	fmt.Printf("cfg.App.HTTP.Host: %v\n", cfg.App.HTTP.Host)
	fmt.Printf("\"sssssssssssssssssssssssssssssssssssssssss\": %v\n", "sssssssssssssssssssssssssssssssssssssssss")
	// 加载全局配置单例
	global = cfg
	return nil
}

// LoadConfigFromEnv 从环境变量中加载配置
func LoadConfigFromEnv() error {
	cfg := newConfig()
	if err := env.Parse(cfg); err != nil {
		return err
	}
	// 加载全局配置单例
	global = cfg
	return nil
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf(fmt.Errorf("Failed to get local IP address: %v", err).Error())
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil && strings.HasPrefix(ipNet.IP.String(), "10.2.") {
			return ipNet.IP.String()
		}
	}
	fmt.Printf(fmt.Errorf("Failed to get local IP address: no matching IP address found").Error())
	return ""
}
