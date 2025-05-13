package stack

type Stack struct {
	data []int
}

func New() Stack {
	return Stack{data: []int{}}
}

// Em Go é definido o receptor da função antes do nome da função
// como por exemplo nesta função, o receptor seria (s *stack)
// Exemplo real: <nome da stack>.push(valor inteiro)
func (s *Stack) Push(valor int) {
	s.data = append(s.data, valor)
}

func (s *Stack) Pop() (valorRetirado int) {
	if len(s.data) == 0 {
		return 0
	}
	valor := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return valor
}
