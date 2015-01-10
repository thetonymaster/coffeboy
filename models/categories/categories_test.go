package categories_test

import (
	"github.com/coopernurse/gorp"
	. "github.com/crowdint/coffeboy/models/categories"
	"github.com/crowdint/coffeboy/models/products"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Categories", func() {
	var (
		dbmap    *gorp.DbMap
		category *Category
	)

	BeforeSuite(func() {
		dbmap, _ = InitDb()
	})

	BeforeEach(func() {
		category = &Category{
			Name: "1234",
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

				err := category.Save(dbmap)
				Expect(err).To(BeNil())

			})
		})
	})

	Describe("Get a register from the database", func() {
		Context("To the test database", func() {
			It("Should not return an error", func() {
				err := category.Save(dbmap)
				Expect(err).To(BeNil())

				category2, err := Get(category.ID, dbmap)

				category2.Products = nil
				Expect(err).To(BeNil())
				Expect(category2).To(Equal(category))

				product := products.Product{
					CategoryID: category.ID,
					Name:       "blahblahblaj",
				}

				err = product.Save(dbmap)
				Expect(err).To(BeNil())

				category2, err = Get(category.ID, dbmap)
				Expect(err).To(BeNil())
				Expect(category2.Products).To(ContainElement(product))

			})
		})
	})

	Describe("Update a register from the database", func() {
		Context("To the test database", func() {
			It("Should not return an error and update the register", func() {
				err := category.Save(dbmap)
				Expect(err).To(BeNil())
				category.Name = "1111"
				category.Update(dbmap)

				category2, err := Get(category.ID, dbmap)

				Expect(err).To(BeNil())
				Expect(category2.Name).To(Equal("1111"))
			})
		})
	})

	Describe("Get registers from the database", func() {
		Context("From the test database", func() {
			It("Should return no errors and the items", func() {
				category.Save(dbmap)
				category2 := Category{
					Name: "1111",
				}
				category2.Save(dbmap)

				category3 := Category{
					Name: "12222",
				}
				category3.Save(dbmap)

				product := products.Product{
					CategoryID: category.ID,
					Name:       "ABD",
				}

				product.Save(dbmap)

				category.Products = []products.Product{product}
				cats, err := GetAll(dbmap)

				Expect(err).To(BeNil())
				Ω(cats).Should(ContainElement(*category))
				Ω(cats).Should(ContainElement(category2))
				Ω(cats).Should(ContainElement(category3))

				Ω(category.Products).Should(ContainElement(product))

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
