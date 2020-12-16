package schema

import "time"

type Todo struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Note     string    `json:"note"`
	Deadline time.Time `json:"deadline"`
}
