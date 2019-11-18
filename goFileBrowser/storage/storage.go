package storage

import (
	"golang-test/goFileBrowser/auth"
	"golang-test/goFileBrowser/settings"
	"golang-test/goFileBrowser/share"
	"golang-test/goFileBrowser/users"
)

// Storage is a storage powered by a Backend whih makes the neccessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    *users.Storage
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}
