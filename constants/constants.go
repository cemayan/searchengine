package constants

const (
	XSearchEngineQuery = "X-SearchEngine-Query"
)

// Db represents of databases
type Db int

const (
	Redis Db = iota + 1
	MongoDb
)

var Db2Str = map[Db]string{Redis: "redis", MongoDb: "mongodb"}
var Str2Db = map[string]Db{"redis": Redis, "mongodb": MongoDb}

// DbType represents of purpose
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
	Scheduler
	Scraper
)

var ProjectMap = map[Project]string{ReadApi: "read", WriteApi: "write", Scheduler: "scheduler", Scraper: "scraper"}

const (
	RecordMetaDataPrefix   = "recordmetadata"
	RecordPrefix           = "record"
	RecordDatabase         = "searchengine"
	RecordMetaDataDatabase = "searchenginemetadata"
)

// DbName represents of table/collection
type DbName int

const (
	Record DbName = iota + 1
	RecordMetadata
)

var DbName2Str = map[DbName]string{Record: "record", RecordMetadata: "recordmetadata"}
