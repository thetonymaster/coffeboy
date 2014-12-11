package products

import (
	"database/sql"
	"os"

	"github.com/coopernurse/gorp"
	//For science
	_ "github.com/lib/pq"
)

type Product struct {
	ID         int64   `db:"id" json:"id,omitempty"`
	CategoryID int64   `json:"category-id" db:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price" db:"price"`
	ImageURL   string  `json:"image-url"`
	Stock      int     `json:"stock" db:"stock"`
}

func (product *Product) Save(dbmap *gorp.DbMap) error {
	return dbmap.Insert(product)

}

func Get(id int64, dbmap *gorp.DbMap) (*Product, error) {
	var product Product
	err := dbmap.SelectOne(&product, "select * from products where id=$1", id)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (product *Product) Update(dbmap *gorp.DbMap) error {
	_, err := dbmap.Update(product)
	return err
}

func (product *Product) Delete(dbmap *gorp.DbMap) error {
	_, err := dbmap.Delete(product)
	return err
}

func GetAll(dbmap *gorp.DbMap) ([]Product, error) {
	var products []Product
	_, err := dbmap.Select(&products, "select * from products order by id")
	if err != nil {
		return nil, err
	}
	return products, nil
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
	dbmap.AddTableWithName(Product{}, "products").SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}
