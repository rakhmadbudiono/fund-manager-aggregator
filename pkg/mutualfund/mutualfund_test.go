package mutualfund_test

import (
	"fmt"
	"testing"

	"github.com/rakhmadbudiono/fund-manager-aggregator/pkg/mutualfund"

	"github.com/stretchr/testify/assert"
)

var (
	name          string               = "Gopher Capital"
	fundtype      string               = "YOLO"
	aum           float64              = 69420
	risk          mutualfund.RiskLevel = mutualfund.RiskLow
	beta          float32              = 1
	alpha         float32              = 0
	sharpe        float32              = 1
	treynor       float32              = 0
	trackingError float32              = 0
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

func TestNewMutualFundWithDetail(t *testing.T) {
	detail := mutualfund.Detail{
		AUM:  aum,
		Risk: risk,
	}

	mf := mutualfund.NewMutualFund(
		name,
		fundtype,
		mutualfund.WithDetail(detail),
	)

	assert.Equal(
		t,
		name,
		mf.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, mf.GetName()),
	)

	assert.Equal(
		t,
		detail,
		mf.GetDetail(),
		fmt.Sprintf("Expect detail to be %v, but got %v", detail, mf.GetDetail()),
	)
}
