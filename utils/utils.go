package utils

import (
	"database/sql"
	"log"
	"os"

	"github.com/coopernurse/gorp"
	"github.com/crowdint/coffeboy/imageutil"
)

func UploadAndResizeImage(maxWidth, maxHeight uint, fileBytes []byte, path string) {
	fileResized, err := imageutil.Resize(fileBytes, maxWidth, maxHeight)
	if err != nil {
		log.Printf("Error %s\n", err.Error())
		return
	}

	err = imageutil.Upload(fileResized, path)
	if err != nil {
		log.Printf("Error %s\n", err.Error())
		return
	}

}

func UploadImages(fileBytes []byte, identifier, folder string) {
	go imageutil.Upload(fileBytes, folder+"/"+identifier+".jpg")
	go UploadAndResizeImage(400, 300, fileBytes, folder+"/"+identifier+"_small.jpg")
}

func CreateTableWithID(tableName string, entity interface{}) (*gorp.DbMap, error) {
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
	dbmap.AddTableWithName(entity, tableName).SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}
