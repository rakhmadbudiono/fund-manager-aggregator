package pasardana_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/rakhmadbudiono/mutual-fund-aggregator/internal/pasardana"
	"github.com/rakhmadbudiono/mutual-fund-aggregator/pkg/mutualfund"

	"github.com/stretchr/testify/assert"
)

func TestCountMutualFunds(t *testing.T) {
	expectedCount := 69420

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response := []byte(strconv.Itoa(expectedCount))

		rw.Write(response)
	}))
	defer server.Close()

	pasardana.BaseURL = server.URL
	count, err := pasardana.CountMutualFunds()

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedCount, count)
}

func TestGetMutualFundByID(t *testing.T) {
	performance := `{"Beta":1, "JensenAlpha":0, "SharpeRatio":1, "TreynorRatio":0, "TrackingError":0}`
	response := fmt.Sprintf(`{"Name":"Gopher Capital", "Type":4, "AssetUnderManagements":[{"Value":69420}], "Performances": [%s]}`, performance)
	mock := []byte(response)

	expected := mutualfund.NewMutualFund(
		"Gopher Capital",
		4, // Other
		mutualfund.WithIndicator(mutualfund.Indicator{
			Beta:          1,
			JensenAlpha:   0,
			SharpeRatio:   1,
			TreynorRatio:  0,
			TrackingError: 0,
		},
		),
		mutualfund.WithAUM(69420),
	)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(mock)
	}))
	defer server.Close()

	pasardana.BaseURL = server.URL
	mf, err := pasardana.GetMutualFundByID(1)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, mf)
}
