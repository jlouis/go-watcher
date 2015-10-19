//Package db implements  the connection and querying infrastructure
//to our mysql database.
package db

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shopgun/matilde/env"
	"sync"
	"time"
)

const (
	/*
	 *query for a single id
	 */
	offerGetFieldQuery = `
	SELECT
		secureid,
		UNIX_TIMESTAMP(publish) as publish,
		UNIX_TIMESTAMP(runfrom) as runfrom,
		UNIX_TIMESTAMP(expires) as expires,
		price,
		preprice,
		catalog
	FROM main.offer
	WHERe id = ?`
	catalogGetFieldQuery = `
	SELECT secureid,
		UNIX_TIMESTAMP(publish) as publish,
		UNIX_TIMESTAMP(runfrom) as runfrom,
		UNIX_TIMESTAMP(expires) as expires
	FROM main.catalog
	WHERE id = ? `
	userGetFieldQuery = `
	SELECT
		IFNULL(gender,0) as gender,
		IFNULL(birthyear,0)
	FROM main.user
	WHERE id = ? `
	appGetFieldQuery = "SELECT `group` FROM main.appgroupmemberships WHERE app = ?"
	/*
	 *queries to perform a db dump
	 */
	DumpOffersQuery = `
	SELECT
		secureid,
		UNIX_TIMESTAMP(publish) as publish,
		UNIX_TIMESTAMP(runfrom) as runfrom,
		UNIX_TIMESTAMP(expires) as expires,
		id,
		price,
		preprice,
		catalog
	FROM main.offer
	`
	DumpCatalogsQuery = `
	SELECT secureid,
		UNIX_TIMESTAMP(publish) as publish,
		UNIX_TIMESTAMP(runfrom) as runfrom,
		UNIX_TIMESTAMP(expires) as expires,
		id
	FROM main.catalog
	`
	DumpUsersQuery = `
	SELECT
		IFNULL(gender,0) as gender,
		IFNULL(birthyear,0) as birthyear,
		id
	FROM main.user
	`
	DumpAppsQuery = "SELECT `group`, app FROM main.appgroupmemberships"
)

// Connector wraps around a sql DB pointer to expose  querying methods
// NOTE danger ! don't close else need to prepare once more even though the connection is potentially ok
//defer c.prepOfferQuery.Close()
type Connector struct {
	DB               *sqlx.DB
	prepOfferQuery   *sql.Stmt
	prepAppQuery     *sql.Stmt
	prepCatalogQuery *sql.Stmt
	prepUserQuery    *sql.Stmt
	Offers           map[int32]Offer
	Catalogs         map[int32]Catalog
	Groups           map[int32][]int32
	Users            map[int32]User
	NeedsUpdate      bool
	sync.RWMutex
}

//Connect to mysqldb using the  info/credentail stored in the env.
func Connect(e env.VARS) (c Connector, err error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%v:%v)/%s", e.DbUser, e.DbPwd, e.DbAddres, e.DbPort, e.DbName))
	if err != nil {
		return c, err
	}
	c.DB = db
	db.SetMaxIdleConns(e.DbMaxIdleConns)
	db.SetMaxOpenConns(e.DbMaxOpenConns)
	// prepare the queries
	offer, err := db.Prepare(offerGetFieldQuery)
	if err != nil {
		return c, err
	}
	c.prepOfferQuery = offer
	app, err := db.Prepare(appGetFieldQuery)
	if err != nil {
		return c, err
	}
	c.prepAppQuery = app
	user, err := db.Prepare(userGetFieldQuery)
	if err != nil {
		return c, err
	}
	c.prepUserQuery = user
	catalog, err := db.Prepare(catalogGetFieldQuery)
	if err != nil {
		return c, err
	}
	c.prepCatalogQuery = catalog
	return c, nil
}

// Offer describes the offer event retrieved from csv dump of mysqldb
type Offer struct {
	CatalogID sql.NullInt64   `db:"catalog"`
	Expires   sql.NullFloat64 `db:"expires"`
	PrePrice  sql.NullFloat64 `db:"preprice"`
	Price     sql.NullFloat64 `db:"price"`
	Publish   sql.NullFloat64 `db:"publish"`
	RunFrom   sql.NullFloat64 `db:"runfrom"`
	Secureid  sql.NullString  `db:"secureid"`
	Id        sql.NullInt64   `db:"id"`
}

// User describes the user entry retrieved from db
type User struct {
	Byear  sql.NullInt64 `db:"birthyear"`
	Gender sql.NullInt64 `db:"gender"`
	Id     sql.NullInt64 `db:"id"`
}

// Catalog describes the catalog entry retrieved from db
type Catalog struct {
	Expires  sql.NullFloat64 `db:"expires"`
	Publish  sql.NullFloat64 `db:"publish"`
	RunFrom  sql.NullFloat64 `db:"runfrom"`
	Secureid sql.NullString  `db:"secureid"`
	Id       sql.NullInt64   `db:"id"`
}

//App describes the app entry retreived from db
type App struct {
	Group  sql.NullInt64 `db:"group"`
	Id     sql.NullInt64 `db:"app"`
	Groups []int32
}

// OfferGetField  loads an offer with offerid from the db
func (c Connector) OfferGetField(offerid int32) (offer Offer, err error) {
	c.RLock()
	// if in memory just use it
	if offer, ok := c.Offers[offerid]; ok {
		return offer, nil
	}
	c.RUnlock()
	c.NeedsUpdate = true
	err = c.prepOfferQuery.QueryRow(offerid).Scan(&offer.Secureid, &offer.Publish, &offer.RunFrom, &offer.Expires, &offer.Price, &offer.PrePrice, &offer.CatalogID)
	switch {
	case err != nil:
		return offer, fmt.Errorf("%v for offerid: %v", err, offerid)
	default:
		return offer, nil
	}
}

// UserGetField loads an user with userid from the db
func (c *Connector) UserGetField(userid int32) (user User, err error) {
	c.RLock()
	// if in memory just use it
	if user, ok := c.Users[userid]; ok {
		return user, nil
	}
	c.RUnlock()
	// else query and mark as need update
	c.NeedsUpdate = true
	err = c.prepUserQuery.QueryRow(userid).Scan(&user.Gender, &user.Byear)
	switch {
	case err != nil:
		return user, fmt.Errorf("%v for userid: %v", err, userid)
	default:
		return user, nil
	}
}

// CatalogGetField loads an user with userid from the db
func (c Connector) CatalogGetField(catalogid int32) (catalog Catalog, err error) {
	c.RLock()
	// if in memory just use it
	if catalog, ok := c.Catalogs[catalogid]; ok {
		return catalog, nil
	}
	c.RUnlock()
	// else query and mark as need update
	c.NeedsUpdate = true
	err = c.prepCatalogQuery.QueryRow(catalogid).Scan(&catalog.Secureid, &catalog.Publish, &catalog.RunFrom, &catalog.Expires)
	switch {
	case err != nil:
		return catalog, fmt.Errorf("%v for catalogid: %v", err, catalogid)
	default:
		return catalog, nil
	}
}

//DumpOffers dumps all the information from the db
func (c *Connector) DumpOffers() error {
	data := make(map[int32]Offer)
	pp := []Offer{}
	err := c.DB.Select(&pp, DumpOffersQuery)
	if err != nil {
		return err
	}
	for _, i := range pp {
		data[int32(i.Id.Int64)] = i
	}
	c.Lock()
	c.Offers = data
	c.Unlock()
	return nil
}

//DumpCatalogs dumps all the information from the db
func (c *Connector) DumpCatalogs() error {
	data := make(map[int32]Catalog)
	pp := []Catalog{}
	err := c.DB.Select(&pp, DumpCatalogsQuery)
	if err != nil {
		return err
	}
	for _, i := range pp {
		data[int32(i.Id.Int64)] = i
	}
	c.Lock()
	c.Catalogs = data
	c.Unlock()
	return nil
}

//DumpUsers dumps all the information from the db
func (c *Connector) DumpUsers() error {
	data := make(map[int32]User)
	pp := []User{}
	err := c.DB.Select(&pp, DumpUsersQuery)
	if err != nil {
		return err
	}
	for _, i := range pp {
		data[int32(i.Id.Int64)] = i
	}
	c.Lock()
	c.Users = data
	c.Unlock()
	return nil
}

// AppGetField loads an app with appid from the db
// NOTE that as of 14 Oct 2015	there are no apps
// with multiple groups, and AFAIK one can't really store
// arrays in mysql.
func (c *Connector) AppGetField(appid int32) (groups []int32, err error) {
	// if in memory just use it
	c.RLock()
	if app, ok := c.Groups[appid]; ok {
		return app, nil
	}
	c.RUnlock()
	// else query and mark as need update
	c.NeedsUpdate = true
	rows, err := c.prepAppQuery.Query(appid)
	for rows.Next() {
		var group int32
		err = rows.Scan(&group)
		groups = append(groups, group)
	}
	rows.Close()
	switch {
	case err != nil:
		if err == sql.ErrNoRows {
			// we have not app-id group relationship. return []
			return []int32{}, nil
		}
		return []int32{}, err
	default:
		return groups, nil
	}
}

//DumpApps dumps all the information from the db
func (c *Connector) DumpApps() error {
	data := make(map[int32][]int32)
	pp := []App{}
	err := c.DB.Select(&pp, DumpAppsQuery)
	if err != nil {
		return err
	}
	for _, i := range pp {
		if val, ok := data[int32(i.Id.Int64)]; ok {
			val = append(val, int32(i.Group.Int64))
			data[int32(i.Id.Int64)] = val
		} else {
			data[int32(i.Id.Int64)] = []int32{int32(i.Group.Int64)}
		}
	}
	// make sure is thread safe
	c.Lock()
	c.Groups = data
	c.Unlock()
	return nil
}

// Update checks if the dump needs updating and updates the db in the baground
func (c *Connector) Update() chan struct{} {
	updated := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				if c.NeedsUpdate {
					err := c.DumpApps()
					log.Debugf("Loading apps")
					if err != nil {
						log.Error(err)
					}
					err = c.DumpUsers()
					log.Debugf("Loading users")
					if err != nil {
						log.Error(err)
					}
					err = c.DumpCatalogs()
					log.Debugf("Loading catalogs")
					if err != nil {
						log.Error(err)
					}
					err = c.DumpOffers()
					log.Debugf("Loading offers")
					if err != nil {
						log.Error(err)
					}
					updated <- struct{}{}
				}
				updated <- struct{}{}
				c.NeedsUpdate = false
			}

		}
	}()
	return updated
}
