package switch_statement

import "fmt"

func PrintSwitchStatement() {
	var value = 10

	switch value {
	case 10:
		fmt.Println("the value is correct")
	case 20:
		fmt.Println("the value is not correct")
	default:
		fmt.Println("undefined value")
	}
}
