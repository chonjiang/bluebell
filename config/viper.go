package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// InitViper 初始化viper，管理配置文件
func InitViper(filePath string) error {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {             // 读取配置信息失败
		return fmt.Errorf("viper.ReadInConfig() failed, err: %v\n", err)
	}

	if err := viper.Unmarshal(Conf); err != nil {
		return fmt.Errorf("unmarshal conf failed, err:%s \n", err)
	}

	return nil
}
