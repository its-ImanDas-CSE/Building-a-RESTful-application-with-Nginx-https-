package main

import (
	"encoding/csv" // A package for reading and writing CSV files.
	"fmt"
	"io"       // Provides basic I/O functionality like reading and writing.
	"log"      //  For logging errors and messages.
	"net/http" // Implements HTTP client and server functionality.
	"os"       //
	"strconv"  // Provides conversions between string and numeric types.

	"github.com/gin-gonic/gin" // A third-party package for creating HTTP web applications in Go (Gin web framework.
)

// CSV: A custom Go struct used to represent a single record from the CSV file.
// SiteID, FxiletID, Name, Criticality, RelevantComputerCount: Fields corresponding to columns in the CSV file.
type CSV struct {
	SiteID                string
	FxiletID              string
	Name                  string
	Criticality           string
	RelevantComputerCount string
}

// convertToJson(): This function reads a CSV file (fixlets.csv), processes its data, and converts it into a 2D slice of CSV structs.
func convertToJson() [][]CSV {

	jsonFiles := [][]CSV{}

	fmt.Println("inside the function") // Opens the fixlets.csv file.
	f, err := os.Open("fixlets.csv")

	if err != nil {

		log.Fatal("Not able to open a file  ", err.Error())
		return jsonFiles // return empty 2D Slice
	}

	defer f.Close()
	file := csv.NewReader(f)

	_, err1 := file.Read()
	if err1 != nil {
		log.Fatal("another error while reading the file ")
	}

	for {

		jsonFile := []CSV{}

		if err1 == io.EOF {
			return jsonFiles
		}

		if err1 != nil {
			log.Fatal("there is some error ", err.Error())
			return jsonFiles
		}

		for i := 0; i < 20; i++ {

			data, err := file.Read()
			err1 = err
			if err == io.EOF {
				return jsonFiles
			}

			if err != nil {
				log.Fatal("there is some error ", err.Error())
				return jsonFiles
			}

			record := CSV{
				SiteID:                data[0],
				FxiletID:              data[1],
				Name:                  data[2],
				Criticality:           data[3],
				RelevantComputerCount: data[4]}

			jsonFile = append(jsonFile, record)

			fmt.Println(jsonFile)
		}
		fmt.Println(jsonFile)
		jsonFiles = append(jsonFiles, jsonFile)

	}
	return jsonFiles

}
func readCSV() {

	f, err := os.Open("fixlets.csv")

	if err != nil {

		log.Fatal("Not able to open a file  ", err.Error())
	}
	file := csv.NewReader(f)
	if _, err1 := file.Read(); err1 != nil {
		log.Fatal("another error while reading the file ")
	}

	for {
		data, err := file.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("there is some error ", err.Error())
		}

		for i := 0; i < len(data); i++ {
			fmt.Print(data[i])
		}
		fmt.Println()

	}

}

func main() {

	jsonfile := convertToJson()

	// for _, records := range jsonfile {

	//  for _, record := range records {
	//      fmt.Println(record)
	//  }
	//  fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	// }
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, jsonfile)
	})
	r.GET("/ListFile", func(c *gin.Context) {

		pageNoStr := c.DefaultQuery("pageNo", "1")

		// Convert the pageNo from string to integer
		pageNo, err := strconv.Atoi(pageNoStr)
		if err != nil || pageNo <= 0 || pageNo > len(jsonfile) {
			// Handle error if the pageNo is invalid or out of range
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid page number or out of range",
			})
			return
		}

		c.JSON(http.StatusOK, jsonfile[pageNo-1])
	})

	r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
