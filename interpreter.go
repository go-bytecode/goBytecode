package main

type OpCode byte

type ScopeContext struct {
	Memory   *Memory
	Stack    *Stack
    code []byte
}



func Run(code []byte) (ret []byte, err error) {
	var (
		JT          = newFrontierInstructionSet()
		op          OpCode        // current opcode
		mem         = NewMemory() // bound memory
		stack       = newstack()  // local stack
		callContext = &ScopeContext{
			Memory: mem,
			Stack:  stack,
			code:   code,
		}
		pc = uint64(0) // program counter
		_  []byte      // result of the opcode execution function
	)
	// Don't move this deferrred function, it's placed before the capturestate-deferred method,
	// so that it get's executed _after_: the capturestate needs the stacks before
	// they are returned to the pools


	for {

		op = OpCode(code[pc])
		operation := JT[op]
		if operation == nil {
			return nil, err
		}
		if sLen := stack.len(); sLen < operation.minStack {
			return nil, err
		}

		ret, err = operation.execute(&pc, callContext)
		// if the operation clears the return data (e.g. it has returning data)
		// set the last return to the result of the operation.

		switch {
		case err != nil:
			return nil, err
		case operation.halts:
			return ret, nil
		case !operation.jumps:
			pc++
		}
	}

}