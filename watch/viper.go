package watch

import (
	"bluebell/config"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// WatchConfigChanged 实时监听配置文件变化
func WatchConfigChanged() error {

	viper.WatchConfig()
	// 回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被人修改啦...")
		zap.L().Info("config file changed", zap.String("name", in.Name))
		if err := viper.Unmarshal(config.Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
	return nil
}
