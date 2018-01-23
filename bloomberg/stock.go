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

func NewImportedStock(fields []string) *Stock {
	return &Stock{
		CompanyID:            fields[0],
		Exchange:             fields[1],
		CurrentPrice:         fields[2],
		AbsoluteChange:       fields[3],
		PercentChange:        fields[4],
		MarketOpen:           fields[5],
		PrevMarketClose:      fields[6],
		Volume:               fields[7],
		MarketCap:            fields[8],
		PERatio:              fields[9],
		BloombergPERatio:     fields[10],
		BloombergPEGRatio:    fields[11],
		SharesOutstanding:    fields[12],
		PriceToBookRatio:     fields[13],
		PriceToSalesRatio:    fields[14],
		OneYearReturnPct:     fields[15],
		MonthAvgVolume:       fields[16],
		EPS:                  fields[17],
		BloombergEPSYear:     fields[18],
		Dividend:             fields[19],
		LastDividendReported: fields[20],
		Address:              fields[21],
		PhoneNumber:          fields[22],
		QrtRevenue:           fields[23],
		QrtNetIncome:         fields[24],
		QrtProfitMargin:      fields[25],
		QrtTotalAssets:       fields[26],
		QrtLiabilities:       fields[27],
		QrtDebtToAssetsRatio: fields[28],
		QrtOperating:         fields[29],
		QrtInvesting:         fields[30],
		QrtFinancing:         fields[31],
	}
}
