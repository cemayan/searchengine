package constants

type Db int

const (
	Redis Db = iota + 1
	MongoDb
)

var Db2Str = map[Db]string{Redis: "redis", MongoDb: "mongodb"}
var Str2Db = map[string]Db{"redis": Redis, "mongodb": MongoDb}

type DbType int

const (
	Read DbType = iota + 1
	Write
)

var DbTypeMap = map[DbType]string{Read: "read", Write: "write"}

type Project int

const (
	ReadApi Project = iota + 1
	WriteApi
)

var ProjectMap = map[Project]string{ReadApi: "read", WriteApi: "write"}

const (
	RedisJsonPrefix = "record"
)
