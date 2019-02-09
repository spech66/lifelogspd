package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

// Journal contains csv/json data
type Journal struct {
	Date string `json:"date"`
	Text string `json:"text"`
}

// JournalList contains csv/json data
type JournalList struct {
	Date string `json:"date"`
}

// GetJournal returns all journal list data
func GetJournal() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		journalList := journalReadAllData(c.Get("config").(*helper.Config).JournalPath)
		if journalList == nil {
			return c.JSONBlob(http.StatusOK, []byte(`[]`))
		}
		jsonData := helper.DataToJSON(&journalList)
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}

// GetJournalByDate returns specific journal data
func GetJournalByDate() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		date := c.Param("date")
		fmt.Println("Get by date", date)
		if date == "" || !helper.CheckDateSeperated(date) {
			return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
		}

		journal := journalReadData(c.Get("config").(*helper.Config).JournalPath, date)
		jsonData := helper.DataToJSON(&[]Journal{journal})
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}

// PostJournal saves new journal data
func PostJournal() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		journal := Journal{
			Date: c.FormValue("date"),
			Text: c.FormValue("text"),
		}
		fmt.Println(journal)
		if journal.Date == "" || !helper.CheckDateSeperated(journal.Date) {
			return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
		}

		journalAddData(c.Get("config").(*helper.Config).JournalPath, &journal)
		jsonData := helper.DataToJSON(&[]Journal{journal})
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}

// DeleteJournal deletes journal data
func DeleteJournal() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		date := c.Param("date")
		fmt.Println("Delete by date", date)
		if date == "" || !helper.CheckDateSeperated(date) {
			return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
		}

		deleted := journalDeleteFile(c.Get("config").(*helper.Config).JournalPath, date)
		if !deleted {
			return c.JSONBlob(http.StatusNotFound, []byte(`[]`))
		}
		return c.JSONBlob(http.StatusOK, []byte(`[]`))
	}
}

func journalReadAllData(path string) []JournalList {
	fmt.Println("Journal data from", path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var journals []JournalList
	for _, f := range files {
		if f.IsDir() {
			subPath := filepath.Join(path, f.Name())
			fmt.Println(subPath)

			subPathFiles, err := ioutil.ReadDir(subPath)
			if err != nil {
				panic(err)
			}

			for _, sf := range subPathFiles {
				if !sf.IsDir() {
					data := JournalList{
						Date: sf.Name(),
					}
					journals = append(journals, data)
				}
			}
		}
	}

	return journals
}

func journalReadData(path string, date string) Journal {
	filename := filepath.Join(path, date[:4], date)
	fmt.Println("Journal data from", filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Journal not found, creating empty one", filename)
		data := Journal{
			Date: date,
			Text: "",
		}
		return data
	}

	data := Journal{
		Date: date,
		Text: string(content[:]),
	}
	return data
}

func journalAddData(path string, data *Journal) {
	yearPath := filepath.Join(path, data.Date[:4])
	filename := filepath.Join(yearPath, data.Date)
	fmt.Println("Write journal at", filename)

	// Check for YEAR directory
	if _, err := os.Stat(yearPath); os.IsNotExist(err) {
		fmt.Println("Create journal dir", yearPath)
		err = os.MkdirAll(yearPath, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(data.Text)
	if err != nil {
		panic(err)
	}
}

func journalDeleteFile(path string, date string) bool {
	filename := filepath.Join(path, date[:4], date)
	fmt.Println("Delete journal", filename)

	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Delete journal did not find file", filename)
		return false
	}

	fmt.Println("Deleted journal", filename)
	return true
}
