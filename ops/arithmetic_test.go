package ops

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/vmgen"
)

func TestAddPositive(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	executeAddition(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(intAsBig(o1+o2)) == 0, "wrong add value")
}

func TestAddNegative(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := -4, -2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	executeAddition(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(intAsBig(o1+o2)) == 0, "wrong add value")
}

func TestAddZero(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := -4, 0
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	executeAddition(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(intAsBig(o1+o2)) == 0, "wrong add value")
}

func TestSubPositive(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	executeSubtraction(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(intAsBig(o1-o2)) == 0, "wrong sub value")
}

func TestSubNegative(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := -4, -2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	executeSubtraction(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(intAsBig(o1-o2)) == 0, "wrong sub value")
}

func TestSubZero(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := -4, 0
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	executeSubtraction(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(intAsBig(o1-o2)) == 0, "wrong sub value")
}

func TestMul(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	executeMultiplication(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(intAsBig(o1*o2)) == 0, "wrong mul value")
}

func TestDiv(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	executeDivision(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(intAsBig(o1/o2)) == 0, fmt.Sprintf("wrong div value: %d", c.Int64()))
}

func TestMod(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	executeModulo(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(intAsBig(o1%o2)) == 0, fmt.Sprintf("wrong mod value: %d", c.Int64()))
}

func TestAddMod(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2, o3 := 4, 2, 1
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	s.Push(intAsBig(o3).Bytes())
	executeAddMod(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	//c := new(big.Int).SetBytes(s.Pop())

}

func TestMulMod(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2, o3 := 4, 2, 1
	s.Push(intAsBig(o1).Bytes())
	s.Push(intAsBig(o2).Bytes())
	s.Push(intAsBig(o3).Bytes())
	executeMulMod(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	//c := new(big.Int).SetBytes(s.Pop())
}

func TestConcatBothFull(t *testing.T) {
	s := new(vmgen.Stack)
	s1 := "hello"
	s2 := "world"
	s.Push([]byte(s1))
	s.Push([]byte(s2))
	executeConcat(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	p := s1 + s2
	goutil.Assert(t, len(s.Pop()) == len([]byte(p)), "wrong value")
}

func TestConcatBothEmpty(t *testing.T) {
	s := new(vmgen.Stack)
	s1 := ""
	s2 := ""
	s.Push([]byte(s1))
	s.Push([]byte(s2))
	executeConcat(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	p := s1 + s2
	goutil.Assert(t, len(s.Pop()) == len([]byte(p)), "wrong value")
}

func TestConcatOneEmpty(t *testing.T) {
	s := new(vmgen.Stack)
	s1 := "hello"
	s2 := ""
	s.Push([]byte(s1))
	s.Push([]byte(s2))
	executeConcat(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	p := s1 + s2
	goutil.Assert(t, len(s.Pop()) == len([]byte(p)), "wrong value")
}
