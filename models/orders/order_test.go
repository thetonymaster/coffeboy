package orders_test

import (
	"time"

	"github.com/coopernurse/gorp"
	. "github.com/crowdint/coffeboy/models/orders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Order", func() {

	var (
		dbmap *gorp.DbMap
	)

	BeforeSuite(func() {
		dbmap, _ = InitDb()
	})

	Describe("Create table in database", func() {
		Context("In a test database", func() {
			It("Create table and return no errors", func() {
				dbm, err := InitDb()
				Ω(err).Should(BeNil())
				Ω(dbm).ShouldNot(BeNil())
				dbm.Db.Close()
			})
		})
	})

	Describe("Save an order to the database", func() {
		Context("With a test databse", func() {
			It("Save it and do not return errors", func() {
				order := Order{
					UserID:  1,
					Created: time.Now().Format("2006-01-02T15:04:05.999999999Z07:00"),
				}

				err := order.Save(dbmap)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Get an order from the database", func() {
		Context("Wirh a test database", func() {
			It("It should save an order and then retrieve it", func() {
				order := Order{
					UserID:  1,
					Created: time.Now().Format("2006-01-02T15:04:05.999999999Z07:00"),
				}
				err := order.Save(dbmap)
				Expect(err).To(BeNil())

				order2, err := GetOrder(order.ID, dbmap)
				Expect(err).To(BeNil())

				Ω(*order2).Should(Equal(order))

			})
		})
	})

	AfterSuite(func() {
		dbmap.DropTables()
		dbmap.Db.Close()
	})

})
