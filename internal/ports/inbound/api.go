package inbound

type APIPorts interface {
	PricesPort
	AppMode
	PriceExchangePort
}

type PricesPort interface {
	GetLatestPrice()
}

type PriceExchangePort interface {
}

type AppMode interface {
	SetMode(live bool)
}
