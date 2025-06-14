package parser

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/JeanLeonHenry/gonotes/db"
)

type Result struct {
	StudentIndex  int
	QuestionIndex int
	Points        float64
}

type Test struct {
	StudentsNames  []string
	QuestionsNames []string
	PointTotals    []float64
	Results        []Result
	Desc           string
}

// TestFromRecords assumes records have been validated
// The returned Test object has students and questions numbered from 0
func TestFromRecords(records [][]string) Test {
	// Parse second line
	var pointsTotals []float64
	for _, val := range records[1][1:] {
		// we assume records have been validated
		maxPoints, _ := strconv.ParseFloat(val, 64)
		pointsTotals = append(pointsTotals, maxPoints)
	}

	// Parse first col and main content
	var students []string
	var results []Result
	for studentIndex, currentRecord := range records[2:] {
		students = append(students, currentRecord[0])
		for questionIndex, content := range currentRecord[1:] {
			points, _ := strconv.ParseFloat(content, 64)
			results = append(results, Result{
				StudentIndex:  studentIndex,
				QuestionIndex: questionIndex,
				Points:        points,
			})

		}
	}

	// Use first line as is
	return Test{
		StudentsNames:  students,
		QuestionsNames: records[0][1:],
		PointTotals:    pointsTotals,
		Results:        results,
	}
}

// Am I recreating an ORM ?

// TestFromDB gathers the test info from db into a Test struct
// INFO: Returned Test struct has no description
func TestFromDB(queries *db.Queries, ctx context.Context, t db.Test) (Test, error) {
	var test Test
	dbRows, err := queries.GetResultsFromTest(ctx, t.ID)
	if err != nil {
		return Test{}, err
	}
	// log.Println("Got the following rows from db : ", dbRows[:3], "...")
	for _, row := range dbRows {
		// Rows are sorted by student name then by question name
		// Check if we changed student or question
		if !slices.Contains(test.StudentsNames, row.StudentName) {
			test.StudentsNames = append(test.StudentsNames, row.StudentName)
		}
		if !slices.Contains(test.QuestionsNames, row.QuestionName.String) {
			test.QuestionsNames = append(test.QuestionsNames, row.QuestionName.String)
			test.PointTotals = append(test.PointTotals, row.MaxPoints)
		}
		// We're now sure those aren't -1
		studentIndex := slices.Index(test.StudentsNames, row.StudentName)
		questionIndex := slices.Index(test.QuestionsNames, row.QuestionName.String)
		// Build the current result
		test.Results = append(test.Results, Result{
			StudentIndex:  studentIndex,
			QuestionIndex: questionIndex,
			Points:        row.Points,
		})
	}
	// log.Println("Parsed into", test)
	return test, nil
}

func (t *Test) ExportReport() {
	// TODO: print individual reports in a pretty pdf
	// Use maroto.io ?

	// log.Println("Exporting report.")
}
