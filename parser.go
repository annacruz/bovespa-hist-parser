package main

type Parser struct {
	trading_day_interval Interval
}

type Interval struct {
	start int
	stop  int
}

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

func NewParser() *Parser {
	var parser = &Parser{
		trading_day_interval: Interval{
			start: 3,
			stop:  7,
		},
	}
	return parser

}

func main() {
}
