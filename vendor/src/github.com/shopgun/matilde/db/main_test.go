package db

import (
	"github.com/shopgun/matilde/env"
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

/*
 *Tests
 */
var e env.VARS
var connection Connector
var doneupdating chan struct{}

func init() {
	e = env.Init()
	log.SetLevel(log.DebugLevel)
	// NOTE slip db stuff on travis for now
	if os.Getenv("travis") != "" {
		e.DbName = ""
		// we can then set up mysql travis clone
	} else {
		var err error
		connection, err = Connect(e)
		if err != nil {
			panic(err)
		}
		err = connection.DB.Ping()
		if err != nil {
			panic(err)
		}
	}
}

func TestConnect(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	connection, err := Connect(e)
	if err != nil {
		t.Error(err)
	}
	err = connection.DB.Ping()
	if err != nil {
		t.Error(err)
	}
}

// TestGroup tests wheat happens when we have no appgroups
// we expect a empty result and error
func TestGroup(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	err := connection.DB.Ping()
	if err != nil {
		t.Error(err)
	}
	res, err := connection.AppGetField(999)
	if len(res) != 0 {
		t.Errorf("got %v, wanted %v", res, 0)
	}
	if err != nil {
		t.Error(err)
	}
	val, err := connection.AppGetField(25)
	if err != nil {
		t.Error(err)
	}
	if val[0] != 8 {
		t.Errorf("wanted group 8")
	}
	val, err = connection.AppGetField(23)
	if err != nil {
		t.Error(err)
	}
	if val[0] != 8 {
		t.Errorf("wanted group 8")
	}
	val, err = connection.AppGetField(22)
	if err != nil {
		t.Error(err)
	}
	if val[0] != 8 {
		t.Errorf("wanted group 8")
	}
	if val[1] != 19 {
		t.Errorf("wanted group 19")
	}
}

// TestOffer chekc that we don't get an error
// and we get correct results for one query
func TestOffer(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	offer, err := connection.OfferGetField(1922)
	if err != nil {
		t.Error(err)
	}
	if offer.Secureid.String != "5d59KN" {
		t.Errorf("Got %v, wanted 5d59KN", offer.Secureid)
	}
}

func TestUser(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	user, err := connection.UserGetField(31)
	if err != nil {
		t.Error(err)
	}
	if user.Gender.Int64 != 1 {
		t.Errorf("Got %v, wanted 1", user.Gender)
	}
	if user.Byear.Int64 != 1989 {
		t.Errorf("Got %v, wanted 1989", user.Byear)
	}
	//test nils
	user, err = connection.UserGetField(5149)
	if err != nil {
		t.Error(err)
	}
}

func TestCatalog(t *testing.T) {
	doneupdating = connection.Update()
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	catalog, err := connection.CatalogGetField(186)
	if err != nil {
		t.Error(err)
	}
	if catalog.Secureid.String != "87bfKm" {
		t.Errorf("Got %v, wanted 1", catalog.Secureid)
	}
	if catalog.Expires.Float64 != 1270944000 {
		t.Errorf("Got %v, wanted 1270944000", catalog.Expires)
	}
	t.Log("just waiting baby")
	<-doneupdating
}

func TestGroupDump(t *testing.T) {
	if e.DbName == "" {
		t.Skip("Skipping DB Tests, $DB_NAME not set")
	}
	//load dumps
	doneupdating = connection.Update()
	<-doneupdating
	// now get the results
	val, err := connection.AppGetField(22)
	t.Log(val)
	if err != nil {
		t.Error(err)
	}
	if val[0] != 8 {
		t.Errorf("wanted group 8")
	}
	if val[1] != 19 {
		t.Errorf("wanted group 19")
	}
}

/*
 *Benchmaks
 */

func fromdb() []int32 {
	rows, _ := connection.AppGetField(999)
	return rows
}

func BenchmarkInDb(b *testing.B) {
	// run the fromdb function b.N times
	for n := 0; n < b.N; n++ {
		fromdb()
	}
}

func BenchmarkOfferDb(b *testing.B) {
	for n := 0; n < b.N; n++ {
		offer, _ := connection.OfferGetField(1922)
		_ = offer.Secureid
	}
}
