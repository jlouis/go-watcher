
Matilide  [![Build Status](https://magnum.travis-ci.com/shopgun/matilde.svg?token=H7MjHi74teZgv8JHTYhx&branch=feature/future)](https://magnum.travis-ci.com/shopguncom/matilde)
=========
[![forthebadge](http://forthebadge.com/images/badges/built-by-hipsters.svg)](http://forthebadge.com)


Matilde is is eTilbudsavis cleaner. Written in go to avoid type problems, and use concurrency (and parallelism) easily.

### Changes from go-dev/master (which are in "production")

No more sorting. It's too slow to do properly and we don't care. If sorting needs to be done it has to be done in the logshipper or
when events are streamed to matilde with a sliding window (i.e. 10 seconds).

Matilde/future is just a humble and transparent  cleaner: you feed her dirty stuff and get back clean stuff.

#Logic

The conversion logic is specified in [doc.md](./doc.md). Which is used as a guide with pseudo code to follow what happens to old events.
Will be obsolete with api changes.

if input json is:

```json
{
"a":0,
"Z": "0",
"b":"foo",
"c":"bar"
}
```

then :

 - the types are checked and converted
 - if a field needs to be **nuked** is not declared in the go struct that unmarshals the json blob
 - if a field needs to be **inserted** either for future proofness or because it's missing in old events but present in recent events it will be inserted and its value is defaulted to the zero of the type we want it to be.

```json
{
"a":0,
"Z":0,
"b":"foo",
"k":"string default is empty string",
"w": "int default is 0",
"y": "float default is 0.0",
}
```

then:

- transformation are performed on the fields and on their names
- if a sane missing value or default can be inserted either arbitrarily (i.e. api <1 ) or via db lookup it's done here.

```json
{
"a":0,
"aBetterName":0,
"b":"Foo",
"k":"looked up value in a db",
"w":0 ,
"y":0.0
}
```

#architecture:
## Environment
Matilde is configured via env variables.
Core variables are baked into env.go.

```go
type VARS struct {
	CPU            int
	DbAddres       string
	DbMaxIdleConns int
	DbMaxOpenConns int
	DbName         string
	DbPort         int
	DbPwd          string
	DbUser         string
	Level          log.Level
	Stream         bool
}
```
Other are loaded and defined on a per package basis for now.
To signal the test suites that the testing is done on travis we set f.ex

- travis=true

To set up aws we use f.ex:

- AWS_ACCESS_KEY=unset
- AWS_SECRET_KEY=unset




## Data Flow

- normal operation: stream from erlang port.
- reload/transform: stream/read from old data (needs infrastructural work-from msgpack).
- reclean: read and stream from old json data.


# Requirements
- Go
- A properly set up environment
- A live mysql connection to our db.
