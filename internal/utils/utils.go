package utils

import (
	"net/http"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/vitorduarte/phonebook/internal/contact"
)

func CompareTwoContactArrays(a1, a2 []contact.Contact) bool {
	if len(a1) == 0 && len(a2) == 0 {
		return true
	}

	less := func(a, b contact.Contact) bool { return a.Id < b.Id }
	isEqual := cmp.Diff(a1, a2, cmpopts.SortSlices(less)) == ""
	return isEqual
}

// GetIdFromURL extracts the id from the url based on prefix
func GetIdFromPath(r *http.Request, prefix string) (id string) {
	if strings.Contains(r.URL.Path, prefix) {
		id = strings.TrimPrefix(r.URL.Path, prefix)
	}
	return
}
