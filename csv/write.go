package csv

import (
	"encoding/csv"
	"github.com/DanielFillol/API_Caller/models"
	"os"
	"path/filepath"
)

//Write function takes in a fileName string, a folderName string, and a slice of models.WriteStruct named decisions.
//	It creates a CSV file in the specified folder with the given file name and writes the data from decisions into the file.
//	This function uses two helper functions, createFile and generateRow.
func Write(fileName string, folderName string, decisions []models.WriteStruct) error {
	var rows [][]string

	rows = append(rows, generateHeaders())

	for _, decision := range decisions {
		rows = append(rows, generateRow(decision))
	}

	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		return err
	}

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

//createFile function takes in a file path and creates a file in the specified directory. It returns a pointer to the created file and an error if there is any.
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

//generateHeaders function returns a slice of strings containing the header values for the CSV file.
func generateHeaders() []string {
	return []string{
		"SearchedName",
		"ID",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"Name",
		"Classification",
		"Metaphone",
		"NameVariations",
	}
}

//generateRow function takes in a single models.WriteStruct argument and returns a slice of strings containing the values to be written in a row of the CSV file.
//	It uses a loop to concatenate all the NameVariations into a single string separated by " | "
func generateRow(result models.WriteStruct) []string {
	return []string{
		result.SearchName,
		result.ID,
		result.CreatedAt,
		result.UpdatedAt,
		result.DeletedAt,
		result.Name,
		result.Classification,
		result.Metaphone,
		result.NameVariations,
	}
}
