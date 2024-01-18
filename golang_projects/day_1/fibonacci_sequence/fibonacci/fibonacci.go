package fibonacci

// Fibonacci Sequence
// 1, 1, 2, 3, 5, 8, 13, 21, ...
// A(n) = A(n-1) + A(n-2) with A(0) = 1 A(1) = 1

type Sequence []int

func (s *Sequence) FibonacciElement(position int) int {
	if position > len(*s) {
		s.FibonacciSequenceIncrement(position)
	}
	return (*s)[position-1]
}

func (s *Sequence) FibonacciPartialSequence(initialPosition int, finalPosition int) []int {
	return (*s)[initialPosition:finalPosition]
}

func (s *Sequence) FibonacciSequenceIncrement(position int) {
	for i := len(*s) - 1; i < position-1; i++ {
		nextSequenceElement := (*s)[i] + (*s)[i-1]
		*s = append(*s, nextSequenceElement)
	}
}

func (s *Sequence) FibonacciSequenceInitializer() {
	*s = append(*s, 1, 1)
}

func SequenceInitializer() Sequence {
	s := make(Sequence, 0)
	s.FibonacciSequenceInitializer()
	return s
}
