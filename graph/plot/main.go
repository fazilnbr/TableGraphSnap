package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Example JSON data
var jsonData = `
[
  {"Name": "Alice", "Age": 30},
  {"Name": "Bob", "Age": 25},
  {"Name": "Charlie", "Age": 35},
  {"Name": "James", "Age": 55},
  {"Name": "Jon", "Age": 15},
  {"Name": "Tom", "Age": 36}
]
`

func main() {
	graphWithTwoValue()
	lineGraph()
}

func graphWithTwoValue() {
	// Parse JSON data
	var data []map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(jsonData)).Decode(&data); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Extract x and y values for the graph
	var names []string
	var ages []float64
	for _, item := range data {
		names = append(names, item["Name"].(string))
		ages = append(ages, item["Age"].(float64))
	}

	// Create a new plot
	p := plot.New()

	// Create a bar plotter
	bars, err := plotter.NewBarChart(plotter.Values(ages), vg.Points(50))
	if err != nil {
		log.Fatalf("Error creating bar chart: %v", err)
	}

	// Add the bars to the plot
	p.Add(bars)

	// Customize the plot ticks to label the bars
	p.NominalX(names...)

	// Set the title and labels
	p.Title.Text = "Age of People"
	p.X.Label.Text = "Name"
	p.Y.Label.Text = "Age"

	// Save the plot to a file
	if err := p.Save(10*vg.Inch, 6*vg.Inch, "file/graph.png"); err != nil {
		log.Fatalf("Error saving plot: %v", err)
	}
	fmt.Println("Graph saved as graph.png")
}

func lineGraph() {

	p := plot.New()

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "Name"
	p.Y.Label.Text = "Age"

	err := plotutil.AddLinePoints(p,
		"First", randomPoints())
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "file/points.png"); err != nil {
		panic(err)
	}
}

// randomPoints returns some random x, y points.
func randomPoints() plotter.XYs {

	// Parse JSON data
	var data []map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(jsonData)).Decode(&data); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Extract x and y values for the graph
	var names []string
	var ages []float64
	for _, item := range data {
		names = append(names, item["Name"].(string))
		ages = append(ages, item["Age"].(float64))
	}

	pts := make(plotter.XYs, len(ages))
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = ages[i]
	}
	return pts
}
