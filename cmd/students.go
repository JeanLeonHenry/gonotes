package cmd

import (
	"fmt"
	"log"

	"github.com/JeanLeonHenry/gonotes/importer"
	"github.com/JeanLeonHenry/gonotes/utils"
	"github.com/spf13/cobra"
)

// studentsCmd represents the students command
var studentsImporterCmd = &cobra.Command{
	Use:   "students file.csv",
	Short: "Import students names from csv",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		names, err := importer.GetStudentsNamesFromCSV(args[0])
		if err != nil {
			log.Fatalf("Error reading csv file : %v", err)
		}
		class := utils.AskUser("Class? ", utils.NotAllSpaces)
		params := importer.StudentFactory(names, class)
		log.Printf("About to write class %v with %v names from file %v", class, len(params), args[0])
		for _, student := range params {
			if err := queries.CreateStudent(ctx, student); err != nil {
				log.Fatalf("Error writing %v to class %v : %v", student, class, err)
			}
		}
		log.Printf("Successful write.")
	},
}

var studentsReaderCmd = &cobra.Command{
	Use:   "students",
	Short: "List students names from a class",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		class := args[0]
		students, err := queries.GetClass(ctx, class)
		if err != nil {
			log.Fatalf("Error reading students name from class %v : %v", class, err)
		}
		fmt.Printf("Class %v - %v students\n---\n", class, len(students))
		for _, student := range students {
			fmt.Println(student.Name)
		}

	},
}

func init() {
	importCmd.AddCommand(studentsImporterCmd)
	listCmd.AddCommand(studentsReaderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// studentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// studentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
