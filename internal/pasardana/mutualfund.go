package pasardana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/rakhmadbudiono/mutual-fund-aggregator/pkg/mutualfund"
)

type snapshotMutualFundResponse struct {
	Name         string                        `json:"Name"`
	FundType     mutualfund.MutualFundType     `json:"Type"`
	Performances []snapshotPerformanceResponse `json:"Performances"`
	AUMs         []snapshotAUMResponse         `json:"AssetUnderManagements"`
}

type snapshotPerformanceResponse struct {
	Beta          float32 `json:"Beta"`
	JensenAlpha   float32 `json:"JensenAlpha"`
	SharpeRatio   float32 `json:"SharpeRatio"`
	TreynorRatio  float32 `json:"TreynorRatio"`
	TrackingError float32 `json:"TrackingError"`
}

type snapshotAUMResponse struct {
	Value float64 `json:"Value"`
}

func CountMutualFunds() (int, error) {
	res, err := Client.Get(
		fmt.Sprintf("%s/FundSearchResult/GetCount", BaseURL),
	)
	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	return byteToInt(body)
}

func byteToInt(bytes []byte) (int, error) {
	text := string(bytes)

	number, err := strconv.Atoi(text)
	if err != nil {
		return 0, nil
	}

	return number, nil
}

func GetMutualFundByID(ID int) (mutualfund.IMutualFund, error) {
	res, err := Client.Get(
		fmt.Sprintf("%s/FundService/GetSnapshot/%d", BaseURL, ID),
	)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return byteToMutualFund(body)
}

func byteToMutualFund(bytes []byte) (mutualfund.IMutualFund, error) {
	mfResp := &snapshotMutualFundResponse{}

	if err := json.Unmarshal(bytes, mfResp); err != nil {
		return nil, err
	}

	perfsLength := len(mfResp.Performances)
	if perfsLength == 0 {
		return nil, fmt.Errorf("pasardana.byteToMutualFund: empty performance")
	}

	lastPerf := mfResp.Performances[perfsLength-1]
	indicator := mutualfund.Indicator{
		Beta:          lastPerf.Beta,
		JensenAlpha:   lastPerf.JensenAlpha,
		SharpeRatio:   lastPerf.SharpeRatio,
		TreynorRatio:  lastPerf.TreynorRatio,
		TrackingError: lastPerf.TrackingError,
	}

	aumsLength := len(mfResp.AUMs)
	if aumsLength == 0 {
		return nil, fmt.Errorf("pasardana.byteToMutualFund: empty AUM")
	}

	lastAUM := mfResp.AUMs[aumsLength-1]
	aum := lastAUM.Value

	return mutualfund.NewMutualFund(
		mfResp.Name,
		mfResp.FundType,
		mutualfund.WithIndicator(indicator),
		mutualfund.WithAUM(aum),
	), nil
}
