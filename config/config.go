// 读取配置并转为结构体
// Created: 2025/4/14

package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	LoginName   string `yaml:"login_name"`
	Password    string `yaml:"password"`
	LoginTypeId string `yaml:"login_type_id"`
	UnivCode    string `yaml:"univ_code"`
}

// LoadConfig 从指定路径读取 YAML 配置文件，并返回 Config 对象
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
