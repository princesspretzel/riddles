package main

import (
	"fmt"
)

func runTests() bool {

	passOr := orTest()
	passEqual := equalTest()
	passXor := xorTest()
	passImplication := implicationTest()
	passAnd := andTest()
	passNot := notTest()

	if passOr && passEqual && passXor && passImplication && passAnd && passNot {
		return true
	}
	return false

}

// 1. Using only AND and NOT, build OR
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
func notTest() bool {

	true := not(false)
	false := not(true)

	if true == true && false == false {
		return true
	}
	return false

}
