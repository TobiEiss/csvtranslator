package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of csvtranslator",
	Long:  `All software has versions. This is csvtranslators's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("csvtranslator v0.1 -- HEAD")
	},
}
