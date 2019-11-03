package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/TobiEiss/csvtranslator/pkg/reader"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "transform",
		Short: "Transforms a csv to another csv",
		Long: `A transformer to transform .csv-Files.
			Use help-command to get more informations`,
		Run: func(cmd *cobra.Command, args []string) {},
	}
	source    string
	output    string
	template  string
	seperator string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sourceCSV, err := reader.ReadCsv(source, seperator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println(sourceCSV)
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&source, "source", "s", "source.csv", "Source-file to transform")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "output.csv", "output-file which is transformed")
	rootCmd.PersistentFlags().StringVarP(&output, "template", "t", "template.csv", "template-file which should be applied")
	rootCmd.PersistentFlags().StringVarP(&seperator, "seperator", "c", ";", "the seperator of your CSV")
}
