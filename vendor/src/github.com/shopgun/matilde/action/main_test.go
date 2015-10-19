// TODO  revamp this test suite
/*
 *this test suite tests the un-marshaling logic
 i.e. do the events acutally un-marshall  the way we
 want it.
 Given a speficic string--> event.
*/

package action

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/shopgun/matilde/db"
	"github.com/shopgun/matilde/env"
	"github.com/shopgun/matilde/events/clean"
	"github.com/shopgun/matilde/events/event"
	"os"
	"reflect"
	"testing"

	log "github.com/Sirupsen/logrus"
)

var line []byte

var e env.VARS
var connection db.Connector

func init() {
	e = env.Init()
	if os.Getenv("travis") != "" {
		e.DbName = ""
		// we can then set up mysql travis clone
	} else {
		var err error
		connection, err = db.Connect(e)
		if err != nil {
			log.Fatal(err)
		}
		err = connection.DB.Ping()
		if err != nil {
			log.Fatal(err)
		}
	}
}
func TestEmptyJson(t *testing.T) {
	eventTest := Event{}
	line := []byte(``)
	eventTest.Init(line, line)
	if len(eventTest.Errors) < 1 {
		t.Errorf("should be error")
	}
}

func TestType(t *testing.T) {
	line := []byte(`{
	"scope": "offer",
	"action": "view"}
	`)
	eventTest := Event{}
	eventTest.Init(line, []byte("offer.view"))
	if eventTest.Type != "offer.view" {
		t.Errorf("got %s wanted %s", eventTest.Type, "offer.view")
	}
	//TODO decide what to do with this test
	eventTest = Event{}
	line = []byte(`{
	"action": "view"}
	`)
	eventTest.Init(line, line)
	if len(eventTest.Errors) < 1 {
		//t.Errorf("should be error")
	}
	eventTest = Event{}
	line = []byte(`{
	"action": ""}
	`)
	eventTest.Init(line, line)
	if len(eventTest.Errors) < 1 {
		//t.Errorf("should be error")
	}
}

// TestZeroValues checks if there are non zero values when input is empty json blob.
func TestZeroValues(t *testing.T) {
	// empty string and db object
	in, _ := UnmarshalEvent(line)
	out := new(event.Event)
	cond := reflect.DeepEqual(in, out)
	if cond == false {
		t.Errorf("All fields must be zero, imput string is empty")
		fmt.Printf("%+v\n", in)
	}
}

//TestMissingRequired tests for errors when the reuired fields are missing.
//by design there should be an error which means the event is skipped
func TestMissingRequired(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	in, _ := UnmarshalEvent(line)
	_, err := CleanOfferView(connection, in)
	if err == nil {
		t.Log(err)
		t.Errorf("Timestamp can't be empty, and error is nil. ")
	}

}

func randomstring(size int32) string {
	rb := make([]byte, size)
	_, err := rand.Read(rb)

	if err != nil {
		fmt.Println(err)
	}
	rs := base64.URLEncoding.EncodeToString(rb)
	return rs
}

func TestJsonparser(t *testing.T) {
	// random string  with errors
	rstring := randomstring(32)
	_, err := UnmarshalEvent([]byte(rstring))
	if err == nil {
		t.Errorf("Should't unmarshall a random string into a vaild structure ")
	}
}

func TestIp(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	line := []byte(`{"uuid": "8F0DA3E6A6F311E1B35A12313B05658A",
	"device": 1,
	"app": 22,
	"ip": 111111536617607,
	"timestampMs": 1330560015000,
	"timestamp": 1330560015,
	"scope": "offer",
	"action": "view",
	"date_year": "2012-01-01T00:00:00+00:00",
	"date_quarter": "2012-01-01T00:00:00+00:00",
	"date_month": "2012-03-01T00:00:00+00:00",
	"date_week": "2012-02-27T00:00:00+00:00",
	"date_day": "2012-03-01T00:00:00+00:00",
	"date_hour": "2012-03-01T00:00:00+00:00",
	"date_minute": "2012-03-01T00:00:00+00:00",
	"date_second": "2012-03-01T00:00:15+00:00",
	"year": 2012,
	"quarter": 1}`)

	var err error
	_, err = UnmarshalEvent(line)
	if err == nil {
		t.Log(err)
		t.Errorf("this should fail")
	}
}

func TestEventLeaking(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	line := []byte(`{"uuid": "8F0DA3E6A6F311E1B35A12313B05658A",
	"device": 1,
	"app": 22,
	"ip": 1536617607,
	"timestampMs": 1330560015000,
	"timestamp": 1330560015,
	"scope": "offer",
	"action": "click",
	"date_year": "2012-01-01T00:00:00+00:00",
	"date_quarter": "2012-01-01T00:00:00+00:00",
	"date_month": "2012-03-01T00:00:00+00:00",
	"date_week": "2012-02-27T00:00:00+00:00",
	"date_day": "2012-03-01T00:00:00+00:00",
	"date_hour": "2012-03-01T00:00:00+00:00",
	"date_minute": "2012-03-01T00:00:00+00:00",
	"date_second": "2012-03-01T00:00:15+00:00",
	"offer": 5088,
	"year": 2012,
	"quarter": 1}`)
	var err error
	in, err := UnmarshalEvent(line)
	if err != nil {
		t.Log(err)
		t.Errorf("this should not fail")
	}
	out, errs := CleanOfferClick(connection, in)
	for _, err := range errs {
		if err != nil {
			t.Logf("%v", err)
			t.Errorf("this should not fail")
		}
	}
	// now use  as semni empty string if there is event leaking
	// the event will not be empty
	line = []byte(`{"timestamp": 10, "offer":1992, "uuid":"22azGR",
	"scope": "offer",
	"action": "view"}`)
	in, err = UnmarshalEvent(line)
	if err != nil {
		t.Log(err)
		t.Errorf("this should not fail")
	}
	out, errs = CleanOfferClick(connection, in)
	for _, err := range errs {
		if err != nil && err != sql.ErrNoRows {
			t.Log(err)
			t.Errorf("this should not fail")
		}
	}
	// this event should be empty
	// except for required fields
	// check dirty event
	empty := new(event.Event)
	cond := reflect.DeepEqual(in, empty)
	if cond != false {
		t.Errorf("%v - %v", empty, in)
	}
	// check cleaned event

	emptyClean := new(clean.Offer_Click)
	cond = reflect.DeepEqual(out, emptyClean)
	if cond != false {
		t.Errorf("%v - %v", emptyClean, out)
	}
}

// TestBEventcheck check if the clean function is able to handle properly
// events with wrong types
func TestBadEventcheck(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	line := []byte(`
	{
	"uuid": "efaf172f1e44432e89218c5223b39dce",
	"timestampMs": 1350017977436,
	"timestamp": 1350017977,
	"scope": "offer",
	"action": "search",
	"user_id":99999999
	}`)
	Type := []byte("offer.search")
	e := Event{}
	e.Init(line, Type)
	e.Clean(connection)
	if len(e.Errors) != 2 {
		t.Errorf("Sould just have two errors! Got %v", e.Errors)
	}
	t.Log(e.Errors)
}

// TestEventcheck check if the clean function is able to handle properly
// events with wrong types
func TestEventcheck(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	line := []byte(`
	{
	"distance": 10000,
	"offset": 0,
	"limit": 40,
	"results": 3,
	"query": 1850,
	"offers": [
	567935,
	560224,
	574840
	],
	"dealers": [
	9,
	54,
	6
	],
	"uuid": "efaf172f1e44432e89218c5223b39dce",
	"device": 8,
	"app": 32,
	"ip": 3165721174,
	"timestampMs": 1350017977436,
	"timestamp": 1350017977,
	"scope": "offer",
	"action": "search",
	"date_year": "2012-01-01T00:00:00+00:00",
	"date_quarter": "2012-10-01T00:00:00+00:00",
	"date_month": "2012-10-01T00:00:00+00:00",
	"date_week": "2012-10-08T00:00:00+00:00",
	"date_day": "2012-10-12T00:00:00+00:00",
	"date_hour": "2012-10-12T04:00:00+00:00",
	"date_minute": "2012-10-12T04:59:00+00:00",
	"date_second": "2012-10-12T04:59:37+00:00",
	"year": 2012,
	"quarter": 4,
	"month": 10,
	"week": 41,
	"day": 12,
	"hour": 4,
	"minute": 59,
	"second": 37,
	"geocoded": 0,
	"accuracy": 1,
	"latitude": 55.67831,
	"longitude": 12.53461,
	"geohash": "u3butpzmtt",
	"locationDetermined": 1350025178
	}`)
	in, err := UnmarshalEvent(line)
	if err != nil {
		t.Log(err)
		t.Errorf("unmarhsaling.. this should not fail")
	}
	out := clean.Offer_View{}
	out, errors := CleanOfferView(connection, in)
	if errors != nil {
		t.Log(errors)
		t.Errorf("this should not fail")
	}
	errors = eventChecker(in, out)
	if errors != nil {
		t.Log(errors)
		t.Errorf("this should not fail event is valid")
	}
	// fuck the event
	out.Uuid = ""
	errors = eventChecker(in, out)
	if errors == nil {
		t.Log(errors)
		t.Errorf("this should fail uuid is empty")
	}
	out.Type = ""
	errors = eventChecker(in, out)
	if errors == nil {
		t.Log(errors)
		t.Errorf("this should fail uuid is empty")
	}
	e := new(Event)
	e.Raw = *in
	e.Type = "offer.search"
	e.Clean(connection)
	for _, i := range e.Errors {
		if i != nil {
			t.Errorf("no mismatch in action.scope should not fail")
		}
	}
	e = new(Event)
	e.Raw = *in
	e.Type = "offer.view"
	e.Clean(connection)
	for _, i := range e.Errors {
		if i == nil {
			t.Errorf("mismatch in action.scope should fail")
		}
	}
	e = new(Event)
	e.Raw = *in
	e.Type = "foo.bar"
	e.Clean(connection)
	for _, i := range e.Errors {
		if i == nil {
			t.Errorf("mismatch in action.scope should fail")
		}
	}
}

//TestMissingInJson check if the timestamp are properly hanlded
func TestMissingInJson(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	line := []byte(`
		{
		"dealers": [
		9,
		54,
		6
		],
		"stores": [],
		"uuid": "efaf172f1e44432e89218c5223b39dce",
		"timestampMs": 1350017977436,
		"timestamp": 1350017977,
		"scope": "offer",
		"action": "search"
		}`)
	var err error
	in, err := UnmarshalEvent(line)
	if err != nil {
		t.Log(err)
		t.Errorf("unmarhsaling.. this should not fail")
	}
	_, errs := CleanOfferSearch(connection, in)
	for _, err := range errs {
		if err != nil {
			t.Log(err)
			t.Errorf("this should not fail")
		}
	}
}

func TestMissingInJsontwo(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	line := []byte(`
	{
	"dealers": [
	9,
	54,
	6
	],
	"uuid": "efaf172f1e44432e89218c5223b39dce",
	"timestampMs": 1350017977436,
	"timestamp": 1350017977,
	"scope": "offer",
	"action": "search"
	}`)
	var err error
	in, err := UnmarshalEvent(line)
	if err != nil {
		t.Log(err)
		t.Errorf("unmarhsaling.. this should not fail")
	}
	_, errs := CleanOfferSearch(connection, in)
	for _, err := range errs {
		if err != nil {
			t.Log(err)
			t.Errorf("this should not fail")
		}
	}
}

// test duration
func TestDuration(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	line := []byte(`{"uuid": "8F0DA3E6A6F311E1B35A12313B05658A",
	"device": 1,
	"app": 22,
	"timestampMs": 100000,
	"timestamp": 100,
	"scope": "catalog",
	"action": "pageview",
	"duration": -100,
	"year": 2012,
	"quarter": 1}`)
	var err error
	in, err := UnmarshalEvent(line)
	if err != nil {
		t.Log(err)
		t.Errorf("unmarhsaling.. this should not fail")
	}
	out, errs := CleanCatalogPageview(connection, in)
	for _, err := range errs {
		if err == nil {
			t.Log(err)
			t.Errorf("this should fail")
		}
	}
	value, err := event.GetValue("Duration", out)
	if value != float64(0) {
		t.Errorf("got %+v", out)
		t.Errorf("got %v wanted %v", value, 0)
	}
}

//TestTImestamp check if the timestamp are properly hanlded
func TestTimestamp(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	line := []byte(`{"uuid": "8F0DA3E6A6F311E1B35A12313B05658A",
	"device": 1,
	"app": 22,
	"timestampMs": 100000,
	"timestamp": 100,
	"scope": "offer",
	"action": "view",
	"date_year": "2012-01-01T00:00:00+00:00",
	"date_quarter": "2012-01-01T00:00:00+00:00",
	"date_month": "2012-03-01T00:00:00+00:00",
	"date_week": "2012-02-27T00:00:00+00:00",
	"date_day": "2012-03-01T00:00:00+00:00",
	"date_hour": "2012-03-01T00:00:00+00:00",
	"date_minute": "2012-03-01T00:00:00+00:00",
	"date_second": "2012-03-01T00:00:15+00:00",
	"year": 2012,
	"quarter": 1}`)
	var err error
	in, err := UnmarshalEvent(line)
	if err != nil {
		t.Log(err)
		t.Errorf("unmarhsaling.. this should not fail")
	}
	out := clean.Offer_View{}
	out, errs := CleanOfferView(connection, in)
	for _, err := range errs {
		if err != nil {
			t.Log(err)
			t.Errorf("this should not fail")
		}
	}
	value, err := event.GetValue("Timestamp", out)
	if value != 100.0 {
		t.Log(err)
		t.Errorf("got %+v", out)
		t.Errorf("got %s wanted %s", value, 100.0)
	}
	line = []byte(`{"uuid": "8F0DA3E6A6F311E1B35A12313B05658A",
	"device": 1,
	"app": 22,
	"timestamp": 100,
	"scope": "offer",
	"action": "view",
	"date_year": "2012-01-01T00:00:00+00:00",
	"date_quarter": "2012-01-01T00:00:00+00:00",
	"date_month": "2012-03-01T00:00:00+00:00",
	"date_week": "2012-02-27T00:00:00+00:00",
	"date_day": "2012-03-01T00:00:00+00:00",
	"date_hour": "2012-03-01T00:00:00+00:00",
	"date_minute": "2012-03-01T00:00:00+00:00",
	"date_second": "2012-03-01T00:00:15+00:00",
	"year": 2012,
	"quarter": 1}`)

	in, err = UnmarshalEvent(line)
	if err != nil {
		t.Log(err)
		t.Errorf("unmarhsaling.. this should not fail")
	}
	out, errs = CleanOfferView(connection, in)
	for _, err := range errs {
		if err != nil {
			t.Log(err)
			t.Errorf("this should not fail")
		}
	}
	value, err = event.GetValue("Timestamp", out)
	if value != 100.0 {
		t.Log(err)
		t.Errorf("got %+v", out)
		t.Errorf("got %s wanted %s", value, 100.0)
	}

}

func TestRecover(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	line := []byte(`
	{
	"uuid": "efaf172f1e44432e89218c5223b39dce",
	"scope": "user",
	"user_id":999,
	"action": "create"
	}`)
	var err error
	in, err := UnmarshalEvent(line)
	if err != nil {
		t.Errorf("unmarhsaling.. this should not fail")
	}
	_, err = CleanUserCreate(connection, in)
	if err == nil {
		t.Log(err)
		t.Errorf("this should fail")
	}
	// now try with a completely empty event and make sure we don't panic and
	/// get an error back.
	line = []byte(`
	{
	"uuid": "efaf172f1e44432e89218c5223b39dce",
	"scope": "offer",
	"user_id":999,
	"action": "suggest"
	}`)
	in, err = UnmarshalEvent(line)
	if err != nil {
		t.Errorf("unmarhsaling.. this should not fail")
	}
	_, err = CleanOfferSuggest(connection, in)
	if err == nil {
		t.Log(err)
		t.Errorf("this should fail")
	}
}
