package set

var exists = struct{}{}

type Set struct {
	hmap map[string]struct{}
}

func New() *Set {
	return &Set{
		hmap: make(map[string]struct{}),
	}
}

func (s Set) isEmpty() bool {
	return s.Size() == 0
}

func (s Set) Size() int {
	return len(s.hmap)
}

func (s *Set) Add(value string) error {
	s.hmap[value] = exists
	return nil
}

func (s *Set) Remove(value string) (string, error) {
	delete(s.hmap, value)
	return value, nil
}

func (s *Set) Has(value string) (bool, error) {
	_, c := s.hmap[value]
	return c, nil
}
