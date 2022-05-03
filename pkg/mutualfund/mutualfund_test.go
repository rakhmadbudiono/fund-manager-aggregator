package mutualfund_test

import (
	"fmt"
	"testing"

	"github.com/rakhmadbudiono/mutual-fund-aggregator/pkg/mutualfund"

	"github.com/stretchr/testify/assert"
)

var (
	name          string  = "Gopher Capital"
	fundtype      string  = "YOLO"
	aum           float64 = 69420
	beta          float32 = 1
	alpha         float32 = 0
	sharpe        float32 = 1
	treynor       float32 = 0
	trackingError float32 = 0
)

func TestNewMutualFund(t *testing.T) {
	mf := mutualfund.NewMutualFund(name, fundtype)

	assert.Equal(
		t,
		name,
		mf.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, mf.GetName()),
	)
}

func TestNewMutualFundWithIndicator(t *testing.T) {
	indicator := mutualfund.Indicator{
		Beta:          beta,
		JensenAlpha:   alpha,
		SharpeRatio:   sharpe,
		TreynorRatio:  treynor,
		TrackingError: trackingError,
	}

	mf := mutualfund.NewMutualFund(
		name,
		fundtype,
		mutualfund.WithIndicator(indicator),
	)

	assert.Equal(
		t,
		name,
		mf.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, mf.GetName()),
	)

	assert.Equal(
		t,
		indicator,
		mf.GetIndicator(),
		fmt.Sprintf("Expect indicator to be %v, but got %v", indicator, mf.GetIndicator()),
	)
}

func TestNewMutualFundWithAUM(t *testing.T) {
	mf := mutualfund.NewMutualFund(
		name,
		fundtype,
		mutualfund.WithAUM(aum),
	)

	assert.Equal(
		t,
		name,
		mf.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, mf.GetName()),
	)

	assert.Equal(
		t,
		aum,
		mf.GetAUM(),
		fmt.Sprintf("Expect aum to be %v, but got %v", aum, mf.GetAUM()),
	)
}
