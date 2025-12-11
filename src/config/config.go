package config

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type LogMessage struct {
	Delay   uint64
	Message string
}

type Config struct {
	MessageFilePath string

	Messages []LogMessage
}

func GetConfigFromViper(viper_instance *viper.Viper) (*Config, error) {
	viper_instance.BindEnv("message_separator")
	viper_instance.SetDefault("message_separator", ";;")
	message_separator := viper_instance.GetString("message_separator")

	viper_instance.BindEnv("time_designation_sequence")
	viper_instance.SetDefault("time_designation_sequence", "<<")
	time_designation := viper_instance.GetString("time_designation_sequence")

	viper_instance.BindEnv("messages")
	viper_instance.SetDefault("messages", "")
	raw_messages := viper_instance.GetString("messages")

	var messages []LogMessage
	if len(raw_messages) == 0 {
		messages = make([]LogMessage, 0)
	} else {
		messages = createLogMessages(raw_messages, message_separator, time_designation)
	}

	config := Config{Messages: messages, MessageFilePath: ""}
	return &config, nil
}

func createLogMessages(raw_message string, separator string, time_designation string) []LogMessage {
	time_designation_length := len(time_designation)
	log_messages := make([]LogMessage, 0)

	for sub_message := range strings.SplitSeq(raw_message, separator) {
		if len(sub_message) == 0 {
			continue
		}

		var milisecond_delay uint64 = 0
		message := ""

		if len(sub_message) >= time_designation_length && sub_message[:time_designation_length] == time_designation {

			split_sub_message := strings.SplitN(sub_message[time_designation_length:], " ", 2)
			raw_message_delay := split_sub_message[0]
			delay, err := strconv.ParseUint(raw_message_delay, 10, 64)

			if err != nil {
				slog.Error(fmt.Sprintf("failed to parse \"%s\" as a valid delay for message part: \"%s\"", raw_message_delay, sub_message))
			} else {
				if len(split_sub_message) == 2 {
					message = split_sub_message[1]
				}
				milisecond_delay = delay
			}
		} else {
			message = sub_message
		}
		if len(message) > 0 {
			log_messages = append(log_messages, LogMessage{
				Message: message, Delay: milisecond_delay,
			})
		}

	}
	return log_messages
}
