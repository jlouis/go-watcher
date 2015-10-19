package event

import (
	"encoding/json"
	"fmt"
	"github.com/shopgun/matilde/db"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"

	libgeohash "github.com/TomiHiltunen/geohash-golang"
)

/*
*
*Types
*
 */

// Event contains  shared attribudes between all the kind of events.
// List of ignored field
//remote_addr  `year` `quarter` `month` `week` `day` `hour` `minute` `second`
type Event struct {
	Action          string        `json:"action"`
	Api             string        `json:"api"`
	Api_app_version string        `json:"api_app_version"`
	Api_build       string        `json:"api_build"`
	Apienv          string        `json:"api_env"`
	App_id          MaybeInt      `json:"app"`
	Buy_url         string        `json:"buy_url"`
	Catalog         MaybeInt      `json:"catalog"`
	Catalogs        MaybeIntArray `json:"catalogs"`
	Currency        string        `json:"currency"`
	Dealer          MaybeInt      `json:"dealer"`
	Dealerid        MaybeInt      `json:"dealer_id"`
	Dealers         MaybeIntArray `json:"dealers"`
	Description     MaybeString   `json:"description"`
	Duration        float64       `json:"duration"`
	IP              uint32        `json:"ip"`
	// NOTE  this does not look right to me (we never used camel case)
	Is_owner bool  `json:"isOwner"`
	Limit    int32 `json:"limit"`
	Location
	Offer       MaybeInt      `json:"offer"`
	Offers      MaybeIntArray `json:"offers"`
	Offset      int32         `json:"offset"`
	Orientation MaybeInt      `json:"orientation"`
	Pages       MaybeIntArray `json:"pages"`
	PrePrice    MaybeFloat    `json:"preprice"`
	Price       MaybeFloat    `json:"price"`
	Push
	Query        MaybeString   `json:"query"`
	Query_string string        `json:"query_string"`
	Scope        string        `json:"scope"`
	Server_ip    uint32        `json:"server_ip"`
	Store        MaybeInt      `json:"store"`
	Stores       MaybeIntArray `json:"stores"`
	Shop_item    MaybeInt      `json:"shoppingitem"`
	Shop_list    MaybeInt      `json:"shoppinglist"`
	Results      MaybeInt      `json:"results"`
	Tick         bool          `json:"tick"`
	// NOTE this holds the finer details of the event
	/// and is thus event dependent
	Typeof MaybeString `json:"type"`
	Times
	User
	User_agent   string      `json:"user_agent"`
	View_session MaybeString `json:"view_session"`
	Errors       []error
}

// User contains attributes that characterize an user.
type User struct {
	BYear             MaybeInt64  `json:"birthYear"`
	Email             string      `json:"email"`
	Gender            MaybeString `json:"gender"`
	Locale            string      `json:"locale"`
	Name              string      `json:"name"`
	Id                MaybeInt    `json:"user"`
	User_id           MaybeInt    `json:"user_id"`
	Is_Dealer_admin   MaybeBool   `json:"is_dealer_admin"`
	Is_archive_user   MaybeBool   `json:"is_archive_user"`
	Is_uuid_ephemeral MaybeBool   `json:"is_uuid_ephemeral"`
	Provider          string      `json:"provider"`
	Uuid              MaybeString `json:"uuid"`
}

// Push  contains the infromation of a push event NOTE we never disucssed about this event.
type Push struct {
	Endpoint_id MaybeInt `json:"endpoint_id"`
	Push_type   string   `json:"push_type"`
}

// GetGender extracts the gender from a databse dump
// gender 1 is male gender 2 is female
func (e *Event) GetGender(connector db.Connector) (out string) {
	var gender string
	casea := e.User.Id.Value != 0
	caseb := e.User.User_id.Value != 0
	if casea || caseb {
		if e.User.Gender.Value == "" {
			var err error
			var user db.User
			if casea {
				user, err = connector.UserGetField(e.User.Id.Value)
			} else {
				user, err = connector.UserGetField(e.User_id.Value)
			}
			if err != nil {
				err = fmt.Errorf("GetGender %v", err)
				e.Errors = append(e.Errors, err)
			}
			// handle nil response from db
			if user.Gender.Valid {
				gender = strconv.Itoa(int(user.Gender.Int64))
			}
		} else {
			gender = e.User.Gender.Value
		}
	} else {
		// no id no gender duh
		gender = ""
	}
	// make sure the int response from different events is
	// changed to male/female/empty
	switch {
	case gender == "0":
		// NOTE return null stirng for empty gender
		gender = ""
	case gender == "1":
		gender = "male"
	case gender == "2":
		gender = "female"
	}
	return gender
}

// GetByear  extracts the birthday year from a databse dump
// NOTE if a value has more than 4 digits it's probably bogus.
func (e *Event) GetByear(connector db.Connector) (out int64) {
	casea := e.User.Id.Value != 0
	caseb := e.User.User_id.Value != 0
	if casea || caseb {
		if e.User.BYear.Value64 == 0 {
			var err error
			var user db.User
			if casea {
				user, err = connector.UserGetField(e.User.Id.Value)
			} else {
				user, err = connector.UserGetField(e.User_id.Value)
			}
			if err != nil {
				err = fmt.Errorf("GetByear %v", err)
				e.Errors = append(e.Errors, err)
			}
			if user.Byear.Valid {
				return user.Byear.Int64
			}
		}
		return int64(e.User.BYear.Value64)
	}
	return 0
}

// GetCatalogId  looks up the catalog id from the database
// NOTE that it is legit to have no catalog id f.ex
// single offers.
func (e *Event) GetCatalogId(db db.Connector) (catalogId int32) {
	if e.Offer.Value != 0 {
		if e.Catalog.Value == 0 {
			db_resp, err := db.OfferGetField(e.Offer.Value)
			if err != nil {
				err = fmt.Errorf("GetCatalogId %v", err)
				e.Errors = append(e.Errors, err)
				return 0
			}
			if db_resp.CatalogID.Valid {
				catalogId = int32(db_resp.CatalogID.Int64)
				return catalogId
			}
		}
		return e.Catalog.Value
	}
	return 0
}

// Times contains all the time related attributes of an event.
type Times struct {
	S     MaybeInt   `json:"timestamp"`
	Ms    MaybeInt64 `json:"timestampMs"`
	Pub   MaybeInt   `json:"publish"`
	Run   MaybeInt   `json:"runfrom"`
	Exp   MaybeInt   `json:"expires"`
	Range string     `json:"timeRange"`
}

// Timestamp converts input timestamps (seconds, and  Ms) to seconds
// which is what we want to store the timestamp.
func (t Times) Timestamp() float64 {
	// unix time zero
	var t0 time.Time
	t0 = time.Unix(0, 0)
	// add the int amount of second from to t0
	// note that this consider that if a filed is missing form the json
	// the value defaults to zero as a int type
	if t.Ms.Value64 == 0 {
		deltatime := t0.Add(time.Duration(t.S.Value) * time.Second)
		seconds := float64(deltatime.Unix())
		return seconds
	}
	deltatime := time.Duration(t.Ms.Value64) * time.Millisecond
	seconds := deltatime.Seconds()
	return seconds
}

// PublishTimeStamp returns either the timestamp extracted from the json blob
// or looks it up in the database dump using offer.id
func (e *Event) PublishTimeStamp(db db.Connector) (publishTimeStamp float64) {
	if e.Times.Pub.Value == 0 {
		switch e.Scope {
		case "offer":
			db, err := db.OfferGetField(e.Offer.Value)
			if err != nil {
				err = fmt.Errorf("PublishTimeStamp %v", err)
				e.Errors = append(e.Errors, err)
			}
			if db.Publish.Valid {
				publishTimeStamp = db.Publish.Float64
			}
			return publishTimeStamp
		case "catalog":
			db, err := db.CatalogGetField(e.Catalog.Value)
			if err != nil {
				err = fmt.Errorf("PublishTimeStamp %v", err)
				e.Errors = append(e.Errors, err)
			}
			if db.Publish.Valid {
				publishTimeStamp = db.Publish.Float64
			}
			return publishTimeStamp
		}
	}
	// Unix seconds as float, as per spec.
	return float64(e.Times.Pub.Value)
}

// UDuration returns either the duration extracted from the json blob, or zero and appends the error if the duration is negative
func (e *Event) UDuration() float64 {
	if e.Duration < 0 {
		e.Errors = append(e.Errors, fmt.Errorf("Negative event duration"))
		return 0
	}
	return e.Duration

}

// RunFrom returns either the timestamp extracted from the json blob
// or looks it up in the database dump using offer.id
func (e *Event) RunFrom(db db.Connector) (RunFromTimeStamp float64) {
	if e.Times.Run.Value == 0 {
		switch e.Scope {
		case "offer":
			db, err := db.OfferGetField(e.Offer.Value)
			if err != nil {
				err = fmt.Errorf("RunFrom %v", err)
				e.Errors = append(e.Errors, err)
			}
			if db.RunFrom.Valid {
				RunFromTimeStamp = db.RunFrom.Float64
			}
			return RunFromTimeStamp
		case "catalog":
			db, err := db.CatalogGetField(e.Catalog.Value)
			if err != nil {
				err = fmt.Errorf("RunFrom %v", err)
				e.Errors = append(e.Errors, err)
			}
			if db.RunFrom.Valid {
				RunFromTimeStamp = db.RunFrom.Float64
			}
			return RunFromTimeStamp
		}
	}

	// Unix seconds as float, as per spec.
	return float64(e.Times.Run.Value)
}

// Expires return the expiration date of event with id=id
func (e *Event) Expires(db db.Connector) (ExpiresTimeStamp float64) {
	if e.Times.Exp.Value == 0 {
		switch e.Scope {
		case "offer":
			db, err := db.OfferGetField(e.Offer.Value)
			if err != nil {
				err = fmt.Errorf("Expires %v", err)
				e.Errors = append(e.Errors, err)

			}
			if db.Expires.Valid {
				ExpiresTimeStamp = db.Expires.Float64
			}
			return ExpiresTimeStamp
		case "catalog":
			db, err := db.CatalogGetField(e.Catalog.Value)
			if err != nil {
				err = fmt.Errorf("Expires %v", err)
				e.Errors = append(e.Errors, err)
			}
			if db.Expires.Valid {
				ExpiresTimeStamp = db.Expires.Float64
			}
			return ExpiresTimeStamp
		}
	}
	// Unix seconds as float, as per spec.
	return float64(e.Times.Exp.Value)
}

// Location contains all the location related attributes of an event.
type Location struct {
	GeoHash   string     `json:"geohash"`
	Geocoded  int32      `json:"geocoded"`
	Lat       MaybeFloat `json:"latitude"`
	Long      MaybeFloat `json:"longitude"`
	ReqRadius float64    `json:"distance"`
}

// GeoCoded returns if the event is geocoded or the location is extracted in a different way.
func (l *Location) GeoCoded() bool {
	if l.Geocoded == 1 {
		return true
	}
	return false
}

// Latitude returns the latitude of an event
func (l Location) Latitude() float64 {
	in := l.Lat
	// handle v1 default
	if in.Value == 55.730521 {
		return 0.0
	}
	// return parsed value, 0.0 if not in json
	return in.Value
}

// Longitude returns the latitude of an event
func (l Location) Longitude() float64 {
	in := l.Long
	if in.Value == 12.464888 {
		return 0.0
	}
	// return parsed value, 0.0 if not in json
	return in.Value
}

// Geohash returns the geohas with max preicsion (12 digits)
func (l Location) Geohash() string {
	// handle v1 defalut value for nil (i.e. somewhere
	// in Herlev 55.730521,12.464888
	if l.Lat.Value == 55.730521 && l.Long.Value == 12.464888 {
		return ""
	}
	// handle nulls
	if l.Lat.Value == 0 && l.Long.Value == 0 {
		return ""
	}
	// calculate if empty and if len != full precision (12)
	if l.GeoHash == "" || len(l.GeoHash) != 12 {
		var geohash string
		geohash = libgeohash.Encode(l.Lat.Value, l.Long.Value)
		return geohash
		// return  if already exist
	}
	return l.GeoHash
}

// TODO this could be merged into type maybeint  (?)

// MaybeInt64 is a strcture that holds a value that **must** be a integer64
// but may be a string in the original data
type MaybeInt64 struct {
	Value64 int64
}

// MaybeInt is a strcture that holds a value that **must** be a integer32
// but may be a string in the original data
type MaybeInt struct {
	Value    int32
	IntArray MaybeIntArray
}

// MaybeBool is a strcture that holds a value that **must** be a bool
// but may be a string in the original data
type MaybeBool struct {
	Value bool
}

// MaybeFloat is a strcture that holds a value that **must** be a float
// but may be a string in the original data
type MaybeFloat struct {
	Value float64
}

// MaybeString is a strcture that holds a value that **must** be a string
// but may be a bool or number   in the original data
type MaybeString struct {
	Value string
}

// MaybeIntArray is a strcture that holds a value that **must** be a string
// but may be a string array or int43 array in the original data
type MaybeIntArray struct {
	IntArray []int32
}

/*
 *
 *Methods
 *
 */

// UnmarshalJSON teaches MaybeIntArray how to parse itself
func (a *MaybeIntArray) UnmarshalJSON(b []byte) (err error) {
	var intarr []int32
	var strarr []string
	// unmarshall array of ints into array of ints
	// which gives nil error if the
	// input byte stream is indeed an array of ints
	if err = json.Unmarshal(b, &intarr); err == nil {
		a.IntArray = intarr
	} else {
		var intarr []int32
		err = json.Unmarshal(b, &strarr)
		if err != nil {
			return err
		}
		for _, value := range strarr {
			intvalue, _ := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			intarr = append(intarr, int32(intvalue))
			a.IntArray = intarr
		}
	}
	return err
}

// GetArray returns the underling array from a MaybeIntArray type.
func (a *MaybeIntArray) GetArray() []int32 {
	arr := a.IntArray
	bytes := make([]int32, len(arr))
	copy(bytes, arr)
	return bytes
}

// UnmarshalJSON teaches MaybeString how to parse itself
func (a *MaybeString) UnmarshalJSON(b []byte) (err error) {
	s := ""
	// unmarshall string into string
	// which gives nil error if the
	// input byte stream is indeed a string
	if err = json.Unmarshal(b, &s); err == nil {
		a.Value = s
	} else {
		// hack to handle special cases
		if string(b) == "false" {
			// YOLO
			a.Value = ""
			err = nil
			// if input is a int
			// unmarhsall will fail
			// so cast to string
		} else {
			// NOTE this may be a bug in api
			a.Value = string(b)
			err = nil
		}
	}
	return err
}

// UnmarshalJSON teachs maybeint how to parse itself
func (a *MaybeInt) UnmarshalJSON(b []byte) (err error) {
	s, n := "", int32(0)
	var arr MaybeIntArray
	// unmarshall string into string
	// which gives nil error if the
	// input byte stream is indeed a string
	var temp int64
	if err = json.Unmarshal(b, &s); err == nil {
		// convert the string to int
		if (s == "") || (s == "null") {
			// except if it is null
			a.Value = 0
			return nil
		}
		temp, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		a.Value = int32(temp)
		return nil
	}
	// unmarshall int into int
	// which gives nil error  if the
	// input byte stream is indeed a int
	if err = json.Unmarshal(b, &n); err == nil {
		a.Value = n
		return nil
	}
	// unmarshall intarray into intarray
	// which gives nil error  if the
	// input byte stream is indeed a intarray
	if err = json.Unmarshal(b, &arr); err == nil {
		a.IntArray.IntArray = arr.GetArray()
		return nil
	}
	return err
}

// UnmarshalJSON teachs maybeint64 how to parse itself
func (a *MaybeInt64) UnmarshalJSON(b []byte) (err error) {
	s, n := "", int64(0)
	// unmarshall string into string
	// which gives nil error if the
	// input byte stream is indeed a string
	var temp int64
	if err = json.Unmarshal(b, &s); err == nil {
		// convert the string to int
		if (s == "") || (s == "null") {
			// except if it is null
			a.Value64 = 0
		} else {
			temp, err = strconv.ParseInt(s, 10, 64)
			a.Value64 = temp
		}
	} else {
		// unmarshall int into int
		// which gives nil error  if the
		// input byte stream is indeed a int
		err = json.Unmarshal(b, &n)
		a.Value64 = n
	}
	return err
}

// UnmarshalJSON teachs maybebool how to parse itself
func (a *MaybeBool) UnmarshalJSON(b []byte) (err error) {
	str, bol := "", false
	// unmarshall string into string
	// which gives nil error if the
	// input byte stream is indeed a string
	if err = json.Unmarshal(b, &str); err == nil {
		// convert the string to bool, strip spaces else fails
		// empty string means false by spec
		if len(str) < 1 {
			a.Value = false
			err = nil
		} else {
			a.Value, err = strconv.ParseBool(strings.Trim(str, " "))
		}
	} else {
		// unmarshall bool into bool
		// which gives nil error  if the
		// input byte stream is indeed a int
		err = json.Unmarshal(b, &bol)
		a.Value = bol
	}
	return err
}

// UnmarshalJSON teachs maybefloat how to parse itself
func (a *MaybeFloat) UnmarshalJSON(b []byte) (err error) {
	s, n := "", float64(0)
	// unmarshall string into string
	// which gives nil error  if the
	// input byte stream is indeed a string
	if err = json.Unmarshal(b, &s); err == nil {
		// empty string is null
		if s == "" {
			a.Value = 0
			// convert string to float
		} else {
			a.Value, err = strconv.ParseFloat(s, 64)
		}
	} else {
		// unmarshall float into float
		// which gives nil error  if the
		// input byte stream is indeed a int
		err = json.Unmarshal(b, &n)
		a.Value = n
	}
	return err
}

/*
 *
 *functions
 *
 */

// intToIP converts a integer rapresentation of an ip to the
// corresponding net.IP rapresentation,  works only for ipv4.
// as desired since the phps ip2long only works for ipv4.
func intToIP(ip uint32) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ip & 0xFF)
	bytes[1] = byte((ip >> 8) & 0xFF)
	bytes[2] = byte((ip >> 16) & 0xFF)
	bytes[3] = byte((ip >> 24) & 0xFF)
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

// Ipv4Ipv6 performs the conversion from ipv4 to ipv6
func (e *Event) Ipv4Ipv6(ipv4 uint32) string {
	var ip4 net.IP
	// handle the empty field from json
	if ipv4 == 0 {
		return ""
	}
	// this an ipv4 because of php api
	// but NOTE that internally it may be also
	// ipv6
	ip4 = intToIP(ipv4)
	if ip4 == nil {
		// note this actualy will never happen
		// an unit will alwasy be convetend into a valid ip
		// it may not be the correct ip.
		err := fmt.Errorf("IPV4 string conversion error")
		e.Errors = append(e.Errors, err)
		return ""
	}
	ip4stirng := ip4.String()
	ipv6string := fmt.Sprintf("::ffff:%v", ip4stirng)
	return ipv6string
}

// ApiAppVersion returns the string of the version identifer or nil string
// if not defined ("", or "null")
func (e Event) ApiAppVersion() string {
	if len(e.Api_app_version) < 1 {
		return ""
	} else if e.Api_app_version == "null" {
		return ""
	} else {
		return e.Api_app_version
	}
}

// GetPrice returns the value of the price of an offer with id=id
func (e *Event) GetPrice(db db.Connector) (offerPrice float64) {
	if e.Price.Value == 0 {
		offer, err := db.OfferGetField(e.Offer.Value)
		if offer.Price.Valid {
			offerPrice = offer.Price.Float64
		}
		if err != nil {
			err = fmt.Errorf("GetPrice:%v", err)
			e.Errors = append(e.Errors, err)
		}
		return float64(offerPrice)
	} // float, as per spec
	return e.Price.Value
}

// GetPrePrice returns the value of the preprice of an offer with id=id
func (e *Event) GetPrePrice(db db.Connector) (offerPrePrice float64) {
	if e.PrePrice.Value == 0 {
		offer, err := db.OfferGetField(e.Offer.Value)
		if offer.PrePrice.Valid {
			offerPrePrice = offer.PrePrice.Float64
		}
		if err != nil {
			err = fmt.Errorf("GetPrePrice:%v", err)
			e.Errors = append(e.Errors, err)
		}
		return float64(offerPrePrice)

	} // float, as per spec
	return e.PrePrice.Value
}

//ApiEnv return the description of the enviroment api (edge prod dev..)
func (e Event) ApiEnv() string {
	if len(e.Apienv) < 1 || e.Apienv == "null" {
		return "production"
	}
	return e.Apienv
}

// GetDealers makes sure we can extract a list of dealer id
// also in the casee case wehre the informatio nwas not correctly
// stored into the dealers field but in the dealer field
// this was a bug in the bridge between api 1.0 and api 2.0
func (e Event) GetDealers() []int32 {
	if e.Dealer.IntArray.IntArray != nil {
		// this handles the corner case
		// where the event has "delaer = [1,3,3]"
		// make sure we use intarray
		return e.Dealer.IntArray.GetArray()
	}
	return e.Dealers.GetArray()
}

// AppGroups return a list of the groups an app_id belongs to
func (e *Event) AppGroups(db db.Connector) []int32 {
	appID := e.App_id.Value
	if appID != 0 {
		appgroups, err := db.AppGetField(appID)
		if err != nil {
			err = fmt.Errorf("AppGroups:%v", err)
			e.Errors = append(e.Errors, err)
		}
		return appgroups
	}
	return []int32{}
}

// Type returns a scope.action
func (e Event) Type() (eventType string) {
	if e.Action == "" {
		return ""
	}
	if e.Scope == "" {
		return ""
	}
	eventType = fmt.Sprintf("%v.%v", e.Scope, e.Action)
	return eventType
}

// CheckCurrency returns the currency of the offer price and preprice
// we enforce iso values (i.e. longer than 1 shorter than 3)
func CheckCurrency(in string) (out string) {
	if in == "" {
		out = "DKK"
	}
	if in == "1" {
		out = "DKK"
	}
	if len(in) > 1 && len(in) < 3 {
		out = ""
	}
	return
}

/*
*
* Reflection methods.
*  This will fail if the interface is nil.
*  which is fine because we don't want an nil event anyway.
 */

// ZeroCheck returns an array containg true if field is zero of type
// or  false is not zero of the type
// TODO panics if f is not  fileds are not exported
// TODO panics if f is not  struct.
func ZeroCheck(f interface{}) []bool {
	val := reflect.Indirect(reflect.ValueOf(f))
	var values []bool
	var cond bool
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		switch valueField.Interface().(type) {
		default:
			cond = valueField.Interface() == reflect.Zero(typeField.Type).Interface()
			values = append(values, cond)
		case []int32:
			value, _ := valueField.Interface().([]int32)
			if len(value) == 0 {
				cond = true
			}
			values = append(values, cond)

		case MaybeIntArray:
			value, _ := valueField.Interface().(MaybeIntArray)
			if len(value.IntArray) == 0 {
				cond = true
			}
			values = append(values, cond)
		}
	}
	return values
}

// GetNames returns an array containg the names of the
// field of the struct
func GetNames(f interface{}) []string {
	val := reflect.Indirect(reflect.ValueOf(f))
	var values []string
	for i := 0; i < val.NumField(); i++ {
		nameField := val.Type().Field(i).Name
		values = append(values, nameField)
	}
	return values
}

// FieldZeroCheck returns true if a field it's zero  false if it's not zero.
func FieldZeroCheck(name string, f interface{}) bool {
	val := reflect.Indirect(reflect.ValueOf(f))
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if typeField.Name == name {
			switch valueField.Interface().(type) {
			default:
				cond := valueField.Interface() == reflect.Zero(typeField.Type).Interface()

				return cond
			case []int32:
				value, _ := valueField.Interface().([]int32)
				//NOTE
				if len(value) == 0 {
					cond := true
					return cond
				}
			case MaybeIntArray:
				value, _ := valueField.Interface().(MaybeIntArray)
				if len(value.IntArray) == 0 {
					cond := true
					return cond
				}
			case MaybeInt:
				value := valueField.Interface().(MaybeInt)
				cond := value.Value == reflect.Zero(typeField.Type).Interface()
				return cond
			}
		}
	}
	return false
}

// GetValue returns the value of the field  "name"
// return nil + error if there is no such field
func GetValue(name string, f interface{}) (interface{}, error) {
	val := reflect.Indirect(reflect.ValueOf(f))
	var cond interface{}
	var err error
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if typeField.Name == name {
			err = nil
			cond = valueField.Interface()
			return cond, err
		}
	}
	cond = nil
	err = fmt.Errorf("Wrong field name or non-existing field. Did you mispell the name?")
	return cond, err
}

// TestEq checks if two int32 slices are the same
func TestEq(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return true
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
