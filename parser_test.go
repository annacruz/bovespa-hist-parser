package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var parser Parser

func do_parse_or_fatal(t *testing.T, line string) *Cotation {
	cotation, err := parser.Do(line)
	if err != nil {
		t.Errorf("Error in parse")
	}
	return cotation
}

func Test_parse_return_stock_code(t *testing.T) {
	line := "012016110102AALR3       010ALLIAR      ON      NM   R$  000000000180600000000018380000000001716000000000177400000000017900000000001765000000000179002072000000000000996200000000001767698100000000000000009999123100000010000000000000BRAALRACNOR6100"

	cotation := do_parse_or_fatal(t, line)

	assert.Equal(t, "AALR3", cotation.stock_code, "it should be equal")
}

func Test_parse_return_trading_day(t *testing.T) {
	line := "012016110102AALR3       010ALLIAR      ON      NM   R$  000000000180600000000018380000000001716000000000177400000000017900000000001765000000000179002072000000000000996200000000001767698100000000000000009999123100000010000000000000BRAALRACNOR6100"

	cotation := do_parse_or_fatal(t, line)

	assert.Equal(t, "20161101", cotation.trading_day, "it should be equal")
}
func Test_parse_return_price_on_trade_start(t *testing.T) {
	line := "012016110102AALR3       010ALLIAR      ON      NM   R$  000000000180600000000018380000000001716000000000177400000000017900000000001765000000000179002072000000000000996200000000001767698100000000000000009999123100000010000000000000BRAALRACNOR6100"

	cotation := do_parse_or_fatal(t, line)

	assert.Equal(t, "180,60", cotation.price_on_start, "it should be equal")
}
