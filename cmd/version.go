package cmd

import (
	"fmt"

	"github.com/invit/ec2meta/internal/lib/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ec2meta",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ec2meta %s -- %s\n", version.Version, version.Commit)
	},
}
