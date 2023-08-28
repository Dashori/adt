package redisrepo

type RedisRepository interface {
	Get(string) (string, error)
	Set(string, string) error
	Del(string) error
}
