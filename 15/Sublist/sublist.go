package sublist

import "slices"

// Relation type is defined in relations.go file.

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

func Sublist(l1, l2 []int) Relation {
	if slices.Compare(l1, l2) == 0 {
		return RelationEqual
	} else if len(l1) == 0 {
		return RelationSublist
	} else if len(l2) == 0 {
		return RelationSuperlist
	} else if len(l1) > len(l2) {
		status := false
		for i := 0; i < len(l1)-len(l2)+1; i++ {
			slice := l1[i : i+len(l2)]
			for j := 0; j < len(slice); j++ {
				if slice[j] == l2[j] {
					status = true
				} else {
					status = false
					break
				}
			}
			if status {
				return RelationSuperlist
			}
		}
	} else if len(l2) > len(l1) {
		status := false
		for i := 0; i < len(l2)-len(l1)+1; i++ {
			slice := l2[i : i+len(l1)]
			for j := 0; j < len(slice); j++ {
				if slice[j] == l1[j] {
					status = true
				} else {
					status = false
					break
				}
			}
			if status {
				return RelationSublist
			}
		}
	}
	return RelationUnequal
}
