/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// log level "reciever" var
var logLevel string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-kubernetes-controllers",
	Short: "Sample application with zerolog",
	Long:  "Lorem ipsum delor sit amet.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initLogger()
	},
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

func initLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	level := parseLogLevel(logLevel)
	configureLogger(level)
	// zerolog.SetGlobalLevel(level)

	// Logger = log.Logger
}

func parseLogLevel(lvl string) zerolog.Level {
	switch strings.ToLower(lvl) {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}

func configureLogger(level zerolog.Level) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.SetGlobalLevel(level)
	if level == zerolog.TraceLevel {
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			return fmt.Sprintf("%s:%d", file, line)
		}
		zerolog.CallerFieldName = "caller"
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "2006-01-02 15:04:05.000",
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.LevelFieldName,
				zerolog.CallerFieldName,
				zerolog.MessageFieldName,
			},
		}).With().Caller().Logger()
	} else if level == zerolog.DebugLevel {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "2006-01-02 15:04:05.000",
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.LevelFieldName,
				zerolog.MessageFieldName,
			},
		})
	} else {
		log.Logger = log.Output(os.Stderr)
	}
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
	// rootCmd.PersistentFlags().StringVar(&cfgloFile, "config", "", "config file (default is $HOME/.go-kubernetes-controllers.yaml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "", "Logging level file (default is 'info')")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
