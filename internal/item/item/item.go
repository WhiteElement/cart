package item

import "time"

type Item struct {
	Id      int
	Name    string
	Checked bool
	ListId  int
	Updated time.Time
}
