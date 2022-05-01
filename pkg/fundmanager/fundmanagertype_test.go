package fundmanager_test

import (
	"fmt"
	"testing"

	"github.com/rakhmadbudiono/fund-manager-aggregator/pkg/fundmanager"

	"github.com/stretchr/testify/assert"
)

var (
	fixedincome string = "Pendapatan Tetap"
	mixed       string = "Campuran"
	moneymarket string = "Pasar Uang"
	stock       string = "Saham"
)

func TestNewFixedIncomeFundManager(t *testing.T) {
	fm := fundmanager.NewFundManager(name, fixedincome)

	assert.Equal(
		t,
		name,
		fm.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, fm.GetName()),
	)
}

func TestNewMixedFundManager(t *testing.T) {
	fm := fundmanager.NewFundManager(name, mixed)

	assert.Equal(
		t,
		name,
		fm.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, fm.GetName()),
	)
}

func TestNewMoneyMarketFundManager(t *testing.T) {
	fm := fundmanager.NewFundManager(name, moneymarket)

	assert.Equal(
		t,
		name,
		fm.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, fm.GetName()),
	)
}

func TestNewStockFundManager(t *testing.T) {
	fm := fundmanager.NewFundManager(name, stock)

	assert.Equal(
		t,
		name,
		fm.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, fm.GetName()),
	)
}
