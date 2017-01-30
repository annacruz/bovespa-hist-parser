package main

import (
	//	"fmt"
	//	"github.com/ericlagergren/decimal"
	. "strings"
)

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
	stock_code     string `json:"stock_code`
	trading_day    string `json:"trading_day`
	price_on_start string `json:"price_on_start`
}

func (parser *Parser) Do(line string) (*Cotation, error) {
	cotation := &Cotation{
		stock_code:  TrimSpace(line[parser.stock_code_interval.start:parser.stock_code_interval.stop]),
		trading_day: TrimSpace(line[parser.trading_day_interval.start:parser.trading_day_interval.stop]),
		//		price_on_start: decimal.New(TrimLeft(line[parser.price_start_trading_day_interval.start:parser.price_start_trading_day_interval.stop], "0"), 2),
	}
	return cotation, nil
}

func ConvertNumber(number string) string {
	return ""
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
