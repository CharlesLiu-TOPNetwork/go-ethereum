package util

type ChainNetwork uint8

const (
	Dev ChainNetwork = iota
	Galileo
	NewHorizon
)

const (
	DBDialectMysql   = "mysql"
	DBDialectSqlite3 = "sqlite3"
)

var Network = Dev
