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

	Describe("Update an order from the database", func() {
		Context("With a test database", func() {
			It("It should update an existing order and not return errors", func() {
				order := Order{
					UserID:  1,
					Created: time.Now().Format("2006-01-02T15:04:05.999999999Z07:00"),
				}
				err := order.Save(dbmap)
				Expect(err).To(BeNil())

				newdate := time.Now().Format("06-01-02T15:04:05.999999999Z07:00")
				order.Created = newdate

				err = order.Update(dbmap)
				Expect(err).To(BeNil())

				order2, err := GetOrder(order.ID, dbmap)
				Expect(err).To(BeNil())
				Ω(order2.Created).Should(Equal(newdate))

			})
		})
	})

	Describe("Delete an order from the database", func() {
		Context("With a test database", func() {
			It("It should delete an existing order and return no errors", func() {
				order1 := Order{
					UserID:  2,
					Created: time.Now().Format("2006-01-02T15:04:05.999999999Z07:00"),
				}
				err := order1.Save(dbmap)
				Expect(err).To(BeNil())

				err = order1.Delete(dbmap)
				Expect(err).To(BeNil())

			})
		})
	})

	AfterSuite(func() {
		dbmap.DropTables()
		dbmap.Db.Close()
	})

})
