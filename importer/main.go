package importer

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/JeanLeonHenry/gonotes/db"
	"github.com/JeanLeonHenry/gonotes/utils"
)

// CsvRead reads a csv file given its path.
// Uses ReadAll, potentially dangerous on big files.
func CsvRead(filePath string) ([][]string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	records, err := csv.NewReader(bytes.NewReader(file)).ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func ValidateResultsCSV(filePath string) ([][]string, error) {
	records, err := CsvRead(filePath)
	if err != nil {
		return nil, err
	}
	// Validate records
	if len(records) < 3 {
		return nil, fmt.Errorf("Found less than 3 lines of csv")
	}
	size := len(records[0])
	if size < 2 {
		return nil, fmt.Errorf("First line is of length %v, should be at least 2", size)
	}
	if len(records[1]) != size {
		return nil, fmt.Errorf("First two lines don't have the same length")
	}
	// Extract info from first two lines
	questionNames := records[0][1:]
	pointTotals := records[1][1:]
	// All questions names should be parsable into ranks
	if s, index, err := utils.AllMatch(questionNames, `^E\d+(?:[\.\-\/:]\w+)*$`); err != nil {
		log.Println(questionNames)
		return nil, fmt.Errorf("Question name %v at col %v can't be parsed with error '%v'.", s, index+1, err)
	}
	// All points totals should be parsable into floats
	if s, index, err := utils.AllParsableIntoFloats(pointTotals); err != nil {
		log.Println(pointTotals)
		return nil, fmt.Errorf("Point total %v at col %v can't be parsed to float.", s, index+1)
	}
	// Validate the remaining points records
	for k, currentRecord := range records[2:] {
		if len(currentRecord) != size {
			return nil, fmt.Errorf("Records length mismatch on line %v", k)
		}
		if s, index, err := utils.AllParsableIntoFloats(currentRecord[1:]); err != nil {
			return nil, fmt.Errorf("Point total %v at line %v and col %v can't be parsed to float.", s, k+3, index+2)
		}
	}
	return records, nil
}

// GetStudentsNamesFromCSV reads the given file and return the first column, skipping the first two lines.
// TODO: tests : empty first field in a record...
func GetStudentsNamesFromCSV(filePath string) ([]string, error) {
	records, err := CsvRead(filePath)
	if err != nil {
		return nil, err
	}
	var names []string
	for lineNumber, record := range records[2:] {
		if len(record) == 0 {
			return nil, fmt.Errorf("Found a record with no student name on line %v", lineNumber)
		}
		names = append(names, record[0])
	}
	return names, nil
}

func StudentFactory(names []string, class string) (students []db.CreateStudentParams) {
	for _, name := range names {
		students = append(students, db.CreateStudentParams{
			Name:  name,
			Class: class,
		})
	}
	return
}
