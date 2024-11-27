package parse

import (
	pdf "MyGeneratorPrime/new_pdf"
	Pirates "MyGeneratorPrime/pirate"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func Parsecsv(filepath string) {
	// OPEN
	csvfile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	// SET NEWREADER AND
	csvreader := csv.NewReader(csvfile)
	data, err := csvreader.ReadAll()
	if err != nil {
		fmt.Print(err.Error())
	}
	// replace nuber . by "" 
	for _, i := range data {
		i[1] = strings.ReplaceAll(i[1], ".", "")

		new_pirate, err := Pirates.New(i[0], i[1], i[2])
		if err != nil {
			fmt.Println(err.Error())
			// fmt.Println(new_pirate)
			return
		}
		pdf.New_pdf(new_pirate)
	}
}
