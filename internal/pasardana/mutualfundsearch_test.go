package pasardana_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakhmadbudiono/mutual-fund-aggregator/internal/pasardana"

	"github.com/stretchr/testify/assert"
)

func TestGetAllMutualFunds(t *testing.T) {
	response := `[{"Id":1}, {"Id":2}]`
	mock := []byte(response)

	expected := []int{1, 2}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(mock)
	}))
	defer server.Close()

	pasardana.BaseURL = server.URL
	mf, err := pasardana.GetAllMutualFundIDs()

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, mf)
}
