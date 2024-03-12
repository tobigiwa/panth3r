package http

import (
	"panth3rWaitlistBackend/internal/env"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

const (
	Panth3rHOMEURL string = "https://www.Panth3r.xyz/"
	Production     string = "production"
)

func (a Application) Routes() *chi.Mux {

	r := chi.NewRouter()

	if env.GetEnvVar().Server.Env != Production {
		r.Use(middleware.Logger)
	}

	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Content-Type"},
			ExposedHeaders: []string{"Link"},
		},
	))

	if env.GetEnvVar().Server.Env == Production {
		r.Use(httprate.LimitByIP(3, 15*time.Minute))

		r.Use(cors.Handler(
			cors.Options{
				AllowedOrigins: []string{Panth3rHOMEURL},
				AllowedMethods: []string{"GET", "POST", "OPTIONS"},
				AllowedHeaders: []string{"Accept", "Content-Type"},
				ExposedHeaders: []string{"Link"},
			}))
	}

	r.Get("/", a.healthcheckHandler)
	r.Post("/sendmail", a.sendEmail)

	if env.GetEnvVar().Server.Env != Production {
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://127.0.0.1:"+env.GetEnvVar().Server.Port+"/swagger/doc.json"),
		))
	}

	return r
}
