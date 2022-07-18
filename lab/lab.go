package lab


func BinaryTree(expression string) []Exp {
	var stack[] int
	var start int
	var expressions []Exp
	var tempExp Exp

	for pos, char := range expression{
		// fmt.Println(string(char))

		if string(char) == "("{
			stack = append(stack, pos)
		} else if string(char) == ")"{
			start = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tempExp =  Exp{
				Order:len(stack), 
				Value:expression[start+1: pos],
			}
			expressions = append(expressions, tempExp)
		}

	}
	for i := range expressions {
		fmt.Println(expressions[i].Value, " -- ", expressions[i].Order)
	}
	return expressions
}
