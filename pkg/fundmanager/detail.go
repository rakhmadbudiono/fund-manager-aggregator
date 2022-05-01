package fundmanager

type RiskLevel int

const (
	RiskLow RiskLevel = iota
	RiskMedium
	RiskHigh
)

type Detail struct {
	AUM  float64
	Risk RiskLevel
}
