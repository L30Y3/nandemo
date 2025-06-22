package oauth

import (
	"fmt"
	"net/http"
)

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TODO: Implement Google OAuth2 Login")
}

func MicrosoftLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TODO: Implement Microsoft OAuth2 Login")
}
