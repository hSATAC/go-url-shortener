package shortener

type RedisConfig struct {
	address           string
	namespace         string
	maxIdleConn       int
	maxIdleTimeoutSec int
}
