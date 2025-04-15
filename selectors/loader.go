//  -
// Created: 2025/4/15

package selectors

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type SelectorConfig map[string]interface{}

// Load 动态加载为嵌套map
func Load(path string) (SelectorConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config SelectorConfig
	err = yaml.Unmarshal(data, &config)
	return config, err
}

// Get 路径式访问函数
func Get(config SelectorConfig, path string) (string, error) {
	paths := strings.Split(path, ".")
	current := config
	for i, path := range paths {
		value, exists := current[path]
		if !exists {
			return "", fmt.Errorf("path does not exist: %s", path)
		}

		switch v := value.(type) {
		case SelectorConfig:
			current = v
		case string:
			if len(paths)-i > 1 {
				return "", fmt.Errorf("path is incomplete: %s", path)
			}
			return v, nil
		default:
			return "", fmt.Errorf("invalid node type: %s", path)
		}
	}
	return "", fmt.Errorf("selector not found")
}
