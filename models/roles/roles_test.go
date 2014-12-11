package roles_test

import (
	"github.com/coopernurse/gorp"
	. "github.com/crowdint/coffeboy/models/roles"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Roles", func() {
	var (
		dbmap *gorp.DbMap
		role  Role
	)

	BeforeSuite(func() {
		dbmap, _ = InitDb()
	})

	BeforeEach(func() {
		role = Role{
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
				err := role.Save(dbmap)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Get a register from the database", func() {
		Context("To the test database", func() {
			It("Should not return an error", func() {
				err := role.Save(dbmap)
				Expect(err).To(BeNil())

				role2, err := Get(role.ID, dbmap)

				Expect(err).To(BeNil())
				Expect(role2).To(Equal(&role))
			})
		})
	})

	Describe("Update a register from the database", func() {
		Context("To the test database", func() {
			It("Should not return an error and update the register", func() {
				err := role.Save(dbmap)
				Expect(err).To(BeNil())

				role.Name = "1111"
				role.Update(dbmap)

				role2, err := Get(role.ID, dbmap)

				Expect(err).To(BeNil())
				Expect(role2.Name).To(Equal("1111"))
			})
		})
	})

	Describe("Get registers from the database", func() {
		Context("From the test database", func() {
			It("Should return no errors and the items", func() {
				role.Save(dbmap)
				role2 := Role{
					Name: "1111",
				}
				role2.Save(dbmap)

				role3 := Role{
					Name: "12222",
				}
				role3.Save(dbmap)

				tokens, err := GetAll(dbmap)

				Expect(err).To(BeNil())
				Expect(tokens).Should(ContainElement(role))
				Expect(tokens).Should(ContainElement(role2))
				Expect(tokens).Should(ContainElement(role3))

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
