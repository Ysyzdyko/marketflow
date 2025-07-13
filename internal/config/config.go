package config

import "os"

type Config struct {
	// PostgreSQL
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	// PostgresDSN      string
	// Test DB
	TestDBHost string
	TestDBPort string
	TestDBName string

	// Redis
	RedisHost string
	RedisPort string
}

type RedisConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (r *RedisConfig) ConnRedString() string {
	return "host=" + r.Host +
		"port=" + r.Port
}

func (db *DBConfig) ConnDBString() string {
	return "host=" + db.Host +
		" port=" + db.Port +
		" user=" + db.User +
		" password=" + db.Password +
		" dbname=" + db.DBName +
		" sslmode=disable"
}

func Load() *Config {
	return &Config{
		PostgresHost:     getEnv("POSTGRES_HOST", "postgres"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", ""),
		PostgresDB:       getEnv("POSTGRES_DB", "marketflow"),

		TestDBHost: getEnv("TEST_DB_HOST", "postgres_test"),
		TestDBPort: getEnv("TEST_DB_PORT", "5433"),
		TestDBName: getEnv("TEST_DB_NAME", "marketflow_test"),

		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (c *Config) RedisConfigInitial() *RedisConfig {
	return &RedisConfig{
		Host: c.RedisHost,
		Port: c.RedisPort,
	}
}

func (c *Config) DBConfigByMode(mode bool) *DBConfig {
	if mode {
		// Live mode
		return &DBConfig{
			Host:     c.PostgresHost,
			Port:     c.PostgresPort,
			User:     c.PostgresUser,
			Password: c.PostgresPassword,
			DBName:   c.PostgresDB,
		}
	} else {
		// Test mode
		return &DBConfig{
			Host:     c.TestDBHost,
			Port:     c.TestDBPort,
			User:     c.PostgresUser,
			Password: c.PostgresPassword,
			DBName:   c.TestDBName,
		}
	}
}
