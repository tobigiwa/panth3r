package env

import (
	"fmt"
	"sync"

	"github.com/Netflix/go-env"
)

type Environment struct {
	// Databases struct {
	// 	Mongo struct {
	// 		Host       string `env:"MONGO_HOST,required=true"`
	// 		Username   string `env:"MONGO_USERNAME,required=true" `
	// 		Password   string `env:"MONGO_PASSWORD,required=true" json:"-" yaml:"-" toml:"-"`
	// 		Database   string `env:"MONGO_DATABASE,default=waitlist"`
	// 		Collection string `env:"MONGO_COLLECTION,default=users"`
	// 	}
	// }

	Server struct {
		Port    string `env:"PORT,default=8090"`
		Env     string `env:"ENV,default=development"`
		Version string `env:"VERSION"`
	}

	Mail struct {
		EmailAcc            string `env:"EMAILACC,required=true"`
		EmailPswd           string `env:"EMAILPSWD,required=true"`
		EmailSmtpServerHost string `env:"SMTPHOST,required=true"`
		EmailSmtpServerPort int    `env:"SMTPPORT,required=true"`
	}
}

var (
	once        sync.Once
	environment Environment
)

// LoadAllEnvVars loads required env variables. LoadAllEnvVars contains a sync.Once
func LoadAllEnvVars() {
	once.Do(func() {
		if _, err := env.UnmarshalFromEnviron(&environment); err != nil {
			panic(fmt.Errorf("error: env.UnmarshalFromEnviron(&environment): %v", err))
		}
	})
}

// GetEnvVar gets a particular env. variable and should only be
// called after LoadAllEnvVars() has be called.
func GetEnvVar() Environment {
	return environment
}

func BuildURI(username, password, host string) string {
	return "mongodb+srv://" + username + ":" + password + "@" + host + "/?retryWrites=true&w=majority"
}

// SetSaneDefaults sets defaults value
//
//	if environment.Server.Env == "" {
//		environment.Server.Env = "development"
//	}
//
//	if environment.Databases.Mongo.Database == "" {
//		environment.Databases.Mongo.Database = "waitlist"
//	}
//
//	if environment.Databases.Mongo.Collection == "" {
//		environment.Databases.Mongo.Collection = "users"
//	}
//
//	if environment.Server.Port == "" {
//		environment.Server.Port = "8090"
//	}
// func SetSaneDefaults() {
// 	if environment.Server.Env == "" {
// 		environment.Server.Env = "development"
// 	}

// 	if environment.Databases.Mongo.Database == "" {
// 		environment.Databases.Mongo.Database = "waitlist"
// 	}
// 	if environment.Databases.Mongo.Collection == "" {
// 		environment.Databases.Mongo.Collection = "users"
// 	}
// 	if environment.Server.Port == "" {
// 		environment.Server.Port = "8090"
// 	}
// }
