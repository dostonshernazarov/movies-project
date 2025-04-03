package config

const (
	// DefaultPort is the default port for the server
	DefaultPort = "8060"
)

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}
