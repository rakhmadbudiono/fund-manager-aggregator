package pasardana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type searchMutualFundResponse struct {
	ID int `json:"Id"`
}

func GetAllMutualFundIDs() ([]int, error) {
	res, err := Client.Get(
		fmt.Sprintf("%s/FundSearchResult/GetAll", BaseURL),
	)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return byteToMutualFundIDs(body)
}

func byteToMutualFundIDs(bytes []byte) ([]int, error) {
	mfResp := []searchMutualFundResponse{}

	if err := json.Unmarshal(bytes, &mfResp); err != nil {
		return nil, err
	}

	ids := []int{}
	for _, mf := range mfResp {
		ids = append(ids, mf.ID)
	}

	return ids, nil
}
