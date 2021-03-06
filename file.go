package main

import (
	"encoding/csv"
	"os"
)

func generateStudentScoreCSVTemplate(data []Student) {
	csvfile, err := os.Create("hw_" + HW_ID + ".csv")

	if err != nil {
		panic(err)
	}

	csvwriter := csv.NewWriter(csvfile)

	csvwriter.Write([]string{"student_ID", "student_name", "submit_ID", "Score", "Comment"})
	for _, row := range data {
		_ = csvwriter.Write([]string{row.ID, row.Name, row.SubmitID, "", ""})
	}

	csvwriter.Flush()

	csvfile.Close()
}

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
