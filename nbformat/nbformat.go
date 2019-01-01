package nbformat

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type CodeCell struct {
	Source []string `json:"source"`
	CellType string `json:"cell_type"`
	Metadata map[string]interface{} `json:"metadata"`
	ExecutionCount int `json:"execution_count"`
	Outputs []map[string]interface{} `json:"outputs"`
}

type NBNode struct {
	Metadata map[string]interface{} `json:"metadata"`
	Nbformat int `json:"nbformat"`
	NbformatMinor int `json:"nbformat_minor"`
	Cells []CodeCell `json:"cells"`
}

func Read(fname string) NBNode {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	var nb NBNode
	json.Unmarshal(data, &nb)
	return nb
}

func CreateCodeCell() {

}

func CreateMarkdownCell() {

}

func CreateRawNotebookCell() {

}

func Write() {

}

func CreateNewNotebook() {

}

func ValidateNotebook() {

}

