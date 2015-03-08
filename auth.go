package api

type Credentials struct {
	ApiKey   string
	Username string
	Password string
}

func ApiKey(key string) *Credentials {
	creds := &Credentials{ApiKey: key}
	return creds
}

func UserPass(user string, password string) *Credentials {
	creds := &Credentials{Username: user, Password: password}
	return creds
}
