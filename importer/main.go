package importer

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/JeanLeonHenry/gonotes/db"
)

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

func GetStudentsNamesFromCSV(filePath string) ([]string, error) {
	records, err := CsvRead(filePath)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, v := range records {
		if len(v) == 0 {
			return nil, fmt.Errorf("Found a record with no student name")
		}
		names = append(names, v[0])
	}
	return names, nil
}

func StudentFactory(names []string, class string) []db.CreateStudentParams {
	var students []db.CreateStudentParams
	for _, name := range names {
		students = append(students, db.CreateStudentParams{Name: name, Class: class})
	}
	return students
}
