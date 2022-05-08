package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/rakhmadbudiono/mutual-fund-aggregator/internal/pasardana"
	"github.com/rakhmadbudiono/mutual-fund-aggregator/pkg/mutualfund"
)

func main() {
	pc := &pasardana.PasardanaClient{
		Client:  http.DefaultClient,
		BaseURL: "https://www.pasardana.id/api",
	}

	mfs := ExtractMutualFunds(pc)

	path := "mutual-funds.csv"
	SaveToCSV(mfs, path)
}

func ExtractMutualFunds(client *pasardana.PasardanaClient) []mutualfund.IMutualFund {
	ids, err := client.GetAllMutualFundIDs()
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan mutualfund.IMutualFund)
	mfs := []mutualfund.IMutualFund{}
	for _, id := range ids {
		go func(id int) {
			mf, err := client.GetMutualFundByID(id)
			if err != nil {
				log.Print(err)
			}

			ch <- mf
		}(id)
	}

	for range ids {
		mfs = append(mfs, <-ch)
	}

	return mfs
}

func SaveToCSV(mfs []mutualfund.IMutualFund, path string) {
	records := [][]string{}

	headers := []string{
		"Name",
		"Beta",
		"Jensen Alpha",
		"Sharpe Ratio",
		"Treynor Ratio",
		"Tracking Error",
		"AUM",
	}

	records = append(records, headers)

	for _, mf := range mfs {
		if mf == nil {
			log.Println("saveToCSV: mutual fund not found")
			continue
		}

		records = append(records, mf.ToArrayString())
	}

	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		log.Fatalln("saveToCSV: failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records)
	if err != nil {
		log.Fatal(err)
	}
}
