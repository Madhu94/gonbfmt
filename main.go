package main

import (
	"fmt"
	"gonbformat/nbformat"
)

func main() {
	nb := nbformat.Read("codenb.ipynb")
	cells := nb.Cells
	for i := 0; i < len(cells); i++ {
		fmt.Println(cells[i].Source)
	}
}