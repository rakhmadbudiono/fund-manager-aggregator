package fundmanager_test

import (
	"fmt"
	"testing"

	"github.com/rakhmadbudiono/fund-manager-aggregator/pkg/fundmanager"

	"github.com/stretchr/testify/assert"
)

var (
	name          string                = "Gopher Capital"
	fundtype      string                = "YOLO"
	aum           float64               = 69420
	risk          fundmanager.RiskLevel = fundmanager.RiskLow
	beta          float32               = 1
	alpha         float32               = 0
	sharpe        float32               = 1
	treynor       float32               = 0
	trackingError float32               = 0
)

func TestNewFundManager(t *testing.T) {
	fm := fundmanager.NewFundManager(name, fundtype)

	assert.Equal(
		t,
		name,
		fm.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, fm.GetName()),
	)
}

func TestNewFundManagerWithIndicator(t *testing.T) {
	indicator := fundmanager.Indicator{
		Beta:          beta,
		JensenAlpha:   alpha,
		SharpeRatio:   sharpe,
		TreynorRatio:  treynor,
		TrackingError: trackingError,
	}

	fm := fundmanager.NewFundManager(
		name,
		fundtype,
		fundmanager.WithIndicator(indicator),
	)

	assert.Equal(
		t,
		name,
		fm.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, fm.GetName()),
	)

	assert.Equal(
		t,
		indicator,
		fm.GetIndicator(),
		fmt.Sprintf("Expect indicator to be %v, but got %v", indicator, fm.GetIndicator()),
	)
}

func TestNewFundManagerWithDetail(t *testing.T) {
	detail := fundmanager.Detail{
		AUM:  aum,
		Risk: risk,
	}

	fm := fundmanager.NewFundManager(
		name,
		fundtype,
		fundmanager.WithDetail(detail),
	)

	assert.Equal(
		t,
		name,
		fm.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, fm.GetName()),
	)

	assert.Equal(
		t,
		detail,
		fm.GetDetail(),
		fmt.Sprintf("Expect detail to be %v, but got %v", detail, fm.GetDetail()),
	)
}
