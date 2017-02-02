package main

import (
	"strconv"
	"strings"
)

type Parser struct {
	trading_day_interval             Interval
	ticker_interval                  Interval
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
	ticker         string `json:"ticker`
	trading_day    string `json:"trading_day`
	price_on_start string `json:"price_on_start`
}

func (parser *Parser) Do(line string) (*Cotation, error) {
	cotation := &Cotation{
		ticker:         strings.TrimSpace(line[parser.ticker_interval.start:parser.ticker_interval.stop]),
		trading_day:    strings.TrimSpace(line[parser.trading_day_interval.start:parser.trading_day_interval.stop]),
		price_on_start: ConvertNumber(line[parser.price_start_trading_day_interval.start:parser.price_start_trading_day_interval.stop]),
	}
	return cotation, nil
}

func ConvertNumber(number string) string {
	unformatted_number, _ := strconv.ParseFloat(number, 64)
	return strconv.FormatFloat(unformatted_number/100.0, 'f', 2, 64)
}

func NewParser() *Parser {
	var parser = &Parser{
		trading_day_interval: Interval{
			start: 2,
			stop:  10,
		},
		ticker_interval: Interval{
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
