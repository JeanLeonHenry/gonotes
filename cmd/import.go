/*
Copyright © 2025 Jean-Léon HENRY
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/JeanLeonHenry/gonotes/importer"
	"github.com/JeanLeonHenry/gonotes/utils"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import {students|results} file.csv",
	Short: "Import data in the database, such as student names or test results.",
	Args:  cobra.ExactArgs(2),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "students" {
			names, err := importer.GetStudentsNamesFromCSV(args[1])
			if err != nil {
				log.Fatalf("Error reading csv file : %v", err)
			}
			class := utils.AskUser("Class? ", utils.NotAllSpaces)
			params := importer.StudentFactory(names, class)
			log.Printf("About to write class %v with %v names from file %v", class, len(params), args[1])
			for _, student := range params {
				if err := queries.CreateStudent(ctx, student); err != nil {
					log.Fatalf("Error writing %v to class %v : %v", student, class, err)
				}
			}
			log.Printf("Successful write.")

		} else if args[0] == "results" {
			records, err := importer.ValidateResultsCSV(args[1])
			if err != nil {
				log.Fatalln("Error importing results : ", err)
			}
			fmt.Println(records)
			// TODO: finish results

			// class := utils.AskUser("Class? ", utils.NotAllSpaces)
			// test_date := utils.AskUser("Test date ? ", utils.NotAllSpaces)
			// for _, record := range records[2:] {
			// 	// Look up student
			// 	studentName := record[0]
			// 	studentId, err := queries.GetStudentId(ctx, db.GetStudentIdParams{
			// 		Name:  studentName,
			// 		Class: class,
			// 	})
			// 	if err != nil {
			// 		log.Fatalln("Error looking up student : ", err)
			// 	}
			// 	// Create test
			//
			// 	// Create results based on the record
			// 	for questionIndex, mark := range record[1:] {
			// 		points, err := strconv.ParseFloat(mark, 64)
			// 		if err != nil {
			// 			log.Fatalf("Parsing %v's record : couldn't parse %v as points", studentName, mark)
			// 		}
			// 		queries.CreateResult(ctx, db.CreateResultParams{
			// 			StudentID:  studentId,
			// 			QuestionID: questionIds[questionIndex],
			// 			Points:     points,
			// 		})
			// 	}
			// }
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
