package day11

import (
	"fmt"
	"math/big"
	"testing"
)

func TestAdd(t *testing.T) {

	a := big.NewInt(12312313)
	b := big.NewInt(92589585858)
	c := addOp(*a, *b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
