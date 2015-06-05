package keys

import (
	"appengine"
	"appengine/datastore"
)

func GuestbookKey(context appengine.Context) *datastore.Key {
	return datastore.NewKey(context, "Guestbook", "default_guestbook", 0, nil)
}