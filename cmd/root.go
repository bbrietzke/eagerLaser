package cmd

import (
	"fmt"
	"github.com/bbrietzke/eagerLaser/services"
	"github.com/spf13/cobra"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var cfgFile string
var githubProject string
var sender string
var destination string
var console bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eagerLaser",
	Short: "View Github Pull Requests",
	Long:  `View Github Pull Requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		v := strings.Split(githubProject, "/")

		g := services.GithubPullRequests(v[0], v[1])
		prs, err := g.GetPullRequests()

		if err == nil {
			if console {
				notifier := services.NewConsoleNotifier()
				notifier.Notify(prs)
			} else {
				notifier := services.NewEmailNotificationFormatter(destination, sender)
				notifier.Notify(prs)
			}
		} else {
			fmt.Println(err)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eagerLaser.yaml)")

	rootCmd.PersistentFlags().StringVar(&githubProject, "project", "google/go-github", "The Github Project to query")
	rootCmd.PersistentFlags().StringVar(&sender, "sender", "noone@example.org", "Who is sending the email")
	rootCmd.PersistentFlags().StringVar(&destination, "destination", "mygroup@example.org", "Who are we sending the message to?")
	rootCmd.PersistentFlags().BoolVar(&console, "console", false, "Display list to console")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".eagerLaser" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".eagerLaser")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
