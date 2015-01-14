package orders

import (
	"database/sql"
	"os"

	"github.com/coopernurse/gorp"
	//For science
	_ "github.com/lib/pq"
)

type Order struct {
	ID      int64  `db:"id"`
	UserID  int64  `db:"user_ID"`
	Created string `db:"created"`
}

func (order *Order) Save(dbmap *gorp.DbMap) error {
	return dbmap.Insert(order)
}

func GetOrder(id int64, dbmap *gorp.DbMap) (*Order, error) {
	order := Order{}
	err := dbmap.SelectOne(&order, "SELECT * FROM orders WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (order *Order) Update(dbmap *gorp.DbMap) error {
	_, err := dbmap.Update(order)
	return err
}

func (order *Order) Delete(dbmap *gorp.DbMap) error {
	_, err := dbmap.Delete(order)
	return err
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
	dbmap.AddTableWithName(Order{}, "orders").SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}
