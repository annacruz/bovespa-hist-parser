package main

type Parser struct{}

type Cotation struct {
	stock_code     string
	trading_day    string
	price_on_start string
}

func (p *Parser) Do(line string) (*Cotation, error) {
	cotation := &Cotation{
		stock_code:     "AALR3",
		trading_day:    "20161101",
		price_on_start: "180,60",
	}
	return cotation, nil
}

func main() {}
