package config

/*
   |--------------------------------------------------------------------------
   | Default Database Connection Name
   |--------------------------------------------------------------------------
   |
   | Here you may specify which of the database connections below you wish
   | to use as your default connection for all database work. Of course
   | you may use many connections at once using the Database library.
   |
*/

type DBConnectionNameConfig struct {
	Default string `env:"DB_CONNECTION" envDefault:"mysql"`
}

type DBSqliteConfig struct {
	Url                   string `env:"DATABASE_URL" envDefault:"production"`
	Database              string `env:"DB_DATABASE" envDefault:"database.sqlite"`
	Prefix                string `env:"DB_DATABASE_PREFIX" envDefault:""`
	ForeignKeyConstraints bool   `env:"DB_FOREIGN_KEYS" envDefault:"true"`
}
