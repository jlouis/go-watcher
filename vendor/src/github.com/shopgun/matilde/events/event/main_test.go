/*
 *this test suite tests the conversion logic
 *i.e. given a data structure that has well defined data types does the event conversion logic works how we want it to work ?
 *I.e. do the ips convert to the expeceted ones or
 *are the defaults handled correctly
 */

package event

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/shopgun/matilde/db"
	"github.com/shopgun/matilde/env"
	"os"
	"testing"
)

const (
	path = "../../data/"
)

var ENV env.VARS
var connection db.Connector

func init() {
	ENV = env.Init()
	// NOTE slip db stuff on travis for now
	if os.Getenv("travis") != "" {
		ENV.DbName = ""
		// we can then set up mysql travis clone
	} else {
		var err error
		connection, err = db.Connect(ENV)
		if err != nil {
			log.Fatal(err)
		}
		err = connection.DB.Ping()
		if err != nil {
			log.Fatal(err)
		}
	}
}

type IPs struct {
	in  uint32
	out interface{}
}

var ipList = []IPs{
	IPs{3232235521, "::ffff:192.168.0.1"}, //192.168.0.1
	IPs{0, ""},
}

type apiID struct {
	in  string
	out string
}

var apiIDList = []apiID{
	apiID{"", ""},
	apiID{"null", ""},
	apiID{"4.1.1", "4.1.1"},
}

func TestDefaults(t *testing.T) {
	in := Event{}
	for i := 0; i < len(apiIDList); i++ {
		test := apiIDList[i]
		asd := Event{
			Api_app_version: test.in,
		}
		if out := asd.ApiAppVersion(); out != test.out {
			t.Errorf("ParseIP(%v) = %v, want %v", test.in, out, test.out)
		}
	}
	for i := 0; i < len(ipList); i++ {
		test := ipList[i]
		if out := in.Ipv4Ipv6(test.in); out != test.out {
			t.Errorf("ParseIP(%v) = %v, want %v", test.in, out, test.out)
		}
	}
}

type testme struct {
	Event
	Timestamp float64
	IP        string
	Id        int32
}

func TestZeroField(t *testing.T) {
	atest := testme{Timestamp: 0.0, IP: "", Id: 0}
	out := FieldZeroCheck("Timestamp", atest)
	if out == false {
		t.Errorf("Should return true because input is nil of the required type")
	}
	out = FieldZeroCheck("IP", atest)
	if out == false {
		t.Errorf("Should return true because input is nil of the required type")
	}
	out = FieldZeroCheck("Id", atest)
	if out == false {
		t.Errorf("Should return true because input is nil of the required type")
	}
}

func TestZeros(t *testing.T) {
	atest := new(testme)
	atest.IP = ""
	cond := FieldZeroCheck("IP", atest)
	if cond == false {
		t.Errorf("All fields must be zero, imput string is empty")
		fmt.Printf("%+v\n", atest)
	}
}

func checkname(name string, names []string) bool {
	for _, inname := range names {
		if inname == name {
			return true
		}
	}
	return false
}

// func TestNames checks that reflect gives back the correct values of the field names
func TestNames(t *testing.T) {
	atest := new(testme)
	names := GetNames(atest)
	cond := checkname("Timestamp", names)
	if cond != true {
		t.Errorf("Didn't return the expected fields names")
	}
	cond = checkname("IP", names)
	if cond != true {
		t.Errorf("Didn't return the expected fields names")
	}
	cond = checkname("Id", names)
	if cond != true {
		t.Errorf("Didn't return the expected fields names")
	}
}

func TestValue(t *testing.T) {
	//atest := new(testme)
	atest := testme{IP: "1", Id: 1, Timestamp: 0.0}
	names := GetNames(atest)
	for _, name := range names {
		extractedvalue, err := GetValue(name, atest)
		switch {
		case name == "Timestamp":
			if extractedvalue != 0.0 {
				t.Errorf("wrong value returned")
			}
			if err != nil {
				t.Errorf("Err should be nil")
			}
		case name == "IP":
			if extractedvalue != "1" {
				t.Errorf("wrong value returned")
			}
			if err != nil {
				t.Errorf("Err should be nil")
			}
		case name == "Id":
			var expectedval int32
			expectedval = 1
			if extractedvalue != expectedval {
				t.Errorf("wrong value returned")
			}
			if err != nil {
				t.Errorf("Err should be nil")
			}
		}
		name = "lol"
		extractedvalue, err = GetValue(name, atest)
		if extractedvalue != nil {
			t.Errorf("nonexisting field name should return nil")
		}
		if err == nil {
			t.Errorf("Err shouldn't be nil")
			t.Log(name)
		}
	}
}

func TestAppGroup(t *testing.T) {
	if ENV.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	in := Event{}
	in.App_id.Value = 22
	appgroups := in.AppGroups(connection)
	expected := []int32{8, 19}
	cond := TestEq(appgroups, expected)
	if !cond {
		t.Errorf("epxect%v, got %v", expected, appgroups)
	}
	//expect null array if empty app_id
	in = Event{}
	in.App_id.Value = 1
	appgroups = in.AppGroups(connection)
	expected = []int32{}
	cond = TestEq(appgroups, expected)
	if !cond {
		t.Errorf("epxect%v, got %v", expected, appgroups)
	}
	in = Event{}
	in.App_id.Value = 1
	in.App_id.Value = 99999999
	appgroups = in.AppGroups(connection)
	expected = []int32{}
	cond = TestEq(appgroups, expected)
	if !cond {
		t.Errorf("epxect%v, got %v", expected, appgroups)
	}
}

func TestCurrency(t *testing.T) {
	in := ""
	out := CheckCurrency(in)
	if out != "DKK" {
		t.Errorf("Bad missing value handling")
	}
	in = "1"
	out = CheckCurrency(in)
	if out != "DKK" {
		t.Errorf("Bad missing value handling")
	}
	in = "as"
	out = CheckCurrency(in)
	if out != "" {
		t.Errorf("Bad missing value handling got: %v, wanted empty", out)
	}
}
func TestDealers(t *testing.T) {
	badDealer := `
	{ "scope": "store",
    "action": "list",
    "device": 8,
    "uuid": "85312218b6614a62b0c7842a50d4960e",
    "app": 32,
    "ip": 0,
    "server_ip": 167798254,
    "user": null,
    "timestampMs": 1431390061564,
    "timestamp": 1431390062,
    "geocoded": 1,
    "latitude": 55.03309,
    "longitude": 10.53122,
    "geohash": "u1z3pdkwgp",
    "locationDetermined": "1431397262",
    "granularity": "store",
    "infoLevel": "all",
    "type": "suggested",
    "distance": 15000,
	"stores": [ "12191", "16922", "12202", "16720", "12929", "6428" ],
	"dealer": [ "79", "12", "122", "108", "9", "87", "69", "73", "183", "9" ]}
	`
	in, err := UnmarshalRealJson(badDealer)
	if err != nil {
		t.Error(err)
	}
	dealers := in.GetDealers()
	expected := []int32{79, 12, 122, 108, 9, 87, 69, 73, 183, 9}
	if !TestEq(dealers, expected) {
		t.Errorf("got %+v, weanted %+v", dealers, expected)
	}
	goodDealer := `
	{ "scope": "store",
    "action": "list",
    "device": 8,
    "uuid": "85312218b6614a62b0c7842a50d4960e",
    "app": 32,
    "ip": 0,
    "server_ip": 167798254,
    "user": null,
    "timestampMs": 1431390061564,
    "timestamp": 1431390062,
    "geocoded": 1,
    "latitude": 55.03309,
    "longitude": 10.53122,
    "geohash": "u1z3pdkwgp",
    "locationDetermined": "1431397262",
    "granularity": "store",
    "infoLevel": "all",
    "type": "suggested",
    "distance": 15000,
	"stores": [ "12191", "16922", "12202", "16720", "12929", "6428" ],
	"dealers": [ "79", "12", "122", "108", "9", "87", "69", "73", "183", "9" ]}
	`
	in, err = UnmarshalRealJson(goodDealer)
	dealers = in.GetDealers()
	expected = []int32{79, 12, 122, 108, 9, 87, 69, 73, 183, 9}
	if !TestEq(dealers, expected) {
		t.Errorf("got %+v, weanted %+v", dealers, expected)
	}
}

func TestPrice(t *testing.T) {
	if ENV.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	e := new(Event)
	e.Id.Value = 195262
	prep := 15.00
	check := e.GetPrice(connection)
	if check != prep {
		t.Errorf("got %v, wanted %v", check, 15.0)
	}
	e.Price.Value = 50.0
	check = e.GetPrice(connection)
	if check != 50.0 {
		t.Errorf("got %v, wanted %v", check, e.Price.Value)
	}
}

func TestPrePrice(t *testing.T) {
	if ENV.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	e := new(Event)
	e.Id.Value = 195262
	prep := 19.95
	check := e.GetPrePrice(connection)
	if check != prep {
		t.Errorf("got %v, wanted %v", check, prep)
	}
	e.PrePrice.Value = 50.0
	check = e.GetPrePrice(connection)
	if check != 50.0 {
		t.Errorf("got %v, wanted %v", check, prep)
	}
}

/*
 *Test the unmarshal custom logic
 */

type event struct {
	Test          MaybeBool     `json:"test"`
	Float         MaybeFloat    `json:"float"`
	Int           MaybeInt      `json:"int"`
	String        MaybeString   `json:"str"`
	Int64         MaybeInt64    `json:"int64"`
	MaybeIntArray MaybeIntArray `json:"intarr"`
}

func Unmarshal(lineStr string) (*event, error) {
	var err error
	err = nil
	line := []byte(lineStr)
	in := new(event)
	err = json.Unmarshal(line, &in)
	return in, err
}

func UnmarshalRealJson(lineStr string) (*Event, error) {
	var err error
	err = nil
	line := []byte(lineStr)
	in := new(Event)
	err = json.Unmarshal(line, &in)
	return in, err
}

type Bol struct {
	In  string
	Out bool
}

type Int struct {
	In  string
	Out int32
}

type Int64 struct {
	In  string
	Out int64
}

type Float struct {
	In  string
	Out float64
}

type Str struct {
	In  string
	Out string
}

var bolList = []Bol{
	Bol{`{"test":true}`, true},
	Bol{`{"test":false}`, false},
	Bol{`{"test":"true"}`, true},
	Bol{`{"test":"false"}`, false},
	Bol{`{"test":'true'}`, true},
	Bol{`{"test":'false'}`, false},
	Bol{`{"test":"bonkers"}`, false},
	Bol{`{"test": "   true    "}`, true},
	Bol{`{"test": ""}`, false},
}

var floatList = []Float{
	Float{`{"float": "10.0"}`, 10.0},
	Float{`{"float": 10.0}`, 10.0},
	Float{`{"float": ""}`, 0.0},
}

var intList = []Int{
	Int{`{"int": "10"}`, 10},
	Int{`{"int": 10}`, 10},
	Int{`{"int": ""}`, 0},
}

var maybeIntarrListForMaybeint = []intarr{
	intarr{`{ "int":["10", "10"]} `, []int32{10, 10}},
	intarr{`{ "int":[10, 10]} `, []int32{10, 10}},
	intarr{`{ "int":1} `, []int32{}},
	intarr{`{ "int":"foo"} `, []int32{}},
}

var intList64 = []Int64{
	Int64{`{"int64": "10"}`, 10},
	Int64{`{"int64": 10}`, 10},
	Int64{`{"int64": ""}`, 0},
}

var strList = []Str{
	Str{`{"str": "10"}`, "10"},
	Str{`{"str": 10}`, "10"},
	Str{`{"str": true}`, "true"},
	Str{`{"str": false}`, ""},
}

type intarr struct {
	In  string
	Out []int32
}

var maybeIntarrList = []intarr{
	intarr{`{ "intarr":["10", "10"]} `, []int32{10, 10}},
	intarr{`{ "intarr":[10, 10]} `, []int32{10, 10}},
	intarr{`{ "intarr":true} `, []int32{}},
	intarr{`{ "intarr":false} `, []int32{}},
}

func TestIntarr(t *testing.T) {
	for i := 0; i < len(maybeIntarrList); i++ {
		val := maybeIntarrList[i]
		in, err := Unmarshal(val.In)
		value := in.MaybeIntArray.IntArray
		cond := TestEq(value, val.Out)
		if !cond && err == nil {
			t.Errorf("Got %v wanted %v for %v ", value, val.Out, val.In)
		}

	}
}

func TestBool(t *testing.T) {
	for i := 0; i < len(bolList); i++ {
		val := bolList[i]
		in, err := Unmarshal(val.In)
		if value := in.Test.Value; value != val.Out && err == nil {
			t.Errorf("Got %v, wanted %v", value, val.Out)
		}
	}
}

func TestFloat(t *testing.T) {
	for i := 0; i < len(floatList); i++ {
		val := floatList[i]
		in, err := Unmarshal(val.In)
		if value := in.Float.Value; value != val.Out && err == nil {
			t.Errorf("Got %v, wanted %v", value, val.Out)
		}
	}
}

func TestInt(t *testing.T) {
	for i := 0; i < len(intList); i++ {
		val := intList[i]
		in, err := Unmarshal(val.In)
		if value := in.Int.Value; value != val.Out && err == nil {
			t.Errorf("Got %v, wanted %v", value, val.Out)
		}
	}
	for i := 0; i < len(intList); i++ {
		val := maybeIntarrListForMaybeint[i]
		in, err := Unmarshal(val.In)
		value := in.Int.IntArray.IntArray
		cond := TestEq(value, val.Out)
		if !cond && err == nil {
			t.Errorf("Got %v, wanted %v for %s", value, val.Out, val.In)
		}
	}
}

func TestInt64(t *testing.T) {
	for i := 0; i < len(intList); i++ {
		val := intList64[i]
		in, err := Unmarshal(val.In)
		if value := in.Int64.Value64; value != val.Out && err == nil {
			t.Errorf("Got %v, wanted %v", value, val.Out)
		}

	}
}

func TestStr(t *testing.T) {
	for i := 0; i < len(strList); i++ {
		val := strList[i]
		in, err := Unmarshal(val.In)
		if value := in.String.Value; value != val.Out && err == nil {
			t.Errorf("Got %v, wanted %v", value, val.Out)
		}

	}
}

func TestBadUsr(t *testing.T) {
	if ENV.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	type Usr struct {
		Gender    string
		Id        int32
		OutGender string
		Year      int64
		OutYear   int64
	}
	Users := []Usr{
		// non existing user id
		Usr{"", 99999999, "", 0, 0},
	}
	for i := 0; i < len(Users); i++ {
		val := Users[i]
		in := Event{}
		in.User.BYear.Value64 = val.Year
		in.User.Id.Value = val.Id
		in.Gender.Value = val.Gender
		if value := in.GetGender(connection); value != val.OutGender {
			t.Errorf("Got %v, wanted %v", value, val.OutGender)
		}
		if value := in.GetByear(connection); value != val.OutYear {
			t.Errorf("Got %v, wanted %v", value, val.OutYear)
		}
		if len(in.Errors) != 2 {
			t.Errorf("expect 2 errors got instead %v", in.Errors)
		}
	}

}

func TestUsr(t *testing.T) {
	if ENV.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	type Usr struct {
		Gender    string
		Id        int32
		OutGender string
		Year      int64
		OutYear   int64
	}
	Users := []Usr{
		Usr{"0", 0, "", 0, 0},
		// this is null on the db
		Usr{"0", 5149, "", 0, 0},
		Usr{"0", 11661, "", 1977, 1977},
		Usr{"2", 11661, "female", 0, 1977},
		Usr{"", 11661, "female", 0, 1977},
		Usr{
			Gender:    "1",
			Id:        11661,
			OutYear:   1977,
			OutGender: "male",
		},
	}
	for i := 0; i < len(Users); i++ {
		val := Users[i]
		in := Event{}
		in.User.BYear.Value64 = val.Year
		in.User.Id.Value = val.Id
		in.Gender.Value = val.Gender
		if value := in.GetGender(connection); value != val.OutGender {
			t.Log(in.Errors)
			t.Errorf("Got %v, wanted %v", value, val.OutGender)
		}
		if value := in.GetByear(connection); value != val.OutYear {
			t.Log(in.Errors)
			t.Errorf("Got %v, wanted %v", value, val.OutYear)
		}

	}

}

func TestLocation(t *testing.T) {
	type Loc struct {
		lat      float64
		long     float64
		hash     string
		hashOut  string
		coded    int32
		codedout bool
	}
	Locations := []Loc{
		// lat, long, hash ,code
		Loc{72.0739114882038, -2.021484375, "", "guyd1bxuzyzu", 0, false},
		Loc{55.8498344421387, 9.8280372619629, "u1yvxxhscgexrqyvvgexr", "u1yvxxhscgex", 0, false},
		Loc{55.730521, 12.464888, "", "", 1, true},
		Loc{0.0, 0.0, "", "", 1, true},
		Loc{55.8498344421387, 9.8280372619629, "u1yvxxhscgex", "u1yvxxhscgex", 0, false},
	}
	for i := 0; i < len(Locations); i++ {
		val := Locations[i]
		in := Location{
			Geocoded: val.coded,
			GeoHash:  val.hash,
		}
		in.Lat.Value = val.lat
		in.Long.Value = val.long

		if value := in.GeoCoded(); value != val.codedout {
			t.Errorf("Got %v, wanted %v", value, val.codedout)

		}
		if value := in.Geohash(); value != val.hashOut {
			t.Errorf("got %v, wanted %v", value, val.hashOut)
		}
	}
	val := Locations[2]
	in := Location{
		Geocoded: val.coded,
		GeoHash:  val.hash,
	}
	in.Lat.Value = val.lat
	in.Long.Value = val.long
	if value := in.Latitude(); value != 0.0 {
		t.Errorf("got %v, wanted %v", value, 0.0)
	}
	if value := in.Longitude(); value != 0.0 {
		t.Errorf("got %v, wanted %v", value, 0.0)
	}
	val = Locations[3]
	in = Location{
		Geocoded: val.coded,
		GeoHash:  val.hash,
	}
	in.Lat.Value = val.lat
	in.Long.Value = val.long
	if value := in.Latitude(); value != 0.0 {
		t.Errorf("got %v, wanted %v", value, 0.0)
	}
	if value := in.Longitude(); value != 0.0 {
		t.Errorf("got %v, wanted %v", value, 0.0)
	}
}

func TestTime(t *testing.T) {
	if ENV.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	type time struct {
		id         int32
		s          int32
		ms         int64
		pub        int32
		run        int32
		expires    int32
		out        float64
		outpub     float64
		outrun     float64
		outexpires float64
	}
	var a time
	a.s = 1409581851
	a.ms = 1409581851001
	a.out = 1409581851.001
	timez := []time{
		time{0, 0, 0, 0, 0, 0, 0.0, 0, 0, 0},
		time{0, 1, 0, 0, 0, 0, 1.0, 0, 0, 0},
		time{0, 1, 1000, 0, 0, 0, 1.0, 0, 0, 0},
		time{1923, 0, 1000, 1265504400, 1265504400, 1266109200, 1, 1265504400, 1265504400, 1266109200},
		time{1923, 0, 1000, 0, 0, 0, 1, 1265504400, 1265504400, 1266109200},
		time{1923, 0, 1000, 0, 0, 0, 1, 1265504400, 1265504400, 1266109200},
		a,
	}
	for i := 0; i < len(timez); i++ {
		val := timez[i]
		times := new(Times)
		times.S.Value = val.s
		times.Ms.Value64 = int64(val.ms)
		times.Exp.Value = val.expires
		times.Run.Value = val.run
		times.Pub.Value = val.pub
		event := new(Event)
		event.Times = *times
		event.Id.Value = val.id
		event.Scope = "offer"
		if value := event.Timestamp(); value != val.out {
			t.Errorf("timestamp: \n got %v, wanted %v", value, val.out)
		}
		if value := event.PublishTimeStamp(connection); value != val.outpub {
			t.Log(event.Errors)
			t.Errorf("PublishTimeStamp: \n got %v, wanted %v for %v", value, val.outpub, event.Id.Value)
		}
		if value := event.Expires(connection); value != val.outexpires {
			t.Log(event.Errors)
			t.Errorf("got %v, wanted %v", value, val.out)
		}
		if value := event.RunFrom(connection); value != val.outrun {
			t.Log(event.Errors)
			t.Errorf("got %v, wanted %v", value, val.out)
		}
	}
}

func testNullid(t *testing.T) {

}

func TestType(t *testing.T) {
	type version struct {
		in  string
		out string
	}
	versions := []version{
		version{"", "production"},
		version{"null", "production"},
		version{"production", "production"},
		version{"edge", "edge"},
		version{"dev", "dev"},
	}
	for i := 0; i < len(versions); i++ {
		evt := new(Event)
		val := versions[i]
		evt.Apienv = val.in
		in := evt.ApiEnv()
		if in != val.out {
			t.Errorf("Got %+v, wanted %+v", in, val.out)
		}

	}
	type scopeaction struct {
		scope  string
		action string
		out    string
	}
	types := []scopeaction{
		{"offer", "view", "offer.view"},
		{"offer", "", ""},
		{"", "click", ""},
	}
	for i := 0; i < len(types); i++ {
		val := types[i]
		event := new(Event)
		event.Action = val.action
		event.Scope = val.scope
		res := event.Type()
		if res != val.out {
			t.Errorf("Got %+v, wanted %+v", res, val.out)
		}
	}
}

func TestCatalog(t *testing.T) {
	if ENV.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}

	type ctlg struct {
		Id           int32
		catalogid    int32
		outCatalogid int32
	}
	ctlgs := []ctlg{
		ctlg{0, 0, 0},
		ctlg{1922, 0, 67},
	}
	for i := 0; i < len(ctlgs); i++ {
		val := ctlgs[i]
		in := Event{}
		in.Offer.Value = val.Id
		in.Catalog.Value = val.catalogid
		if value := in.GetCatalogId(connection); value != val.outCatalogid {
			t.Log(in.Errors)
			t.Errorf("Got %v, wanted %v", value, val.outCatalogid)
		}
	}
}

/*
 *Test reflection methods
 */

func TestZero(t *testing.T) {
	// empty string and db object
	type atest struct {
		Anint   int
		Astring string
		Anarray []int32
	}
	in := atest{}
	iszero := ZeroCheck(in)
	for _, cond := range iszero {
		if cond == false {
			t.Errorf("All fields must be zero")
			fmt.Printf("%+v\n", in)
		}
	}
	in = atest{0, "", []int32{}}
	iszero = ZeroCheck(in)
	for _, cond := range iszero {
		if cond == false {
			t.Errorf("All fields must be zero")
			fmt.Printf("%+v\n", in)
		}
	}
	in = atest{1, "foo", []int32{1}}
	iszero = ZeroCheck(in)
	for _, cond := range iszero {
		if cond != false {
			t.Errorf("All fields must be not zero")
			fmt.Printf("%+v\n", in)
		}
	}
}
