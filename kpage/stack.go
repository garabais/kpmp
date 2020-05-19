package kpage

type stack []uint

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(v uint) {
	*s = append(*s, v) // Simply append the new value to the end of the stack
}

func (s *stack) pop() (uint, bool) {
	if s.isEmpty() {
		return 0, false
	}

	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element, true

}
