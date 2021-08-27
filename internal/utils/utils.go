package utils

import (
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
