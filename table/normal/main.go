package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chromedp/chromedp"
)

var jsonData = `
[{
  "id": "a197rn4az31sil2024-01-04b0751ypyk87nk8rk0cblr0nwz@marketplace.amazon.comusfl",
  "seller_id": "A197RN4AZ31SIL",
  "date": "2024-01-04",
  "customer": "7nk8rk0cblr0nwz@marketplace.amazon.com",
  "state_or_region": "fl",
  "country": "us",
  "asin": "B0751YPYK8",
  "cltv": "9.99",
  "operation": "UPSERT",
  "updated_at": "2024-01-31 10:01:39.620458 UTC"
}, {
  "id": "a197rn4az31sil2024-01-05b00anx92e4bfv1tmnn8h9cbn1@marketplace.amazon.comusca",
  "seller_id": "A197RN4AZ31SIL",
  "date": "2024-01-05",
  "customer": "bfv1tmnn8h9cbn1@marketplace.amazon.com",
  "state_or_region": "ca",
  "country": "us",
  "asin": "B00ANX92E4",
  "cltv": "19.99",
  "operation": "UPSERT",
  "updated_at": "2024-01-31 10:02:44.595054 UTC"
}, {
  "id": "a197rn4az31sil2024-01-05b00imisulsysth8d99b5l6wpb@marketplace.amazon.comusfl",
  "seller_id": "A197RN4AZ31SIL",
  "date": "2024-01-05",
  "customer": "ysth8d99b5l6wpb@marketplace.amazon.com",
  "state_or_region": "fl",
  "country": "us",
  "asin": "B00IMISULS",
  "cltv": "49.99",
  "operation": "UPSERT",
  "updated_at": "2024-01-31 10:02:44.595401 UTC"
}, {
  "id": "a197rn4az31sil2024-01-06b074hrgfm8qtsl0s8csznvf82@marketplace.amazon.comusny",
  "seller_id": "A197RN4AZ31SIL",
  "date": "2024-01-06",
  "customer": "qtsl0s8csznvf82@marketplace.amazon.com",
  "state_or_region": "ny",
  "country": "us",
  "asin": "B074HRGFM8",
  "cltv": "49.99",
  "operation": "UPSERT",
  "updated_at": "2024-01-31 10:03:49.525387 UTC"
}, {
  "id": "a197rn4az31sil2024-01-06b0b4fb2g1t8r55rcm3qtw8wj9@marketplace.amazon.comusla",
  "seller_id": "A197RN4AZ31SIL",
  "date": "2024-01-06",
  "customer": "8r55rcm3qtw8wj9@marketplace.amazon.com",
  "state_or_region": "la",
  "country": "us",
  "asin": "B0B4FB2G1T",
  "cltv": "51.99",
  "operation": "UPSERT",
  "updated_at": "2024-01-31 10:03:49.533571 UTC"
}]
`

func main() {
	var data []map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(jsonData)).Decode(&data); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	fmt.Printf("data\n%#v\n\n", data)

	htmlTable := generateHTMLTable(data)
	htmlOrderTable := generateHTMLTableInOrder(data)
	if err := takeScreenshot(htmlTable, "file/table.png"); err != nil {
		log.Fatalf("Error taking screenshot: %v", err)
	}

	file, err := os.Create("table.html")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if _, err := file.Write([]byte(htmlOrderTable)); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Screenshot saved as table.png")
	if err := takeScreenshot(htmlOrderTable, "file/order_table.png"); err != nil {
		log.Fatalf("Error taking screenshot: %v", err)
	}
	fmt.Println("Screenshot saved as table.png")
}

func generateHTMLTable(data []map[string]interface{}) string {
	var table strings.Builder
	table.WriteString(`<table style="" border="2"><thead><tr>`)
	for key := range data[0] {
		table.WriteString(fmt.Sprintf("<th>%s</th>", key))
	}
	table.WriteString("</tr></thead><tbody>")
	for _, item := range data {
		table.WriteString("<tr>")
		for _, value := range item {
			table.WriteString(fmt.Sprintf("<td>%v</td>", value))
		}
		table.WriteString("</tr>")
	}
	table.WriteString("</tbody></table>")
	return table.String()
}

func generateHTMLTableInOrder(data []map[string]interface{}) string {
	keyOrder := []string{"id", "seller_id", "date", "customer", "state_or_region", "country", "asin", "cltv", "operation", "updated_at"}

	var table strings.Builder
	table.WriteString(`<table border="2"><thead><tr>`)
	for _, key := range keyOrder {
		table.WriteString(fmt.Sprintf("<th>%s</th>", key))
	}
	table.WriteString("</tr></thead><tbody>")
	for _, item := range data {
		table.WriteString("<tr>")
		for _, key := range keyOrder {
			table.WriteString(fmt.Sprintf("<td>%v</td>", item[key]))
		}
		table.WriteString("</tr>")
	}
	table.WriteString("</tbody></table>")
	return table.String()
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
