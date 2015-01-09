package categories

import (
	"database/sql"
	"os"

	"github.com/coopernurse/gorp"
	"github.com/crowdint/coffeboy/models/products"
	//For science
	_ "github.com/lib/pq"
)

type Category struct {
	ID       int64              `db:"id" json:"id"`
	Name     string             `db:"name" json:"name"`
	Products []products.Product `db:"-" json:"products,omitempty"`
}

func (category *Category) Save(dbmap *gorp.DbMap) error {
	return dbmap.Insert(category)

}

func Get(id int64, dbmap *gorp.DbMap) (*Category, error) {
	var category Category
	err := dbmap.SelectOne(&category, "select * from categories where id=$1", id)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (category *Category) Update(dbmap *gorp.DbMap) error {
	_, err := dbmap.Update(category)
	return err
}

func (category *Category) Delete(dbmap *gorp.DbMap) error {
	_, err := dbmap.Delete(category)
	return err
}

func GetAll(dbmap *gorp.DbMap) ([]Category, error) {
	var categories []Category
	_, err := dbmap.Select(&categories, "select * from categories order by id")
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func InitDb() (*gorp.DbMap, error) {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, err
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Category{}, "categories").SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}
