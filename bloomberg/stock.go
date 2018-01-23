package bloomberg

type Stock struct {
	CompanyID            string
	Exchange             string
	CurrentPrice         string
	AbsoluteChange       string
	PercentChange        string
	MarketOpen           string
	PrevMarketClose      string
	Volume               string
	MarketCap            string
	PERatio              string
	BloombergPERatio     string
	BloombergPEGRatio    string
	SharesOutstanding    string
	PriceToBookRatio     string
	PriceToSalesRatio    string
	OneYearReturnPct     string
	MonthAvgVolume       string
	EPS                  string
	BloombergEPSYear     string
	Dividend             string
	LastDividendReported string
	Address              string
	PhoneNumber          string
	QrtRevenue           string
	QrtNetIncome         string
	QrtProfitMargin      string
	QrtTotalAssets       string
	QrtLiabilities       string
	QrtDebtToAssetsRatio string
	QrtOperating         string
	QrtInvesting         string
	QrtFinancing         string
}
