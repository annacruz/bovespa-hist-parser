package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Describe("Parsing line", func() {
		var (
			parser   *Parser
			cotation *Cotation
			err      error
			line     string
		)

		BeforeEach(func() {
			line = "012016110102AALR3       010ALLIAR      ON      NM   R$  000000000180600000000018380000000001716000000000177400000000017900000000001765000000000179002072000000000000996200000000001767698100000000000000009999123100000010000000000000BRAALRACNOR6100"
			parser = NewParser()
			cotation, err = parser.Do(line)
		})
		Context("valid line", func() {
			It("retrieve cotation stock code throught interval", func() {
				Expect(cotation.ticker).To(Equal("AALR3"))
			})
			It("retrieve cotation trading day throught interval", func() {
				Expect(cotation.trading_day).To(Equal("20161101"))
			})
			It("retieves cotation price on start throught interval", func() {
				Expect(cotation.price_on_start).To(Equal("18.06"))
			})
		})
	})
	Describe("Converte decimal numbers", func() {
		Context("valid number", func() {
			It("get number in plain text and convert it to an decimal with precision two", func() {
				Expect(ConvertNumber("0000000001806")).To(Equal("18.06"))
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
			Expect(parser.trading_day_interval.start).To(Equal(2))
		})
		It("has trading day interval stop", func() {
			Expect(parser.trading_day_interval.stop).To(Equal(10))
		})
		It("has stock code interval start and stop", func() {
			Expect(parser.ticker_interval.start).To(Equal(12))
			Expect(parser.ticker_interval.stop).To(Equal(24))
		})
		It("has price on start of trading day interval start and stop", func() {
			Expect(parser.price_start_trading_day_interval.start).To(Equal(56))
			Expect(parser.price_start_trading_day_interval.stop).To(Equal(69))
		})
		It("has maximum price interval start and stop", func() {
			Expect(parser.max_price_interval.start).To(Equal(69))
			Expect(parser.max_price_interval.stop).To(Equal(82))
		})
		It("has minimum price interval start and stop", func() {
			Expect(parser.min_price_interval.start).To(Equal(82))
			Expect(parser.min_price_interval.stop).To(Equal(95))
		})
		It("has average price interval start and stop", func() {
			Expect(parser.average_price_interval.start).To(Equal(95))
			Expect(parser.average_price_interval.stop).To(Equal(108))
		})
		It("has last price interval start and stop", func() {
			Expect(parser.last_price_interval.start).To(Equal(108))
			Expect(parser.last_price_interval.stop).To(Equal(121))
		})
	})
})
