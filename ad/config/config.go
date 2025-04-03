package config

const (
	ServerURI       = ":50051"
	ServerPrintURI  = "localhost:50052"
	RedisURI        = "localhost:6379"
	MongoURI        = "mongodb://localhost:27017"
	DBName          = "addb"
	CollectionName  = "ads"
	ExpirationDelay = 30 * 24 * 60 * 60
	CacheDelay      = 5 * 60
)
