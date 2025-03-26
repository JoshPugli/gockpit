package types

type Set[T comparable] map[T]struct{}

// Add an element to the set. If the element already exists, it will not be added again.
func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

// Remove an element from the set if it exists, otherwise do nothing.
func (s Set[T]) Remove(element T) {
	delete(s, element)
}

func NewSet[T comparable](elements ...T) Set[T] {
	set := make(Set[T])
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

// Contains checks if the set contains the specified element.
func (s Set[T]) Contains(element T) bool {
	_, exists := s[element]
	return exists
}

// IsEmpty checks if the set is empty.
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

// Size returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s)
}

// Elements returns a slice of all elements in the set.
func (s Set[T]) Elements() []T {
	elements := make([]T, 0, len(s))
	for element := range s {
		elements = append(elements, element)
	}
	return elements
}
