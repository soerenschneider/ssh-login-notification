package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/soerenschneider/ssh-login-notification/internal"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg"
	"strings"
)

const (
	noGeoResolve     = "no-resolve-geo"
	noDnsResolve     = "no-resolve-dns"
	ignorePrivateIps = "ignore-private-ips"
	telegramToken    = "telegram-token"
	telegramId       = "telegram-id"
)

// FireUp parses the user supplied input and starts this whole mess.
func FireUp() *cobra.Command {
	viper.AutomaticEnv()
	viper.SetConfigName("ssh-login-notification")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/default/")
	viper.AddConfigPath("$HOME/.config/ssh-login-notification")

	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		log.Errorf("Didn't find config file: %v", err)
	} else if err != nil {
		log.Errorf("Couldn't read config file: %v", err)
	}

	main := &cobra.Command{
		Use: "ssh-login-notification",
	}

	main.Flags().BoolP(noGeoResolve, "g", viper.GetBool(strings.Replace(noGeoResolve, "-", "_", -1)), "Do NOT lookup ip geo information")
	main.Flags().BoolP(noDnsResolve, "d", viper.GetBool(strings.Replace(noDnsResolve, "-", "_", -1)), "Do NOT lookup dns information")
	main.Flags().BoolP(ignorePrivateIps, "p", viper.GetBool(strings.Replace(ignorePrivateIps, "-", "_", -1)), "Do nothing when a private IP is detected")
	main.Flags().StringP(telegramToken, "t", viper.GetString(strings.Replace(telegramToken, "-", "_", -1)), "Telegram bot token")
	main.Flags().Int64P(telegramId, "i", viper.GetInt64(strings.Replace(telegramId, "-", "_", -1)), "Telegram message id")

	main.Run = func(cmd *cobra.Command, args []string) {
		options := parseOptions(main)
		pkg.NewCortex(&options).Run()
	}

	return main
}

// parseOptions converts the parsed options to the internally used Options struct.
func parseOptions(cmd *cobra.Command) internal.Options {
	options := internal.Options{}

	noGeoResolve, _ := cmd.Flags().GetBool(noGeoResolve)
	options.GeoLookup = !noGeoResolve

	noDnsLookup, _ := cmd.Flags().GetBool(noDnsResolve)
	options.DnsLookup = !noDnsLookup

	options.IgnorePrivateIps, _ = cmd.Flags().GetBool(ignorePrivateIps)

	options.TelegramToken, _ = cmd.Flags().GetString(telegramToken)
	options.TelegramId, _ = cmd.Flags().GetInt64(telegramId)

	return options
}
