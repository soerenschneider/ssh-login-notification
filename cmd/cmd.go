package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sshnot/internal"
	"sshnot/internal/dispatcher/telegram"
	"sshnot/internal/formatter"
	"strings"
)

const (
	nogeoresolv   = "no-resolve-geo"
	nodnsresolv   = "no-resolve-dns"
	telegramtoken = "telegram-token"
	telegramId    = "telegram-id"
)

func Get() *cobra.Command {
	viper.AutomaticEnv()
	viper.SetConfigName("sshnotification")
	viper.AddConfigPath("/etc/default")
	viper.AddConfigPath("$HOME/.sshnotification")

	err := viper.ReadInConfig()
	if err != nil {
		log.Warnf("Couldn't read conf: %v", err)
	}

	main := &cobra.Command{
		Use: "sshnotification",
	}

	main.Flags().BoolP(nogeoresolv, "g", viper.GetBool(strings.Replace(nogeoresolv, "-", "_", -1)), "Do NOT lookup ip information")
	main.Flags().BoolP(nodnsresolv, "d", viper.GetBool(strings.Replace(nodnsresolv, "-", "_", -1)), "Do NOT lookup ip information")
	main.Flags().StringP(telegramtoken, "t", viper.GetString(strings.Replace(telegramtoken, "-", "_", -1)), "Telegram bot token")
	main.Flags().Int64P(telegramId, "i", viper.GetInt64(strings.Replace(telegramId, "-", "_", -1)), "Telegram message id")

	main.Run = func(cmd *cobra.Command, args []string) {
		options := parseOptions(main)
		scraper := internal.NewScrape(&options)
		formatted := formatter.Format(*scraper.Login)
		output, err := telegram.New(&options)
		if err != nil {
			log.Panic("Could not create telegram bot")
		}
		output.Send(formatted)
	}

	return main
}

func parseOptions(cmd *cobra.Command) internal.Options {
	options := internal.Options{}

	nogeoresolv, _ := cmd.Flags().GetBool(nogeoresolv)
	options.GeoLookup = !nogeoresolv

	nodnslookup, _ := cmd.Flags().GetBool(nodnsresolv)
	options.DnsLookup = !nodnslookup

	options.TelegramToken, _ = cmd.Flags().GetString(telegramtoken)
	options.TelegramId, _ = cmd.Flags().GetInt64(telegramId)

	return options
}
