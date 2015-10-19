// Package env takes care of setting up the environment.
// If a the local development environment variable is not set to true
// it loads from the environment else it defaults to the values
// defined in the const.
package env

import (
	"os"
	"runtime"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

// Fullback local constants
const (
	CPU            = 16
	DbAddres       = "localhost"
	DbMaxIdleConns = 8
	DbMaxOpenConns = 20
	DbName         = "main"
	DbPort         = 3306
	DbPwd          = "fmandaudialvergackle"
	DbUser         = "giulio"
)

// VARS contains the environment variables we want
// we can use this to set up matilde the way we want
type VARS struct {
	CPU            int
	DbAddres       string
	DbMaxIdleConns int
	DbMaxOpenConns int
	DbName         string
	DbPort         int
	DbPwd          string
	DbUser         string
}

func convStringToInt(s string) int {
	if s == "" {
		return 0
	}
	res, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0
	}
	return int(res)
}

// getFromLocalEnv loadds from local shell environment
// if one of the variables is nil it assumes that the
// environment is bad and falls back constants
func getFromLocalEnv() (v *VARS) {
	v = &VARS{
		CPU:            convStringToInt(os.Getenv("CPU")),
		DbAddres:       os.Getenv("DbAddress"),
		DbMaxIdleConns: convStringToInt(os.Getenv("DbMaxIdleConns")),
		DbMaxOpenConns: convStringToInt(os.Getenv("DbMaxOpenConns")),
		DbName:         os.Getenv("DbName"),
		DbPort:         convStringToInt(os.Getenv("DbPort")),
		DbPwd:          os.Getenv("DbPwd"),
		DbUser:         os.Getenv("DbUser"),
	}
	return v
}

// check make sure we control and panic if we loaded crap.
func (v VARS) check() {
	if v.DbMaxOpenConns == 0 && v.DbName != "" {
		log.Panic("Setting unlimited max connections will harras the db too much")
	}
}

// loads from consts
func localDev() (v *VARS) {
	v = &VARS{
		CPU:            CPU,
		DbAddres:       DbAddres,
		DbMaxIdleConns: DbMaxIdleConns,
		DbMaxOpenConns: DbMaxOpenConns,
		DbName:         DbName,
		DbPort:         DbPort,
		DbPwd:          DbPwd,
		DbUser:         DbUser,
	}
	return v
}

// Init initializes the environment
func Init() VARS {
	// check if we are developing locally
	localdev, err := strconv.ParseBool(os.Getenv("localdevelopment"))
	if err == nil {
		if localdev == true {
			v := localDev()
			v.check()
			runtime.GOMAXPROCS(v.CPU)
			return *v
		}
	}
	v := getFromLocalEnv()
	v.check()
	runtime.GOMAXPROCS(v.CPU)
	return *v
}
