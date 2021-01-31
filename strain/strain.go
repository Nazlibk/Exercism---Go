package strain

type Ints []int
type Strings []string
type Lists []Ints

func (list Ints) Keep(f func(n int) bool) Ints {
	var result = make(Ints, 0, len(list))
	if list == nil {
		return nil
	}
	for _, n := range list {
		if f(n) {
			result = append(result, n)
		}
	}
	return result
}

func (list Ints) Discard(f func(n int) bool) Ints {
	var result = make(Ints, 0, len(list))
	if list == nil {
		return nil
	}
	for _, n := range list {
		if !f(n) {
			result = append(result, n)
		}
	}
	return result
}

func (list Strings) Keep(f func(s string) bool) Strings {
	var result = make(Strings, 0, len(list))
	if list == nil {
		return nil
	}
	for _, s := range list {
		if f(s) {
			result = append(result, s)
		}
	}
	return result
}

func (list Strings) Discard(f func(s string) bool) Strings {
	var result = make(Strings, 0, len(list))
	if list == nil {
		return nil
	}
	for _, s := range list {
		if !f(s) {
			result = append(result, s)
		}
	}
	return result
}

func (list Lists) Keep(f func(l []int) bool) Lists {
	var result = make(Lists, 0, len(list))
	if list == nil {
		return nil
	}
	for _, l := range list {
		if f(l) {
			result = append(result, l)
		}
	}
	return result
}

func (list Lists) Discard(f func(l []int) bool) Lists {
	var result = make(Lists, 0, len(list))
	if list == nil {
		return nil
	}
	for _, l := range list {
		if !f(l) {
			result = append(result, l)
		}
	}
	return result
}
