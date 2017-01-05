package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Describe("Parsing line", func() {
		var (
			parser   Parser
			cotation *Cotation
			err      error
			line     string
		)

		BeforeEach(func() {
			line = "012016110102AALR3       010ALLIAR      ON      NM   R$  000000000180600000000018380000000001716000000000177400000000017900000000001765000000000179002072000000000000996200000000001767698100000000000000009999123100000010000000000000BRAALRACNOR6100"
			cotation, err = parser.Do(line)
		})
		Context("valid line", func() {
			It("retrieve cotation stock code", func() {
				Expect(cotation.stock_code).To(Equal("AALR3"))
			})
			It("retrieve cotation trading day", func() {
				Expect(cotation.trading_day).To(Equal("20161101"))
			})
			It("retieves cotation price on start", func() {
				Expect(cotation.price_on_start).To(Equal("180,60"))
			})
		})
	})
	Describe("Intervals", func() {
		var (
			parser *Parser
		)
		BeforeEach(func() {
			parser = NewParser()
		})
		It("has trading day interval start", func() {
			Expect(parser.trading_day_interval.start).To(Equal(3))
		})
		It("has trading day interval stop", func() {
			Expect(parser.trading_day_interval.stop).To(Equal(7))
		})
	})
})
