// 读取配置并转为结构体
// Created: 2025/4/14

package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Credentials Credentials `yaml:"credentials"`
	Browser     Browser     `yaml:"browser"`
}

type Credentials struct {
	LoginName   string `yaml:"login_name"`
	Password    string `yaml:"password"`
	LoginTypeId string `yaml:"login_type_id"`
	UnivCode    string `yaml:"univ_code"`
}

type Browser struct {
	Headless   bool          `yaml:"headless"`
	Timeout    time.Duration `yaml:"timeout"`
	UserAgent  string        `yaml:"user_agent"`
	WindowSize struct {
		Width  int `yaml:"width"`
		Height int `yaml:"height"`
	} `yaml:"window_size"`
	ExecPath string `yaml:"exec_path"`
	Proxy    string `yaml:"proxy"`
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

	// 设置浏览器默认值
	if cfg.Browser.Timeout == 0 {
		cfg.Browser.Timeout = 30 * time.Second
	}
	if cfg.Browser.WindowSize.Width == 0 {
		cfg.Browser.WindowSize.Width = 1920
	}
	if cfg.Browser.WindowSize.Height == 0 {
		cfg.Browser.WindowSize.Height = 1080
	}

	return &cfg, nil
}
