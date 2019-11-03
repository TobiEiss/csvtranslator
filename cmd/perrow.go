package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(perrowCmd)
}

var perrowCmd = &cobra.Command{
	Use:   "perrow",
	Short: "Gernerates for every row all rows of the template.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("csvtranslator v0.1 -- HEAD")
		fmt.Println(source)
	},
}
