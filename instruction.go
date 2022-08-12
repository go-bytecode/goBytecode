package main

import (
	"encoding/binary"
	"github.com/holiman/uint256"
)

func opAdd(pc *uint64, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Add(&x, y)
	return nil, nil
}

func opPush1(pc *uint64, scope *ScopeContext) ([]byte, error) {
	var (
		codeLen = uint64(len(scope.code))
		integer = new(uint256.Int)
	)
	*pc += 1
	if *pc < codeLen {
		scope.Stack.push(integer.SetUint64(uint64(scope.code[*pc])))
	} else {
		scope.Stack.push(integer.Clear())
	}
	return nil, nil
}

func opReturn(pc *uint64, scope *ScopeContext) ([]byte, error) {
    ans:=scope.Stack.pop()
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, ans.Uint64())
	return b, nil
}
