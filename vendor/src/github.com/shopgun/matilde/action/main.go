// Package action provides the functions/types required to perform the actual cleaning.
// The basic idea is to  have an Event struct that holds all the details about the event
// Depending on the type, which needs to be set either using the path of a file or by some routing key, a cleaning operation is performed.
package action

import (
	"encoding/json"
	"expvar"
	"fmt"
	"time"

	"github.com/shopgun/matilde/events/event"
)

var (
	cleaned    = expvar.NewInt("Cleaned events")
	serialized = expvar.NewInt("Serialized events")
)

// Event contains the informations about the event
// Raw, Cleaned, Serialized contain the just-Unmarshaled, cleaned, and seraialized event respectively.
type Event struct {
	BadEvent  bool
	Cleaned   interface{}
	Errors    Errors
	Line      []byte
	Raw       event.Event
	Seralized []byte
	Type      string
	Timestamp time.Time
	//TODO also filename
}

//String make sure Event implement a pretty print interface
func (e Event) String() string {
	msg := fmt.Sprintf("Line: %s, Type: %v,Raw: %v, Cleaned: %+v \n Errors\n  %v \n ", e.Line, e.Type, e.Raw, e.Cleaned, e.Errors)
	return msg
}

// Errors holds a an aray of errors that satisfy the
type Errors []error

func (anerror Errors) Error() string {
	//Error make sure tha Erros satisfied the error interface
	if len(anerror) == 1 {
		return anerror[0].Error()
	}
	idx := len(anerror)
	var msg string
	if idx > 1 {
		for i := 0; i < idx; i++ {
			msg += anerror[i].Error() + "\n"
		}
		return msg
	}
	return ""
}

// LogError stuffs the erros inside the event pointer
// this allows mutiple errors per event
func (e *Event) LogError(err error) {
	if err != nil {
		e.Errors = append(e.Errors, err)
	}
}

// LogErrors stuffs a slice of errrors if not nil inside the
// event pointer this allows mutiple errors per event
func (e *Event) LogErrors(err Errors) {
	for _, anerror := range err {
		if anerror != nil {
			e.Errors = append(e.Errors, anerror)
		}
	}
}

// Init puts the JSONblob in the event pointer
// Unmarshals the jsoon and puts the resultsing
// event.Event inside the event pointer.
func (e *Event) Init(JSONblob []byte, t []byte) {
	e.Line = JSONblob
	e.Type = string(t[:])
	if len(e.Line) < 1 {
		e.LogError(fmt.Errorf("JSON input is less than one char"))
		return
	}
	raw, err := UnmarshalEvent(JSONblob)
	if err != nil {
		e.LogError(err)
	}
	e.Raw = *raw
}

// eventChecker check if a cleaned event is valid
// currently checking for Timestamp, Uuid, and Type
func eventChecker(in *event.Event, out interface{}) (err Errors) {
	if event.FieldZeroCheck("Type", out) == true {
		err = append(err, fmt.Errorf("Empty type does not make any sense."))
	}
	if event.FieldZeroCheck("Timestamp", out) == true {
		err = append(err, fmt.Errorf(`Timestamp can't be empty.`))
	}
	if event.FieldZeroCheck("Uuid", out) == true {
		if cond, _ := event.GetValue("type", out); cond != "user.password_reset" {
			err = append(err, fmt.Errorf(`Uuid can't be empty.`))
		}
	}
	// compare to empty array of errors
	//if not empty log and skip
	if len(in.Errors) > 0 {
		for _, anerror := range in.Errors {
			if anerror != nil {
				err = append(err, anerror)
			}
		}
	}
	return err
}

// UnmarshalEvent un-packs a []byte from a single line
// in  an event.Event struct, which is a struct that contains
// ALL the possible field for our events.
// This allows us to keep it general, but poses some limitation
// on abusing a "same name for different things" policy.
func UnmarshalEvent(line []byte) (*event.Event, error) {
	var err error
	err = nil
	in := new(event.Event)
	err = json.Unmarshal(line, &in)
	return in, err
}
