package bitmax

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	apiKey    = ""
	secretKey = ""
)

var _ = Describe("BitmaxGo", func() {
	var client *Client

	BeforeEach(func() {
		client = CreateClient(apiKey, secretKey)

		userInfo, err := client.UserInfo()
		if err != nil {
			log.Panicln(err)
		}

		client.AccountGroup = userInfo.AccountGroup
	})

	Describe("UserInfo", func() {
		Context("with valid keys", func() {
			It("returns accountGroup", func() {
				userInfo, err := client.UserInfo()
				if err != nil {
					log.Panicln(err)
				}

				Expect(userInfo.AccountGroup > 0).To(BeTrue())
			})
		})

		Context("wtih invalid keys", func() {
			It("returns error", func() {
				c := CreateClient("", "")

				_, err := c.UserInfo()

				Expect(err.Error()).To(Equal("ApiKeyFailure"))
			})
		})
	})

	Describe("Balances", func() {
		It("returns balances", func() {
			balances, err := client.Balances()
			if err != nil {
				log.Panicln(err)
			}

			Expect(len(balances) > 0).To(BeTrue())
		})
	})

	Describe("Balance", func() {
		var symbol string

		Context("with existing symbol", func() {
			BeforeEach(func() {
				symbol = "BTC"
			})

			It("returns balance of symbol", func() {
				balance, err := client.Balance(symbol)
				if err != nil {
					log.Panicln(err)
				}

				Expect(balance.AssetCode).To(Equal(symbol))
			})
		})

		Context("with not existing symbol", func() {
			BeforeEach(func() {
				symbol = "Wrong"
			})

			It("returns error", func() {
				_, err := client.Balance(symbol)

				Expect(err.Error()).To(Equal("Invalid Symbol"))
			})
		})
	})

	Describe("MarketQuote", func() {
		symbol := "ETH-BTC"

		It("returns level1 orderbook data", func() {
			quote, err := client.Quote(symbol)
			if err != nil {
				log.Panicln(err)
			}

			Expect(quote.Symbol).To(Equal("ETH/BTC"))
		})
	})

	Describe("Order", func() {
		// Context("with valid inputs", func() {
		// 	It("buy a eth", func() {
		// 		orderResponse, err := client.Order("buy", "ETH/BTC", "0.018", "0.05")
		// 		if err != nil {
		// 			log.Panicln(err)
		// 		}

		// 		Expect(orderResponse.COID).To(HaveLen(32))
		// 	})
		// })

		Context("when price is too low", func() {
			It("returns error", func() {
				_, err := client.Order("buy", "ETH/BTC", "0.0001", "0.01")

				Expect(err.Error()).
					To(Equal("Price is too low from market price."))
			})
		})

		Context("when amount is too small", func() {
			It("returns error", func() {
				_, err := client.Order("buy", "ETH/BTC", "0.0190", "0.001")

				Expect(err.Error()).To(Equal("Notional is too small."))
			})
		})
	})

	Describe("Products", func() {
		It("returns products", func() {
			products, err := client.Products()
			if err != nil {
				log.Panicln(err)
			}

			Expect(len(products) > 0).To(BeTrue())
		})
	})
	// Describe("Orders", func() {
	// 	It("returns open orders", func() {
	// 		orders, err := client.Orders()
	// 		if err != nil {
	// 			log.Panicln(err)
	// 		}

	// 		Expect(orders[0].BaseAsset).To(Equal("ETH"))
	// 	})
	// })

	// Describe("CancelOrderAll", func() {
	// 	BeforeEach(func() {
	// 		_, err := client.Order("buy", "ETH/BTC", "0.018",
	// 			"0.05")
	// 		if err != nil {
	// 			log.Panicln(err)
	// 		}
	// 	})

	// 	It("cancels all orders", func() {
	// 		res, err := client.CancelOrderAll("ETH/BTC", "buy")
	// 		if err != nil {
	// 			log.Panicln(err)
	// 		}

	// 		Expect(res.Total).To(BeEquivalentTo(1))
	// 		Expect(res.Canceled).To(BeEquivalentTo(1))

	// 	})
	// })

	// Describe("CancelOrder", func() {
	// 	var coid string

	// 	symbol := "ETH/BTC"

	// 	BeforeEach(func() {
	// 		res, err := client.Order("buy", "ETH/BTC", "0.018",
	// 			"0.05")
	// 		if err != nil {
	// 			log.Panicln(err)
	// 		}

	// 		coid = res.COID
	// 	})

	// 	It("returns OrderResponse", func() {
	// 		res, err := client.
	// 			CancelOrder(symbol, coid)
	// 		if err != nil {
	// 			log.Panicln(err)
	// 		}

	// 		Expect(res.Action).To(Equal("cancel"))
	// 	})
	// })
})
