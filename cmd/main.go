// 主要运行程序
// Created: 2025/4/15

package main

import (
	"fmt"
	"log"
	"time"
	"vub-auto-test/auth"
	"vub-auto-test/browser"
	"vub-auto-test/config"
)

func main() {
	fmt.Println("Program started")
	// 1. cg用户配置
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// 2. 获取URL
	loginURL, err := auth.GetLoginURL(cfg.Credentials.LoginName, cfg.Credentials.Password, cfg.Credentials.LoginTypeId, cfg.Credentials.UnivCode)
	if err != nil {
		log.Fatal("Error getting login URL: ", err)
	}
	fmt.Println("Login URL:", loginURL)

	// 3. 启动浏览器实例
	br, err := browser.NewBrowser(cfg.Browser)
	if err != nil {
		log.Fatal("Error creating browser: ", err)
	}
	defer br.Close()

	// 4. 打开网页
	err = browser.OpenPage(br.Context(), loginURL)
	if err != nil {
		log.Fatal("Error opening page: ", err)
	}
	fmt.Println("Browser opened.")
	time.Sleep(3 * time.Second)
}
