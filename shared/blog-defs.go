package shared

import (
	"fmt"
	"strings"
	"time"
)

type Blog struct {
	ID              int       `db:"id"`
	PostOrder       int       `db:"post_order"`
	Image           string    `db:"image"`
	Date            time.Time `db:"date"`
	Name            string    `db:"name"`
	Title           string    `db:"title"`
	Content         string    `db:"content"`
	ShareTwitter    int       `db:"share_twitter"`
	ShareFacebook   int       `db:"share_facebook"`
	ShareInstagram  int       `db:"share_instagram"`
	ShareGooglePlus int       `db:"share_google_plus"`
	Archived        bool      `db:"archived"`
}

func (b *Blog) GetLines() []string {
	return strings.Split(b.Content, "\n")
}

func (b *Blog) GetURL() string {
	return fmt.Sprintf("https://jass.com.au/blog/%d", b.ID)
}
