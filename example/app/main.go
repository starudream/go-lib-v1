package main

import (
	"github.com/starudream/go-lib/flag"
	"github.com/starudream/go-lib/log"
)

var rootCmd = &flag.Command{
	Use:   "root [sub]",
	Short: "My root command",
	Run: func(cmd *flag.Command, args []string) {
		log.Info().Msgf("inside rootCmd Run with args: %v\n", args)
	},
}

var subCmd = &flag.Command{
	Use:   "sub [no options!]",
	Short: "My subcommand",
	Run: func(cmd *flag.Command, args []string) {
		log.Info().Msgf("inside subCmd Run with args: %v\n", args)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Msgf("rootCmd init fail: %v", err)
	}
}

func main() {
	log.Info().Msgf("hello world")
}
