package mutualfund

type MixedMutualFund struct {
	MutualFund
}

func newMixedMutualFund(mf MutualFund) IMutualFund {
	return &MixedMutualFund{
		MutualFund: mf,
	}
}
