package roles

import (
	"database/sql"
	"os"

	"github.com/coopernurse/gorp"

	//For science
	_ "github.com/lib/pq"
)

type Role struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (role *Role) Save(dbmap *gorp.DbMap) error {
	return dbmap.Insert(role)

}

func Get(id int64, dbmap *gorp.DbMap) (*Role, error) {
	var role Role
	err := dbmap.SelectOne(&role, "select * from roles where id=$1", id)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (role *Role) Update(dbmap *gorp.DbMap) error {
	_, err := dbmap.Update(role)
	return err
}

func (role *Role) Delete(dbmap *gorp.DbMap) error {
	_, err := dbmap.Delete(role)
	return err
}

func GetAll(dbmap *gorp.DbMap) ([]Role, error) {
	var roles []Role
	_, err := dbmap.Select(&roles, "select * from roles order by id")
	if err != nil {
		return nil, err
	}
	return roles, nil
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
	dbmap.AddTableWithName(Role{}, "roles").SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}
