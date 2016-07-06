package main

import (
	"fmt"
)

func main() {

	testRun := runTests()

	if testRun {
		fmt.Println("You got it!")
	} else {
		fmt.Println("Try again.")
	}

}

// 1. Using only AND and NOT, build OR
func or(one bool, two bool) bool {
	return !(!one && !two)
}

// 2. Using only AND and NOT, and whatever you can build out of those, build EQUALITY
func equal(one bool, two bool) bool {

	orValue := or(one, two)
	andValue := (one && two)
	equal := or(!orValue, andValue)

	if equal {
		return true
	}
	return false

}

// 3. Using only AND, NOT, EQUAL, and whatever you can build out of those, build XOR
func xor(one bool, two bool) bool {

	xor := !equal(one, two)

	if xor {
		return true
	}
	return false

}

// 4. Using only AND, NOT, EQUAL, XOR, and whatever you can build out of those, build IMPLICATION
func implication(one bool, two bool) bool {

	implication := or(!one, two)

	if implication {
		return true
	}
	return false

}

// (precursor/helper to) 5. Implement NAND
func nand(one bool, two bool) (bool, error) {

	if one == true && two == true {
		return false, nil
	} else if one == true && two == false {
		return true, nil
	} else if one == false && two == true {
		return true, nil
	} else if one == false && two == false {
		return true, nil
	}

	return false, fmt.Errorf("NAND is not implemented correctly")

}

// 5. Using only NAND, make AND
func and(one bool, two bool) bool {

	// The idea is to just run it through twice.
	original, err := nand(one, two)
	clone, err := nand(one, two)
	if err != nil {
		fmt.Println("error: ", err)
		return false
	}

	and, err := nand(original, clone)
	if err != nil {
		fmt.Println("error: ", err)
		return false
	}

	if and {
		return true
	}
	return false

}

// 6. Using only NAND, make NOT
func not(one bool) bool {

	not, err := nand(one, one)
	if err != nil {
		fmt.Println("error: ", err)
		return false
	}

	if not {
		return true
	}
	return false

}
