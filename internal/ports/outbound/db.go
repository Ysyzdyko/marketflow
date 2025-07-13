package outbound

type DbPort interface {
	GetTokenRepository
	PostTokenRepository
	HighestTokenRepository
	LowestTokenRepository
	AverageTokenRepository
}

type GetTokenRepository interface {
	ListPriceBySymbol()
	ListPriceExchangeBySymbol()
}

type PostTokenRepository interface {
	AddPrice()
	UpdatePRice()
}

type HighestTokenRepository interface {
	ListHighestPriceBySymbol()
	ListHighestPriceExcSym()
	ListHighestPriceBySymbolDuration()
	ListHighestPriceBySymbExcDuration()
}

type LowestTokenRepository interface {
	ListLowestPriceBySymbol()
	ListLowestPriceExcSym()
	ListLowestPriceBySymbolDuration()
	ListLowestPriceBySymbExcDuration()
}

type AverageTokenRepository interface {
	ListAvgPriceBySymbol()
	ListAvgPriceExcSym()
	ListAvgPriceBySymbExcDuration()
}
