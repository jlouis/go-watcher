package matilde

import (
	"github.com/shopgun/matilde/action"
	"github.com/shopgun/matilde/db"
	"github.com/shopgun/matilde/io"
)

// Process reads channel and performs cleaning and returns
// cleaned event
func Process(in io.Input, db db.Connector) (event action.Event) {
	event.Init(in.Content, in.Type)
	event.Clean(db)
	return
}
