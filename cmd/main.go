// 主要运行程序
// Created: 2025/4/15

package main

import (
	"fmt"
	"log"
	"vub-auto-test/auth"
	"vub-auto-test/config"
)

func main() {
	fmt.Println("Program started")
	// 用户信息读取器
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// 合成URL跳转至对应学校页面
	loginURL, err := auth.GetLoginURL(cfg.LoginName, cfg.Password, cfg.LoginTypeId, cfg.UnivCode)
	if err != nil {
		log.Fatal("Error getting login URL: ", err)
	}
	fmt.Println("Login URL:", loginURL)
}
