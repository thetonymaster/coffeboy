package orders

import (
	"github.com/coopernurse/gorp"
	"github.com/crowdint/coffeboy/utils"
	//For science
	_ "github.com/lib/pq"

	"encoding/json"
)

type Order struct {
	ID            string             `db:"id" json:"id"`
	UserID        int64              `db:"user_id" json:"user_id"`
	Created       string             `db:"created" json:"created_at"`
	Updated       string             `db:"updated" json:"updated_at"`
	Completed     string             `db:"completed" json:"completed_at"`
	Email         string             `db:"email" json:"email"`
	Quantity      string             `db:"quantity" json:"total_quantity"`
	LineItems     []OrderVariantData `db:"-" json:"line_items"`
	LineItemsJson string             `db:"line_items_json" json:"-"`
}

type OrderVariantData struct {
	ID       string `json:"variant_id"`
	Quantity int    `json:"quantity"`
}

func (order *Order) Save(dbmap *gorp.DbMap) error {
	lineItemsJson, err := json.Marshal(order.LineItems)
	if err != nil {
		return err
	}

	order.LineItemsJson = string(lineItemsJson)

	return dbmap.Insert(order)
}

func (order *Order) Update(dbmap *gorp.DbMap) error {
	lineItemsJson, err := json.Marshal(order.LineItems)
	if err != nil {
		return err
	}

	order.LineItemsJson = string(lineItemsJson)

	_, err = dbmap.Update(order)
	return err
}

func (order *Order) Delete(dbmap *gorp.DbMap) error {
	_, err := dbmap.Delete(order)
	return err
}

func (order *Order) Marshal() ([]byte, error) {
	return json.Marshal(order)
}

func GetOrder(orderId string, dbmap *gorp.DbMap) (*Order, error) {
	order := Order{}
	err := dbmap.SelectOne(&order, "SELECT * FROM orders WHERE id = $1", orderId)
	if err != nil {
		return nil, err
	}

	var lineItems []OrderVariantData

	err = json.Unmarshal([]byte(order.LineItemsJson), &lineItems)
	if err != nil {
		return nil, err
	}

	order.LineItems = lineItems

	return &order, nil
}

func InitDb() (*gorp.DbMap, error) {
	return utils.CreateTableWithID("orders", Order{}, false)
}
