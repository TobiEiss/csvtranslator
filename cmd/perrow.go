package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/TobiEiss/csvtranslator/pkg/reader"
	"github.com/TobiEiss/csvtranslator/pkg/transformer"
	"github.com/TobiEiss/csvtranslator/pkg/writer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(perrowCmd)
}

var perrowCmd = &cobra.Command{
	Use:   "perrow",
	Short: "Gernerates for every row all rows of the template.",
	Run: func(cmd *cobra.Command, args []string) { // read source
		sourceCSV, err := reader.ReadCsv(source, seperator)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// read template
		templateCSV, err := reader.ReadCsv(template, seperator)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// transform
		result := transformer.TransformPerRow(sourceCSV, templateCSV)

		if output != "" {
			// print to file
			if err := writer.WriteCSV(output, result); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			// print to console
			csvWriter := csv.NewWriter(os.Stdout)
			csvWriter.Comma = rune(seperator[0])
			for _, record := range result {
				if err := csvWriter.Write(record); err != nil {
					log.Fatalln("error writing record to csv:", err)
				}
			}

			// Write any buffered data to the underlying writer (standard output).
			csvWriter.Flush()

			if err := csvWriter.Error(); err != nil {
				log.Fatal(err)
			}
		}

	},
}
