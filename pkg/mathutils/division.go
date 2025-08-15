package mathutils
import "errors"

func IntegerDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err 
	}
	quotient := numerator / denominator
	remainder := numerator % denominator
	return quotient, remainder, err
}
