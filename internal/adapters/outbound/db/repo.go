package db

import "database/sql"

type Repo struct {
	Conn *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{Conn: db}
}

func (r *Repo) ListPriceBySymbol()                 {}
func (r *Repo) ListPriceExchangeBySymbol()         {}
func (r *Repo) AddPrice()                          {}
func (r *Repo) UpdatePRice()                       {}
func (r *Repo) ListHighestPriceBySymbol()          {}
func (r *Repo) ListHighestPriceExcSym()            {}
func (r *Repo) ListHighestPriceBySymbolDuration()  {}
func (r *Repo) ListHighestPriceBySymbExcDuration() {}
func (r *Repo) ListLowestPriceBySymbol()           {}
func (r *Repo) ListLowestPriceExcSym()             {}
func (r *Repo) ListLowestPriceBySymbolDuration()   {}
func (r *Repo) ListLowestPriceBySymbExcDuration()  {}
func (r *Repo) ListAvgPriceBySymbol()              {}
func (r *Repo) ListAvgPriceExcSym()                {}
func (r *Repo) ListAvgPriceBySymbExcDuration()     {}
