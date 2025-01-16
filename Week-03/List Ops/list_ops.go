package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if s == nil {
		return initial
	}
	for _, val := range s {
		initial = fn(val, initial)
		if initial == 0 {
			return 0
		}

	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	if s == nil {
		return initial
	}
	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
		if initial == 0 {
			return 0
		}
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	list := make(IntList, 0, s.Length())
	index := 0
	for _, val := range s {
		if fn(val) {
			list = list.Append(s[index : index+1])
		}
		index++
	}
	return list
}

func (s IntList) Length() int {
	length := 0
	for _, val := range s {
		_ = val
		length++
	}
	return length
}

func (s IntList) Append(lst IntList) IntList {
	slice := make(IntList, s.Length()+lst.Length())
	index := 0
	for i := 0; i < s.Length(); i++ {
		slice[i] = s[i]
	}
	for i := s.Length(); i < slice.Length(); i++ {
		slice[i] = lst[index]
		index++
	}
	return slice
}

func (s IntList) Map(fn func(int) int) IntList {
	if s == nil {
		return IntList{}
	}
	for i := 0; i < s.Length(); i++ {
		s[i] = fn(s[i])
	}
	return s
}

func (s IntList) Reverse() IntList {
	slice := make(IntList, s.Length())
	index := 0
	for i := s.Length() - 1; i >= 0; i-- {
		slice[index] = s[i]
		index++
	}
	return slice
}

func (s IntList) Concat(lists []IntList) IntList {
	if lists == nil {
		return s
	}
	length := s.Length() + len(lists)
	for _, list := range lists {
		length += list.Length()
	}

	result := s
	for _, list := range lists {
		result = result.Append(list)
	}
	return result
}
