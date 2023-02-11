package cmd

import (
	"fmt"
	"squarecloud/internal/build"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print cli version",
	Run: func(cmd *cobra.Command, args []string) {
		version := build.Version

		fmt.Printf("Square Cloud CLI version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
