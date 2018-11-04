package environment

import "fmt"

// Postgres  defines the parameters needed to connect to a Postgres database
type Postgres struct {
	Host     string `envconfig:"default=127.0.0.1"`
	Port     int    `envconfig:"default=5432"`
	Username string `envconfig:"default=postgres"`
	Password string `envconfig:"default=postgres"`
	Name     string `envconfig:"default=shorter"`

	MaxOpenConns int  `envconfig:"default=20"`
	LogMode      bool `envconfig:"default=false"`
}

// ConnectionString returns the connection string for a postgres database
func (p Postgres) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.Host, p.Port, p.Username, p.Password, p.Name,
	)
}

