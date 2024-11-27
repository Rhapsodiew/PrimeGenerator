package pdf

import (
	"fmt"
	"strings"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
	Pirate "MyGeneratorPrime/pirate"
)

func New_pdf(pirate Pirate.Pirate) {
	// CREATE PDF
	pdf := gofpdf.New("P", "mm", "A4","")
	pdf.SetFontLocation("/home/semavo_t/MyGeneratorPrime/font_pdf/")
	pdf.AddPage()
	pdf.SetFont("Courier", "B", 40)
	
	// IMAGE BACKGROUND
	bg := "images/wantedVierge.jpg"
	pdf.Image(bg, 2, 0, 205, 0, false, "", 0, "")
	
	pdf.Ln(210)
	
	// pdf.SetFontSize(50)
	pdf.WriteAligned(200, 5, pirate.Name, "C")
	pdf.Ln(30)
	pdf.WriteAligned(0, 5, pirate.Prime, "C")

	// IMAGE PIRATE
	image_wanted := pirate.Img
	if strings.Contains(image_wanted,"https://"){
		httpimg.Register(pdf,image_wanted,"")
	}

	pdf.Image(image_wanted, 22, 64, 165, 123, false, "", 0, "")

	err := pdf.OutputFileAndClose("pdf/"+pirate.Name + ".pdf")
	if err != nil {
		fmt.Println(err.Error())
	}
}
