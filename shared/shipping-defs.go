package shared

type ShippingCode struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Descr string `db:"descr"`
}

type Region struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Descr string `db:"descr"`
}
