package main_test

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	main "github.com/rakhmadbudiono/mutual-fund-aggregator/cmd/aggregator"
	"github.com/rakhmadbudiono/mutual-fund-aggregator/internal/pasardana"
	"github.com/rakhmadbudiono/mutual-fund-aggregator/pkg/mutualfund"

	"github.com/stretchr/testify/assert"
)

func TestExtractMutualFunds(t *testing.T) {
	expected := []mutualfund.IMutualFund{
		mutualfund.NewMutualFund(
			"Gopher Capital",
			4, // other,
			mutualfund.WithIndicator(mutualfund.Indicator{
				Beta:          1,
				JensenAlpha:   0,
				SharpeRatio:   1,
				TreynorRatio:  0,
				TrackingError: 0,
			}),
			mutualfund.WithAUM(69420),
		),
	}

	client := setupPasardanaTestClient()
	mfs := main.ExtractMutualFunds(client)

	assert.Equal(t, expected, mfs)
}

func setupPasardanaTestClient() *pasardana.PasardanaClient {
	firstResponse := []byte(`[{"Id":1}]`)

	performance := `{"Beta":1, "JensenAlpha":0, "SharpeRatio":1, "TreynorRatio":0, "TrackingError":0}`
	json := fmt.Sprintf(`{"Name":"Gopher Capital", "Type":4, "AssetUnderManagements":[{"Value":69420}], "Performances": [%s]}`, performance)
	secondResponse := []byte(json)

	reqNumber := 0
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		reqNumber++

		if reqNumber == 1 {
			rw.Write(firstResponse)
			return
		}

		rw.Write(secondResponse)
	}))

	return &pasardana.PasardanaClient{
		Client:  http.DefaultClient,
		BaseURL: server.URL,
	}
}

func TestSaveToCSV(t *testing.T) {
	expected := [][]string{
		[]string{
			"Name",
			"Beta",
			"Jensen Alpha",
			"Sharpe Ratio",
			"Treynor Ratio",
			"Tracking Error",
			"AUM",
		},
		[]string{
			"Gopher Capital",
			"1.000000",
			"0.000000",
			"1.000000",
			"0.000000",
			"0.000000",
			"69420.000000",
		},
	}

	mfs := []mutualfund.IMutualFund{
		mutualfund.NewMutualFund(
			"Gopher Capital",
			4, // other,
			mutualfund.WithIndicator(mutualfund.Indicator{
				Beta:          1,
				JensenAlpha:   0,
				SharpeRatio:   1,
				TreynorRatio:  0,
				TrackingError: 0,
			}),
			mutualfund.WithAUM(69420),
		),
	}

	main.SaveToCSV(mfs, "test.csv")
	assertCSV(t, "test.csv", expected)
}

func assertCSV(t *testing.T, path string, expectedRows [][]string) {
	rows := [][]string{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		rows = append(rows, record)
	}

	assert.Equal(t, expectedRows, rows)
}
