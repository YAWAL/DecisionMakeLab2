package reporter

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const (
	lp                = "little plant "
	bp                = "big plant "
	w                 = "wait "
	hd                = "high demand "
	ld                = "low demand "
	pi                = "positive information "
	ni                = "negative information "
	dnb               = "do not build "
	calculationString = "5*%.2f*%.2f=%.2f"
	borderString      = "1"
	alignString       = "CM"
	linkString        = ""
)

func Report(calculations map[string]string, input [14]float64) error {
	file := "report.pdf"
	pdf := gofpdf.New("L", "mm", "A4", linkString)

	pdf.SetFont("Arial", linkString, 10)
	basicTable := func() {
		pdf.SetXY(90, 40)

		pdf.CellFormat(20, 20, "A/"+calculations["A"], borderString, 0, alignString, false, 0, linkString)

		pdf.SetXY(20, 95)
		pdf.CellFormat(20, 20, "1/"+calculations["1"], borderString, 0, alignString, false, 0, linkString)

		pdf.SetX(90)
		pdf.CellFormat(20, 20, "B/"+calculations["B"], borderString, 0, alignString, false, 0, linkString)

		pdf.SetX(230)
		pdf.CellFormat(20, 20, "A1/"+calculations["A1"], borderString, 0, alignString, false, 0, linkString)

		pdf.SetXY(90, 155)
		pdf.CellFormat(20, 20, "C/"+calculations["C"], borderString, 0, alignString, false, 0, linkString)

		pdf.SetX(230)
		pdf.CellFormat(20, 20, "B1/"+calculations["B1"], borderString, 0, alignString, false, 0, linkString)

		pdf.SetXY(170, 125)
		pdf.CellFormat(20, 20, "2/"+calculations["2"], borderString, 0, alignString, false, 0, linkString)

	}

	drawLines := func() {
		//1
		pdf.Line(40, 95, 90, 50)
		// 1-2
		pdf.Line(110, 40, 120, 30)
		//2
		pdf.Line(120, 30, 170, 30)
		// 1-3
		pdf.Line(110, 60, 120, 70)
		// 3
		pdf.Line(120, 70, 170, 70)
		//4
		pdf.Line(40, 105, 90, 105)
		//4-5
		pdf.Line(110, 95, 120, 85)
		//5
		pdf.Line(120, 85, 170, 85)
		// 4-6
		pdf.Line(110, 115, 120, 125)
		// 6
		pdf.Line(120, 125, 165, 125)
		//7
		pdf.Line(40, 115, 90, 165)
		//7-8
		pdf.Line(110, 155, 120, 145)
		//8
		pdf.Line(120, 145, 170, 135)
		//7-9
		pdf.Line(110, 175, 120, 185)
		//9
		pdf.Line(120, 185, 170, 185)
		// 10
		pdf.Line(190, 125, 230, 105)
		//10-12
		pdf.Line(250, 95, 260, 85)
		//12
		pdf.Line(260, 85, 290, 85)
		//10-13
		pdf.Line(250, 115, 255, 120)
		// 13
		pdf.Line(255, 120, 290, 120)
		//11
		pdf.Line(190, 145, 230, 165)
		//11-14
		pdf.Line(250, 155, 255, 150)
		//14
		pdf.Line(255, 150, 290, 150)
		//11-15
		pdf.Line(250, 175, 260, 185)
		//15
		pdf.Line(260, 185, 290, 185)

	}
	text := func() {
		pdf.SetFont("Arial", linkString, 12)

		pdf.Text(50, 100, lp)
		pdf.Text(50, 65, bp)
		pdf.Text(60, 135, w)

		pdf.Text(60, 135, w)

		pdf.Text(130, 27, hd)
		pdf.Text(130, 37, calcLine(input[2], input[1]))
		pdf.Text(130, 67, ld+calcLine(input[4], input[3]))

		pdf.Text(130, 82, hd)
		pdf.Text(130, 92, calcLine(input[7], input[6]))
		pdf.Text(130, 122, ld)
		pdf.Text(130, 115, calcLine(input[9], input[8]))

		pdf.Text(130, 135, pi)
		pdf.Text(130, 182, ni)
		pdf.Text(130, 191, dnb)

		pdf.Text(200, 162, lp)
		pdf.Text(200, 107, bp)

		pdf.Text(260, 83, hd)
		pdf.Text(260, 75, calcLine(input[12], input[1]))
		pdf.Text(260, 125, ld)
		pdf.Text(255, 131, calcLine(input[13], input[3]))

		pdf.Text(260, 143, hd)
		pdf.Text(250, 149, calcLine(input[12], input[6]))
		pdf.Text(260, 183, ld)
		pdf.Text(250, 190, calcLine(input[13], input[8]))

	}

	pdf.AddPage()
	basicTable()
	drawLines()
	text()
	return pdf.OutputFileAndClose(file)
}

func calcLine(probability, value float64) string {
	res := 5 * value * probability
	return fmt.Sprintf(calculationString, value, probability, res)
}
