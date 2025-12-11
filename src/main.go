package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/SusanHex/RE-Peater/config"
	"github.com/spf13/viper"
)

func main() {
	slog.Info("Welcome to RE-Peater!")
	viper_instance := viper.NewWithOptions()
	app_config, err := config.GetConfigFromViper(viper_instance)
	if err != nil {
		panic(err)
	}
	slog.Info(fmt.Sprintf("Found %d messages", len(app_config.Messages)))
	slog.Info("Starting to write messages...")
	for {
		for _, log_message := range app_config.Messages {
			time.Sleep(time.Duration(log_message.Delay) * time.Millisecond)
			_, err = fmt.Println(log_message.Message)
			if err != nil {
				panic(err)
			}
		}
	}
}
