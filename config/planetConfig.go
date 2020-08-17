package config

type Config struct {
	Port     int    `env:"PORT" envDefault:"8080"`
	DBURI    string `env:"DB_URI" envDefault:"mongodb://localhost:27017"`
	SWAPIURL string `env:"SWAPI_URL" envDefault:"https://swapi.dev/api"`
}
