package main

type Parser struct {
	trading_day_interval             Interval
	stock_code_interval              Interval
	price_start_trading_day_interval Interval
	max_price_interval               Interval
	min_price_interval               Interval
	average_price_interval           Interval
	last_price_interval              Interval
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
			start: 2,
			stop:  10,
		},
		stock_code_interval: Interval{
			start: 12,
			stop:  24,
		},
		price_start_trading_day_interval: Interval{
			start: 56,
			stop:  69,
		},
		max_price_interval: Interval{
			start: 69,
			stop:  82,
		},
		min_price_interval: Interval{
			start: 82,
			stop:  95,
		},
		average_price_interval: Interval{
			start: 95,
			stop:  108,
		},
		last_price_interval: Interval{
			start: 108,
			stop:  121,
		},
	}
	return parser

}

func main() {
}
