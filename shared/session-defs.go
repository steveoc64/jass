package shared

import "time"

type Session struct {
	ID        int       `db:"id"`
	Date      time.Time `db:"date"`
	UserID    int       `db:"user_id"`
	IP        string    `db:"ip"`
	Referrer  string    `db:"referrer"`
	UserAgent string    `db:"user_agent"`
}
