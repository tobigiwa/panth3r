![Logo](/templ/Panth3r.jpg)

# Waitlist Backend


#### _version: 1.0.0_ - official release

## START SERVER

### DEV MODE:

> $ make dev

Defaults .env variables are enabled.

### PROD MODE:

> hosting services matters

> $ make prod

This strips the binary, suitable for deployment.

## CODE:

### Routes:

**POST:** _"/sendmail"_ **sends user registration email**

**GET:** _"/healthcheck"_ **Report application status**

**GET:** _"/swagger/\*"_ **OpenAPI Schema documentation**

### Frontend endpoint integration:

#### **POST:** _"/sendmail"_

> This endpoint takes a url-encoded form data, see **GET:** _"/swagger/\* for complete docs"_.

- 200 StatusOK - plain text (**EMAIL SENT SUCCESSFULLY**) is sent alongside header.
- 400 StatusBadRequest, invalid query paraameters.
- 500 StatusInternalServerError - server error, header only.
