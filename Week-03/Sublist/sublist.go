// Package sublist provides functionality for comparing two lists
// and determining their relationship, such as whether one is a sublist,
// superlist, or equal to the other.
package sublist

import "slices"

// IsSublist determines whether `smallerList` is a sublist of `greaterList`.
// It returns true if all elements of `smallerList` appear consecutively
// and in the same order within `greaterList`.
//
// Parameters:
//   - greaterList: The larger list in which to search for the smaller list.
//   - smallerList: The list to check as a sublist of `greaterList`.
//
// Returns:
//   - true if `smallerList` is a sublist of `greaterList`; false otherwise.
func IsSublist(greaterList, smallerList []int) bool {
	status := false
	for i := 0; i < len(greaterList)-len(smallerList)+1; i++ {
		sublist := greaterList[i : i+len(smallerList)]
		for j := 0; j < len(sublist); j++ {
			if sublist[j] == smallerList[j] {
				status = true
			} else {
				status = false
				break
			}
		}
		if status {
			return status
		}
	}
	return false
}

// Sublist compares two lists, `l1` and `l2`, and determines their relationship.
// It returns one of the following relations:
//   - RelationEqual: The lists are identical.
//   - RelationSublist: `l1` is a sublist of `l2`.
//   - RelationSuperlist: `l1` is a superlist of `l2`.
//   - RelationUnequal: The lists are neither sublists nor superlists of each other.
//
// Parameters:
//   - l1: The first list to compare.
//   - l2: The second list to compare.
//
// Returns:
//   - A Relation indicating the relationship between the two lists.
func Sublist(l1, l2 []int) Relation {
	if slices.Compare(l1, l2) == 0 {
		return RelationEqual
	} else if len(l1) == 0 {
		return RelationSublist
	} else if len(l2) == 0 {
		return RelationSuperlist
	} else if len(l1) > len(l2) && IsSublist(l1, l2) {
		return RelationSuperlist
	} else if len(l2) > len(l1) && IsSublist(l2, l1) {
		return RelationSublist
	}
	return RelationUnequal
}
