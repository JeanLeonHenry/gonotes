/*
Copyright © 2025 Jean-Léon HENRY
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/JeanLeonHenry/gonotes/db"
	"github.com/JeanLeonHenry/gonotes/importer"
	"github.com/JeanLeonHenry/gonotes/parser"
	"github.com/JeanLeonHenry/gonotes/utils"
	"github.com/spf13/cobra"
)

const (
	studentsArg = "students"
	resultsArg  = "results"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import {students|results} file.csv",
	Short: "Import data in the database, such as student names or test results. You should import students before importing results",
	Args:  cobra.ExactArgs(2),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == studentsArg {
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

		} else if args[0] == resultsArg {
			// Validate input
			records, err := importer.ValidateResultsCSV(args[1])
			if err != nil {
				log.Fatalln("Error validating input file:", err)
			}
			currentTest := parser.TestFromRecords(records)
			fmt.Println("Records parsed into\n", currentTest)
			// Ask for test info
			class := utils.AskUser("Class? ", utils.NotAllSpaces) // TODO: use suggestions from db
			testDate := utils.AskUser("Test date (YYYY-MM-DD)? ", utils.NotAllSpaces)
			parsedDate, err := time.Parse(time.DateOnly, testDate)
			if err != nil {
				log.Fatalln("Wrong date input")
			}
			sqlDesc := sql.NullString{
				String: utils.AskUser("Test description: ", utils.NotAllSpaces),
				Valid:  true,
			}
			// Create test
			testId, err := queries.CreateTestAndReturnID(ctx, db.CreateTestAndReturnIDParams{
				Date:        parsedDate,
				Description: sqlDesc,
			})
			if err != nil {
				log.Fatalln("Error creating test", err)
			}
			for _, currentResult := range currentTest.Results {
				// Find student
				studentName := currentTest.StudentsNames[currentResult.StudentIndex]
				// PERF: cache the id to improve speed
				studentId, err := queries.GetStudentId(ctx, db.GetStudentIdParams{
					Name:  studentName,
					Class: class,
				})
				if err != nil {
					// INFO: we assume the student must have been imported beforehand
					log.Fatalln("Error looking up student: ", err)
				}
				log.Println("Found student", studentName, "in class", class, "at id", studentId)
				// Find question, create if needed
				questionName := currentTest.QuestionsNames[currentResult.QuestionIndex]
				questionId, err := queries.GetQuestionID(ctx, db.GetQuestionIDParams{
					Name:   sql.NullString{String: questionName, Valid: true},
					TestID: testId,
				})
				if err != nil {
					if err == sql.ErrNoRows {
						log.Println("found no question with question name", questionName, "\nCreating it.")
						questionId, err = queries.CreateQuestionAndReturnID(ctx, db.CreateQuestionAndReturnIDParams{
							TestID:    testId,
							MaxPoints: currentTest.PointTotals[currentResult.QuestionIndex],
							Rank:      int64(currentResult.QuestionIndex),
							Name:      sql.NullString{String: questionName, Valid: true},
						})
					} else {
						log.Fatalln("Error looking up question: ", err)
					}
				}
				log.Println("Found question", questionName, "at id", questionId)
				queries.CreateResult(ctx, db.CreateResultParams{
					StudentID:  studentId,
					QuestionID: questionId,
					Points:     currentResult.Points,
				})
			}
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
