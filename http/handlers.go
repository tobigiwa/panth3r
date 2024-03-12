package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"panth3rWaitlistBackend/internal/env"
)

// SendEmail
//
//	@Summary		sends user registration email
//	@Description	sends user registration email
//	@Tags			application
//	@QueryParam		name        description="The name of the user" type="string"
//	@QueryParam		email        description="The email of the user" type="string"
//
//	@Produce		plain
//
//	@Param			name	query		string	true	"Users any preferred name"
//	@Param			email	query		string	true	"valid email address"	Format(email)
//	@Success		200		{string}	string	"OK"
//	@Failure		400		{string}	string	"CLIENT ERROR: BAD REQUEST"
//	@Failure		500		{string}	string	"SERVER ERROR: INTERNAL SERVRER ERROR"
//	@Router			/sendmail [post]
func (a Application) sendEmail(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	name := params.Get("name")
	email := params.Get("email")

	if name == "" || email == "" {
		a.clientError(w, http.StatusBadRequest, fmt.Errorf("%v:name or email was empty", http.StatusText(http.StatusBadRequest)))
		a.logger.LogAttrs(context.TODO(), slog.LevelError, fmt.Sprintf("%v:name or email was empty", http.StatusText(http.StatusBadRequest)))
		return
	}

	if err := sendConfirmationMail(name, email); err != nil {
		a.serverError(w)
		a.logger.LogAttrs(context.TODO(), slog.LevelError, err.Error())
		return
	}

	w.Write([]byte("OK"))
}

type ServerStatus struct {
	Server_status       string
	Application_Env     string
	Application_Version string
}

// HealthcheckHandler
//
//	@Summary		Report application status
//	@Description	return application status
//	@Tags			status
//	@Produce		json
//	@Success		200	{object}	http.ServerStatus	"Server_status:available"
//	@Failure		500	{string}	string				"INTERNAL SERVRER ERROR"
//	@Router			/ [get]
func (a Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	env := ServerStatus{
		Server_status:       "available",
		Application_Env:     env.GetEnvVar().Server.Env,
		Application_Version: env.GetEnvVar().Server.Version + " 1",
	}
	var (
		byteArr []byte
		err     error
	)
	if byteArr, err = json.Marshal(env); err != nil {
		a.logger.LogAttrs(context.TODO(), slog.LevelError, "marshling error from healthcheckHandler"+err.Error())
		a.serverError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteArr)
}
