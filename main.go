package main

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"zap-learn/logger"
	"zap-learn/zap"
)

func watchLogEvent() {
	// 监听配置
	viper.WatchConfig()

	// reload event 事件
	viper.OnConfigChange(func(event fsnotify.Event) {
		// 配置文件热加载
		viper.Unmarshal(&zap.Config)
		logger.Infof("The config %s reload", event.Name)
		if err := logger.Setup(); err != nil {
			logger.Error(err)
			return
		}
	})
}

func main() {
	if err := logger.Setup(); err != nil {
		logger.Error(err)
		return
	}

	watchLogEvent()

	logger.FDebug("prefix", "key1", "value1", "key2", "value2")
	logger.FInfo("Base setup", "config path")

	// 测试日志热加载, 运行后更改 config.yaml 中的 level 等级进行测试
	for {
		logger.Debug("test-prefix-debug", "key", "value")
		time.Sleep(2 * time.Second)
	}
}
