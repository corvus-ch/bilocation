package matcher

// Stack implement a basic stack
type Stack struct {
	data []Matcher
	sp   int
}

func NewStack() *Stack {
	return &Stack{sp: -1}
}

// Pop pops the element atop of the stack
func (s *Stack) Pop() Matcher {
	if s.sp < 0 {
		return nil
	}
	r := s.data[s.sp]
	s.data = s.data[:s.sp]
	s.sp--

	return r
}

// Push pushes an element into the top of the stack
func (s *Stack) Push(b Matcher) {
	s.sp++
	s.data = append(s.data, b)
}

// Drain pops all elements from the stack.
func (s *Stack) Drain() []Matcher {
	if s.sp < 1 {
		return make([]Matcher, 0)
	}
	r := make([]Matcher, 0, s.sp)
	for _, m := range s.data[:s.sp+1] {
		r = append(r, m)
	}
	s.sp = -1
	s.data = s.data[:0]

	return r
}

func (s *Stack) Depth() int {
	return s.sp + 1
}
