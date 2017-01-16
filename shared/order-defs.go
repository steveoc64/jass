package shared

import "time"

type SaleOrder struct {
	ID        int       `db:"id"`
	Date      time.Time `db:"date"`
	UserID    int       `db:"user_id"`
	AddressID int       `db:"address_id"`
	Items     []*SaleOrderItem
}

type SaleOrderItem struct {
	OrderID int `db:"order_id"`
	ItemID  int `db:"item_id"`
	Qty     int `db:"qty"`
}
