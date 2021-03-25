package sessions

import "github.com/gorilla/sessions"

// the byte array passed here is being used as a key to sign our cookies
// "secret-key" is a random string here
var Store = sessions.NewCookieStore([]byte("secret-key"))

