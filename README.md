# Stock Price Scraper

This is a simple Go program that scrapes stock price and change data from Yahoo Finance using the Colly web scraping framework. It collects information for a list of predefined stock tickers, extracts relevant data from the HTML structure of the Yahoo Finance pages, and saves the data to a CSV file.

## How It Works

### Dependencies

This program uses the following dependencies:

- **Colly:** A highly efficient and flexible web scraping framework for Golang.
- **Encoding/CSV:** A package for reading and writing CSV files.

### Workflow

1. **Define Stock Tickers:** The program starts by defining a list of stock tickers to scrape data for.

2. **Initialize Collector:** It creates a new Colly collector instance to handle HTTP requests and HTML parsing.

3. **Request and Error Handlers:**
   - **OnRequest:** Prints the URL being visited.
   - **OnError:** Logs any errors that occur during the scraping process.

4. **HTML Parsing:**
   - **OnHTML:** Defines a handler to extract stock information from the HTML structure of the Yahoo Finance page. It extracts the company name, stock price, and price change percentage.

5. **Scraping:**
   - The program iterates over each stock ticker and visits the corresponding Yahoo Finance page.
   - For each page visited, it extracts the relevant stock information using the defined HTML parsing handler.

6. **Data Export:**
   - After scraping data for all the stocks, it prints the collected data to the console.
   - It then writes the collected data to a CSV file named `stocks.csv`, with columns for company name, price, and change.

## Running the Program

To run the program, ensure you have Go installed on your system. Then, clone this repository and execute the following command:

```bash
go run main.go
