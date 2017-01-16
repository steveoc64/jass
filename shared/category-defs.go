package shared

type Category struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Descr string `db:"descr"`
}
