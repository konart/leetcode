package problems


import (
	"fmt"
	"errors"
)

func comparePar(x, y byte) bool {
	fmt.Println(string(x), string(y))
	switch x {
	case 123:
		return y == 125
	case 91:
		return y == 93
	case 40:
		return y == 41
	default:
		return false
	}
}

type stack struct {
	pars []byte
}

func (s *stack) add(par byte) {
	s.pars = append(s.pars, par)
}

func (s stack) peek() (error, byte) {
	var err error
	if s.empty() {
		err = errors.New("no more elements")
		return err, 0
	}
	return err, s.pars[len(s.pars)-1]
}

func (s *stack) pop() {
	s.pars = s.pars[0:len(s.pars)-1]
}
func (s stack) empty() bool {
	return len(s.pars) == 0
}

func main() {
	string1 := "([{{()}}])]"
	string2 := "{{(){[[]]}}}"
	fmt.Println("res1")
	result := checkString(string1)
	fmt.Println("res2")
	result1 := checkString(string2)
	fmt.Println(result)
	fmt.Println(result1)
}

func checkString(s string) bool {
	st := &stack{[]byte{}}
	arr := []byte(s)
	for _, par := range arr {
		if par == 123 || par == 91 || par == 40 {
			st.add(par)
		} else {
			err, counterPar := st.peek()
			if err != nil {
				return false
			}
			if comparePar(counterPar, par) {
				st.pop()
			} else {
				return false
			}
		}
	}
	if st.empty() {
		return true
	} else {
		return false
	}
}