/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-kubernetes-controllers",
	Short: "Sample application with zerolog",
	Long:  `Lorem ipsum delor sit amet.`,
	Run: func(cmd *cobra.Command, args []string) {
		// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		fmt.Println("Welcome to k8s-controller-tutorial CLI!")
		fmt.Println("This is a sample app for cobra-cli and zerolog demo")
		fmt.Println("Example of zerolog messages:")
		fmt.Println()
		log.Info().Msg("This is an info log")
		log.Debug().Msg("This is a debug log")
		log.Trace().Msg("This is a trace log")
		log.Warn().Msg("This is a warn log")
		log.Error().Msg("This is an error log")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	log.Debug().Msg("running application")
	err := rootCmd.Execute()
	if err != nil {
		log.Error().Err(err).Msg("error occured")
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-kubernetes-controllers.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
