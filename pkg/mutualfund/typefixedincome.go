package mutualfund

type FixedIncomeMutualFund struct {
	MutualFund
}

func newFixedIncomeMutualFund(mf MutualFund) IMutualFund {
	return &FixedIncomeMutualFund{
		MutualFund: mf,
	}
}
