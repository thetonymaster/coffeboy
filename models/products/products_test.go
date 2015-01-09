package products_test

import (
	"github.com/coopernurse/gorp"
	. "github.com/crowdint/coffeboy/models/products"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Products", func() {
	var (
		dbmap   *gorp.DbMap
		product Product
	)

	BeforeSuite(func() {
		dbmap, _ = InitDb()
	})

	BeforeEach(func() {
		product = Product{
			CategoryID: 1,
			Name:       "1234",
		}
	})

	Describe("Try to connect to test database", func() {
		Context("To test database", func() {
			It("Should not return an error", func() {
				db, err := InitDb()
				Expect(err).To(BeNil())
				db.Db.Close()
			})
		})
	})

	Describe("Insert a register from the database", func() {
		Context("To the test database", func() {
			It("Should not return an error", func() {
				err := product.Save(dbmap)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Get a register from the database", func() {
		Context("To the test database", func() {
			It("Should not return an error", func() {
				err := product.Save(dbmap)
				Expect(err).To(BeNil())

				product2, err := Get(product.ID, dbmap)

				Expect(err).To(BeNil())
				Expect(product2).To(Equal(&product))
			})
		})
	})

	Describe("Get a register from a category from the database", func() {
		Context("To the test database", func() {
			It("Should not return an error", func() {
				err := product.Save(dbmap)
				Expect(err).To(BeNil())

				product2, err := GetByCategoryID(product.CategoryID, dbmap)

				Expect(err).To(BeNil())
				Expect(product2).To(Equal([]Product{product}))
			})
		})
	})

	Describe("Update a register from the database", func() {
		Context("To the test database", func() {
			It("Should not return an error and update the register", func() {
				err := product.Save(dbmap)
				Expect(err).To(BeNil())

				product.Name = "1111"
				product.Update(dbmap)

				product2, err := Get(product.ID, dbmap)

				Expect(err).To(BeNil())
				Expect(product2.Name).To(Equal("1111"))
			})
		})
	})

	Describe("Get registers from the database", func() {
		Context("From the test database", func() {
			It("Should return no errors and the items", func() {
				product.Save(dbmap)
				product2 := Product{
					Name: "1111",
				}
				product2.Save(dbmap)

				peoduct3 := Product{
					Name: "12222",
				}
				peoduct3.Save(dbmap)

				tokens, err := GetAll(dbmap)

				Expect(err).To(BeNil())
				Expect(tokens).Should(ContainElement(product))
				Expect(tokens).Should(ContainElement(product2))
				Expect(tokens).Should(ContainElement(peoduct3))

			})
		})
	})

	AfterEach(func() {
		dbmap.TruncateTables()
	})

	AfterSuite(func() {
		dbmap.DropTables()
		dbmap.Db.Close()
	})
})
