package control_structures

const add string = "+"
const subtract string = "-"
const multiply string = "*"
const divide = "/"

func calculate(op string, arg1, arg2 float32) float32 {
	switch op{
		case add: return arg1+arg2
		case subtract:return arg1-arg2
		case multiply:return arg1*arg2
		case divide:return arg1/arg2
	}
	return 0.0
}
