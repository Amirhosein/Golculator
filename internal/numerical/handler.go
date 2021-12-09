package numerical

import "strconv"

func Calculate(num1, num2 float64, op string) float64 {
	switch op {
	case "+":
		return Add(num1, num2)
	case "-":
		return Sub(num1, num2)
	case "*":
		return Mul(num1, num2)
	case "/":
		return Div(num1, num2)
	}
	return 0
}

func ParseNumber(message string) (float64, error) {
	num, err := strconv.ParseFloat(message, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func IsOperator(msg string) bool {
	return msg == "+" || msg == "-" || msg == "*" || msg == "/"
}
