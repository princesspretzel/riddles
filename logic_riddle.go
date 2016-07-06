package main

import (
	"fmt"
)

func main() {

	passOr := orTest()
	passEqual := equalTest()
	passXor := xorTest()
	passImplication := implicationTest()
	passAnd := andTest()
	passNot := notTest()

	if passOr && passEqual && passXor && passImplication && passAnd && passNot {
		fmt.Println("Yay you win!")
	} else {
		fmt.Println("Dun goofed...")
	}

}

// 1. Using only AND and NOT, build OR
func or(one bool, two bool) bool {
	return !(!one && !two)
}

func orTest() bool {

	truetrue := or(true, true)
	truefalse := or(true, false)
	falsetrue := or(false, true)
	falsefalse := or(false, false)

	if truetrue == true && truefalse == true && falsetrue == true && falsefalse == false {
		return true
	}
	return false

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

func equalTest() bool {

	truetrue := equal(true, true)
	truefalse := equal(true, false)
	falsetrue := equal(false, true)
	falsefalse := equal(false, false)

	if truetrue == true && truefalse == false && falsetrue == false && falsefalse == true {
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

func xorTest() bool {

	truetrue := xor(true, true)
	truefalse := xor(true, false)
	falsetrue := xor(false, true)
	falsefalse := xor(false, false)

	if truetrue == false && truefalse == true && falsetrue == true && falsefalse == false {
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

func implicationTest() bool {

	truetrue := implication(true, true)
	truefalse := implication(true, false)
	falsetrue := implication(false, true)
	falsefalse := implication(false, false)

	if truetrue == true && truefalse == false && falsetrue == true && falsefalse == true {
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

func nandTest() bool {

	truetrue, err := nand(true, true)
	truefalse, err := nand(true, false)
	falsetrue, err := nand(false, true)
	falsefalse, err := nand(false, false)
	if err != nil {
		fmt.Println("error: ", err)
		return false
	}

	if truetrue == false && truefalse == true && falsetrue == true && falsefalse == true {
		return true
	}
	return false

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

func andTest() bool {

	truetrue := and(true, true)
	truefalse := and(true, false)
	falsetrue := and(false, true)
	falsefalse := and(false, false)

	if truetrue == true && truefalse == false && falsetrue == false && falsefalse == false {
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

func notTest() bool {

	true := not(false)
	false := not(true)

	if true == true && false == false {
		return true
	}
	return false

}

// // 7. Using only AND and OR, make NOT (might be impossible)
// func notImpossible(one bool, two bool) bool {

// }

// func notImpossibleTest() bool {

// }
