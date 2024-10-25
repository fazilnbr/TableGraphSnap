package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/chromedp/chromedp"
)

type Data struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// JSON data
	jsonData := `[{"name":"John Doe","age":30,"email":"john@example.com"},{"name":"Jane Smith","age":25,"email":"jane@example.com"}]`

	// Parse JSON data
	var data []Data
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Generate HTML content
	htmlContent := generateHTML(data)

	// Write HTML content to a file
	err = ioutil.WriteFile("table.html", []byte(htmlContent), 0644)
	if err != nil {
		fmt.Println("Error writing HTML file:", err)
		return
	}

	// Take a screenshot of the HTML file
	err = takeScreenshot(htmlContent, "screenshot.png")
	if err != nil {
		fmt.Println("Error taking screenshot:", err)
		return
	}

	fmt.Println("Screenshot taken successfully.")
}

func generateHTML(data []Data) string {
	// HTML template
	htmlTemplate := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>JSON to HTML Table</title>
		<style>
			table {
				font-family: Arial, sans-serif;
				border-collapse: collapse;
				width: 100%;
			}
			th, td {
				border: 1px solid #dddddd;
				text-align: left;
				padding: 8px;
			}
			tr:nth-child(even) {
				background-color: #f2f2f2;
			}
		</style>
	</head>
	<body>
		<table>
			<tr>
				<th>Name</th>
				<th>Age</th>
				<th>Email</th>
			</tr>
			{{range .}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.Age}}</td>
				<td>{{.Email}}</td>
			</tr>
			{{end}}
		</table>
	</body>
	</html>`

	// Create a new template and parse the HTML content
	tmpl := template.Must(template.New("html").Parse(htmlTemplate))

	// Execute the template with the provided data
	var resultHTML bytes.Buffer
	err := tmpl.Execute(&resultHTML, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return ""
	}

	return resultHTML.String()
}

func takeScreenshot(html string, filename string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate("data:text/html,"+html),
		chromedp.WaitReady("body"),
		chromedp.FullScreenshot(&buf, 100),
		// chromedp.CaptureScreenshot(&buf),
	); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(buf); err != nil {
		return err
	}

	return nil
}
