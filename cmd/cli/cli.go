package main

import (
	"time"

	"github.com/fagongzi/netproxy/cmd/cli/clicmd"
	"github.com/spf13/cobra"
)

const (
	cliName        = "cli"
	cliDescription = "A simple command line client for netproxy."

	defaultDialTimeout    = 2 * time.Second
	defaultCommandTimeOut = 5 * time.Second
)

var (
	rootCmd = &cobra.Command{
		Use:        cliName,
		Short:      cliDescription,
		SuggestFor: []string{"cli"},
	}
)

func main() {
	rootCmd.PersistentFlags().StringVar(&clicmd.Global.Endpoints, "endpoints", "127.0.0.1:8080", "netproxy api address")

	rootCmd.AddCommand(clicmd.NewListCommand(),
		clicmd.NewUpdateCommand(),
		clicmd.NewPauseCommand(),
		clicmd.NewResumeCommand())

	if err := rootCmd.Execute(); err != nil {
		return
	}
}
