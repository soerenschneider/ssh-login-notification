package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sshnot/internal"
	"sshnot/pkg"
	"strings"
)

const (
	nogeoresolv   = "no-resolve-geo"
	nodnsresolv   = "no-resolve-dns"
	telegramtoken = "telegram-token"
	telegramId    = "telegram-id"
)

// FireUp parses the user supplied input and starts this whole mess.
func FireUp() *cobra.Command {
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

	main.Flags().BoolP(nogeoresolv, "g", viper.GetBool(strings.Replace(nogeoresolv, "-", "_", -1)), "Do NOT lookup ip geo information")
	main.Flags().BoolP(nodnsresolv, "d", viper.GetBool(strings.Replace(nodnsresolv, "-", "_", -1)), "Do NOT lookup dns information")
	main.Flags().StringP(telegramtoken, "t", viper.GetString(strings.Replace(telegramtoken, "-", "_", -1)), "Telegram bot token")
	main.Flags().Int64P(telegramId, "i", viper.GetInt64(strings.Replace(telegramId, "-", "_", -1)), "Telegram message id")

	main.Run = func(cmd *cobra.Command, args []string) {
		options := parseOptions(main)
		pkg.Cortex(&options)
	}

	return main
}

// parseOptions converts the parsed options to the internally
// used Options struct.
func parseOptions(cmd *cobra.Command) internal.Options {
	options := internal.Options{}

	noGeoResolve, _ := cmd.Flags().GetBool(nogeoresolv)
	options.GeoLookup = !noGeoResolve

	noDnsLookup, _ := cmd.Flags().GetBool(nodnsresolv)
	options.DnsLookup = !noDnsLookup

	options.TelegramToken, _ = cmd.Flags().GetString(telegramtoken)
	options.TelegramId, _ = cmd.Flags().GetInt64(telegramId)

	return options
}
