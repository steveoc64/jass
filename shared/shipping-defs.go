package shared

import "fmt"

type ShippingCode struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Descr string `db:"descr"`
}

func (t ShippingCode) ModelId() string {
	return fmt.Sprintf("%d", t.ID)
}

func (t ShippingCode) RootURL() string {
	return "/api/shipping"
}

type Region struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Descr string `db:"descr"`
}

func (t Region) ModelId() string {
	return fmt.Sprintf("%d", t.ID)
}

func (t Region) RootURL() string {
	return "/api/region"
}
