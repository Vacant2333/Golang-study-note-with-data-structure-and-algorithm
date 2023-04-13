package set_map

type Set map[int]struct{}

func New() Set {
	return make(map[int]struct{})
}

func (s Set) Add(element int) {
	s[element] = struct{}{}
}

func (s Set) Remove(element int) {
	delete(s, element)
}

func (s Set) Contains(element int) bool {
	_, ok := s[element]
	return ok
}

func (s Set) Union(s2 Set) {
	for k := range s2 {
		s.Add(k)
	}
}
