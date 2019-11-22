package auth

import (
	"net/http"
	"os"

	"golang-test/goFileBrowser/errors"
	"golang-test/goFileBrowser/settings"
	"golang-test/goFileBrowser/users"
)

// MethodProxyAuth is used to identify no auth.
const MethodProxyAuth settings.AuthMethod = "proxy"

// ProxyAuth is a proxy implementation of an auther.
type ProxyAuth struct {
	Header string `json:"header"`
}

// Auth authenticates the user via an HTTP header.
func (a ProxyAuth) Auth(r *http.Request, sto *users.Storage, root string) (*users.User, error) {
	username := r.Header.Get(a.Header)
	user, err := sto.Get(root, username)
	if err == errors.ErrNotExist {
		return nil, os.ErrPermission
	}

	return user, err
}

// LoginPage tells that proxy auth doesn't require a login page.
func (a ProxyAuth) LoginPage() bool {
	return false
}
