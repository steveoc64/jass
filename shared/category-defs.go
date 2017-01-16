package shared

import "fmt"

type Category struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Descr string `db:"descr"`
}

func (t Category) ModelId() string {
	return fmt.Sprintf("%d", t.ID)
}

func (t Category) RootURL() string {
	return "/api/category"
}
