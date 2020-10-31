package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "ec2meta [command]",
	Short:         "Display EC2 instance metadata",
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
