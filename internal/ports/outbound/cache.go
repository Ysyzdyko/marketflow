package outbound

type Redis interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	MGet(keys ...string) ([]interface{}, error)
	Keys(pattern string) ([]string, error)
	FlushAll() error
}
