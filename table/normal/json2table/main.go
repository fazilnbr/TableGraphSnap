package main

import (
	"fmt"
	"os"

	"github.com/iotdog/json2table/j2t"
	"github.com/leesper/holmes"
)

func main() {
	defer holmes.Start().Stop()

	jsonStr := `[{"title1": "hello", "title2": "world"}, {"title1": "hello", "title2": "github"}, {"title1": "have", "title2": "fun"}]`
	ok, html := j2t.JSON2HtmlTable(jsonStr, []string{"title2", "title1"}, []string{"title1"})
	if ok {
		fmt.Println(html)
	} else {
		fmt.Println("failed to convert json to html table")
	}

	file, err := os.Create("json2table.html")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if _, err := file.Write([]byte(html)); err != nil {
		fmt.Println(err)
	}
}
