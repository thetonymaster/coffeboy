package users

import (
	"database/sql"
	"os"

	"github.com/coopernurse/gorp"
	//For science
	_ "github.com/lib/pq"
)

type User struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	RoleID      int64  `db:"role_id" json:"role-id"`
	Password    string `db:"password" json:"password"`
	ImageURL    string `db:"image_url" json:"image-url"`
	Channel     string `db:"channel" json:"channel"`
	UUID        string `db:"uuid" json:"uuid"`
	DeviceToken string `db:"device_token" json:"device-token"`
	LastName    string `db:"last_name" json:"last-name"`
	Email       string `db:"email" json:"email"`
}

func (user *User) Save(dbmap *gorp.DbMap) error {
	return dbmap.Insert(user)

}

func Get(id int64, dbmap *gorp.DbMap) (*User, error) {
	var user User
	err := dbmap.SelectOne(&user, "select * from users where id=$1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (user *User) Update(dbmap *gorp.DbMap) error {
	_, err := dbmap.Update(user)
	return err
}

func (user *User) Delete(dbmap *gorp.DbMap) error {
	_, err := dbmap.Delete(user)
	return err
}

func GetAll(dbmap *gorp.DbMap) ([]User, error) {
	var users []User
	_, err := dbmap.Select(&users, "select * from users order by id")
	if err != nil {
		return nil, err
	}
	return users, nil
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
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}
