package users_test

import (
	"github.com/coopernurse/gorp"
	. "github.com/crowdint/coffeboy/models/users"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Users", func() {
	var (
		dbmap *gorp.DbMap
		user  User
	)

	BeforeSuite(func() {
		dbmap, _ = InitDb()
	})

	BeforeEach(func() {
		user = User{
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
				err := user.Save(dbmap)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Get a register from the database", func() {
		Context("To the test database", func() {
			It("Should not return an error", func() {
				err := user.Save(dbmap)
				Expect(err).To(BeNil())

				user2, err := Get(user.ID, dbmap)

				Expect(err).To(BeNil())
				Expect(user2).To(Equal(&user))
			})
		})
	})

	Describe("Update a register from the database", func() {
		Context("To the test database", func() {
			It("Should not return an error and update the register", func() {
				err := user.Save(dbmap)
				Expect(err).To(BeNil())

				user.Name = "1111"
				user.Update(dbmap)

				user2, err := Get(user.ID, dbmap)

				Expect(err).To(BeNil())
				Expect(user2.Name).To(Equal("1111"))
			})
		})
	})

	Describe("Get registers from the database", func() {
		Context("From the test database", func() {
			It("Should return no errors and the items", func() {
				user.Save(dbmap)
				user2 := User{
					Name: "1111",
				}
				user2.Save(dbmap)

				user3 := User{
					Name: "12222",
				}
				user3.Save(dbmap)

				tokens, err := GetAll(dbmap)

				Expect(err).To(BeNil())
				Expect(tokens).Should(ContainElement(user))
				Expect(tokens).Should(ContainElement(user2))
				Expect(tokens).Should(ContainElement(user3))

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
