package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func getRowVisitor (s sqlWrapper) func(*xlsx.Row) error {
	return func(r *xlsx.Row) error {
		var pr product
		var y int
		err := r.ForEachCell(func(c *xlsx.Cell) error {
			var x int
			x, y = c.GetCoordinates()
			value, err := c.FormattedValue()
			if err != nil {
				panic(err)
			} else if y > 0 {
				switch x {
				case 0:
					pr.model = value
				case 1:
					pr.company = value
				case 2:
					pr.price, err = c.Int()
				}
			}

			return err
		})
		if y > 0 {
			s.insertProduct(pr)
		}
		return err
	}
}

func main() {
	db := connectDb("db.db")
	defer db.close()

	wb, err := xlsx.OpenFile("sampleFile.xlsx")
	if err != nil {
		panic(err)
	}
	if len(wb.Sheets) == 0 {
		fmt.Println("No sheets in workbook")
	} else {
		sh := wb.Sheet["Sample"]
		_ = sh.ForEachRow(getRowVisitor(db))
	}
}