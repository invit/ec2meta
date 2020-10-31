package cmd

import (
	"fmt"

	"github.com/invit/ec2meta/internal/lib/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get <path>",
	Short: "Returns arbitrary metadata by path",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()

		if err != nil {
			return err
		}

		list, err := c.ResolvePath(args[0])

		if err != nil {
			return err
		}

		if len(list) == 0 {
			return fmt.Errorf("Path %s not found", args[0])
		}

		for _, r := range list {
			fmt.Println(r)
		}

		return nil
	},
}
