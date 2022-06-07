package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	errorEmptyInput          = errors.New("input is empty")
	errorNotTwoOperands      = errors.New("expecting two operands, but received more or less")
	errorCannotParseArgument = errors.New("expecting an operand to be a number")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

var SPACE_REGEXP = regexp.MustCompile(`\s+`)
var EXPRESSION_REGEXP = regexp.MustCompile(`^([+-]?[\d\w]{1,})([+-][\d\w]{1,})$`)

func ThrowError(message string, err error) (resultErr error) {
	return fmt.Errorf(message+": %w", err)
}

func StringSum(input string) (output string, err error) {
	str := SPACE_REGEXP.ReplaceAllString(input, "")

	if str == "" {
		return "", ThrowError("Empty", errorEmptyInput)
	}

	submatches := EXPRESSION_REGEXP.FindStringSubmatch(str)

	if len(submatches) != 3 {
		return "", ThrowError("Operands Amount", errorNotTwoOperands)
	}

	a, err := strconv.Atoi(submatches[1])
	if err != nil {
		return "", ThrowError("Should Be Number (first)", errorCannotParseArgument)
	}

	b, err := strconv.Atoi(submatches[2])
	if err != nil {
		return "", ThrowError("Should be Number (second)", errorCannotParseArgument)
	}

	sum := a + b

	return strconv.Itoa(sum), err
}
