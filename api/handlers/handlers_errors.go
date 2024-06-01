package handlers

const (
	ErrUnableProcessReq = "sorry! unable to process the request at the moment"
	ErrNoKeyProvidedUrl = "no key was provided in url path"
	ErrInvalidUserJson  = "invalid json format for user schema"
	ErrAuth             = "failed to authenticate the user"
	ErrWrongPassword    = "wrong password provided"
	ErrNoEmail          = "no email in the json payload"
)
