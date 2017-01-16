package shared

import (
	"fmt"
	"time"
)

type SaleOrder struct {
	ID        int       `db:"id"`
	Date      time.Time `db:"date"`
	UserID    int       `db:"user_id"`
	AddressID int       `db:"address_id"`
	Items     []*SaleOrderItem
}

func (t SaleOrder) ModelId() string {
	return fmt.Sprintf("%d", t.ID)
}

func (t SaleOrder) RootURL() string {
	return "/api/order"
}

type SaleOrderItem struct {
	OrderID int `db:"order_id"`
	ItemID  int `db:"item_id"`
	Qty     int `db:"qty"`
}
