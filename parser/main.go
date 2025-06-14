package parser

import (
	"context"
	"fmt"
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

// TestFromDB gathers the test info from db into a struct
func TestFromDB(queries *db.Queries, ctx context.Context, t db.Test) (Test, error) {
	var test Test
	dbRow, err := queries.GetResultsFromTest(ctx, t.ID)
	if err != nil {
		return Test{}, err
	}
	fmt.Println(dbRow)
	// TODO: finish up
	return test, nil
}

func (t *Test) ExportReport() {
	// TODO: print individual reports in a pretty pdf
	fmt.Println("Exporting report.")
}
