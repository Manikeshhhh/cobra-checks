package checkcobra



//package main
import (
	"fmt"
	"regexp"
	"strings"
)

// func main() {
// 	fmt.Println("Hello, this check passowrd strength")
// 	str := "Lexrta@2145"
// 	hasSpecial := hasSpecialChar(str)

// 	// Check if the string is valid (has a length greater than 8 and contains a capital letter)

func isvalid() bool{
	isValid := isValidString(str)
	if hasSpecial == true && isValid == true {
		return true
	}
}
//}

func hasSpecialChar(str string) bool {
	// Compile the regular expression that matches special characters
	re := regexp.MustCompile("[^a-zA-Z0-9]+")

	// Check if the string contains a match for the regular expression
	return re.MatchString(str)
}

func isValidString(str string) bool {
	// Check if the string has a length greater than 8
	if len(str) <= 8 {
		return false
	}

	// Convert the string to uppercase
	upperStr := strings.ToUpper(str)

	// Check if the string contains a capital letter
	return str != upperStr
}


//add another function that will check if both is true,if its true return true 
