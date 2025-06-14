/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/JeanLeonHenry/gonotes/parser"
	"github.com/JeanLeonHenry/gonotes/utils"
	"github.com/spf13/cobra"

	fzf "github.com/ktr0731/go-fuzzyfinder"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate the report for a test",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Get tests
		tests, err := queries.GetTest(ctx)
		if err != nil {
			log.Fatalln("Error getting tests")
		}

		choice, _ := utils.AskUserAndSuggest(
			"Pick test > ",
			tests,
			utils.TrueFunc, // TODO: maybe implement a real validation step ?
			func(i int) string {
				currentTest := tests[i]
				date := currentTest.Date.Format(time.RFC822)
				if currentTest.Description.Valid {
					return fmt.Sprintf("%v - %v", currentTest.Description.String, date)
				} else {
					return fmt.Sprintf("%v", date)
				}
			},
			fzf.WithPreviewWindow(func(i, width, height int) string {
				if i == -1 {
					return ""
				}
				classes, err := queries.GetClassFromTest(ctx, tests[i].ID)
				if err != nil {
					log.Fatalln("Error getting classes")
				}
				return strings.Join(classes, "\n")
			}))
		fmt.Println("chose", choice)
		test, err := parser.TestFromDB(queries, ctx, choice)
		if err != nil {
			log.Fatalf("Error parsing test from : %v", err)
		}
		test.ExportReport()
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
