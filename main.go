package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Stock struct {
	company string
	price   string
	change  string
}

func main() {
	
	stockTickers := []string{
		"AAPL", "MSFT", "GOOGL", "AMZN", "TSLA",
		"FB", "NVDA", "JPM", "JNJ", "V",
	}
	stocks := []Stock{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting : ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error: ", err)
	})

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {

		stock := Stock{}

		stock.company = e.ChildText("h1")
		fmt.Println("Company: ", stock.company)

		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Price: ", stock.price)

		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("Change: ", stock.change)

		stocks = append(stocks, stock)
	})

	c.Wait()

	for _, t := range stockTickers {
		c.Visit("https://finance.yahoo.com/quote/" + t + "/")
	}

	fmt.Println(stocks)

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalln("Cannot create file ... ")
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{
		"Company",
		"Price",
		"Change",
	}
	writer.Write(headers) // columns

	for _, t := range stocks {
		record := []string{
			t.company,
			t.price,
			t.change,
		}
		writer.Write(record)
	}

	defer writer.Flush()
}
