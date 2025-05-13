package stack

type Stack struct {
	head int
	data [10]int
}

func New() Stack {
	return Stack{head: 0}
}

// Em Go é definido o receptor da função antes do nome da função
// como por exemplo nesta função, o receptor seria (s *stack)
// Exemplo real: <nome da stack>.push(valor inteiro)
func (s *Stack) Push(valor int) {
	if s.head >= len(s.data) {
		return
	}
	s.data[s.head] = valor
	s.head++
}

func (s *Stack) Pop() (valorRetirado int) {
	if s.head == 0 {
		return 0
	}
	s.head-- // O head aponta para o próximo espaço vazio na fila, por isso s.head-- primeiro
	valorRetirado = s.data[s.head]
	s.data[s.head] = 0
	return valorRetirado
}
