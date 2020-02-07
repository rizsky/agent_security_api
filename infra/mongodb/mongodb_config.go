package mongodb

import "fmt"

// Config stores MongoDB-related configuration
type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

// ConnectionURI :nodoc:
func (config Config) ConnectionURI() string {
	return fmt.Sprintf("mongodb://%s:%s", config.Host, config.Port)
}
