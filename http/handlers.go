package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"panth3rWaitlistBackend/internal/env"
)

// sendEmail
//
//	@Summary		sends user registration email
//	@Description	sends user registration email
//	@Tags			application
//	@QueryParam		name        description="The name of the user" type="string"//	@Produce											plain
//	@Param			name															query		string	true	"Users any preferred name"
//	@Param			email															query		string	true	"valid email address"	Format(email)
//	@Success		200																{string}	string	"OK"
//	@Failure		400																{string}	string	"CLIENT ERROR: BAD REQUEST"
//	@Failure		500																{string}	string	"SERVER ERROR: INTERNAL SERVRER ERROR"
//	@Router			/joinwaitlist [post]
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
//	@Router			/healthcheck [get]
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

// WaitListHandler
//
//	@Summary		sends user registration email
//	@Description	sends user registration email
//	@Tags			application
//	@Accept			x-www-form-urlencoded
//	@Produce		plain
//	@Param			name	formData	string	true	"Users any preferred name"
//	@Param			email	formData	string	true	"valid email address"	Format(email)
//	@Success		200		{string}	string	"OK"
//	@Failure		400		{string}	string	"CLIENT ERROR: BAD REQUEST, INVALID USER FORM DATA"
//	@Failure		409		{string}	string	"CLIENT ERROR: USER WITH EMAIL ALREADY EXIST"
//	@Failure		500		{string}	string	"SERVER ERROR: INTERNAL SERVRER ERROR"
//	@Router			/joinwaitlist [post]
// func (a Application) waitListHandler(w http.ResponseWriter, r *http.Request) {

// 	if err := r.ParseForm(); err != nil {
// 		a.clientError(w, http.StatusBadRequest, err)
// 		a.logger.LogAttrs(context.TODO(), slog.LevelError, err.Error())
// 		return
// 	}

// 	var (
// 		subscriber store.User
// 		err        error
// 	)

// 	if name := r.PostForm.Get("name"); name != "" {
// 		subscriber.Name = name
// 	}
// 	if email := r.PostForm.Get("email"); email != "" {
// 		subscriber.Email = email
// 	}

// 	if (subscriber == store.User{}) {
// 		a.clientError(w, http.StatusBadRequest, fmt.Errorf("empty form data"))
// 		a.logger.LogAttrs(context.TODO(), slog.LevelError, fmt.Errorf("empty form data").Error())
// 		return
// 	}

// 	if a.repository.CheckIfUserExist(subscriber) {
// 		a.clientError(w, http.StatusConflict, fmt.Errorf(http.StatusText(http.StatusConflict)+": USER WITH EMAIL ALREADY EXIST"))
// 		a.logger.LogAttrs(context.TODO(), slog.LevelWarn, fmt.Errorf("USER WITH EMAIL ALREADY EXIST").Error())
// 		return
// 	}

// 	if err = sendConfirmationMail(subscriber.Name, subscriber.Email); err != nil {
// 		a.serverError(w)
// 		a.logger.LogAttrs(context.TODO(), slog.LevelError, err.Error())
// 		return
// 	}

// 	if err := a.repository.SaveToDb(subscriber); err != nil {
// 		if strings.Contains(err.Error(), "duplicate key error") {
// 			a.clientError(w, http.StatusBadRequest, fmt.Errorf("USER WITH EMAIL ALREADY EXIST"))
// 			a.logger.LogAttrs(context.TODO(), slog.LevelError, fmt.Errorf("USER WITH EMAIL ALREADY EXIST: %w", err).Error())
// 			return
// 		}

// 		a.serverError(w)
// 		a.logger.LogAttrs(context.TODO(), slog.LevelError, err.Error())
// 		return
// 	}

// 	w.Write([]byte("EMAIL SENT SUCCESSFULLY"))
// }
