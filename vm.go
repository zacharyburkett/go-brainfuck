package brainfuck

// ============================================================
// Types and globals
// ============================================================

// VM represents the machine state
// of a virtual brainfuck execution machine.
// This state includes program memory,
// tape memory, and instruction pointer.
type VM struct {
	prog    []byte
	instr   int
	p       int
	tape    [2048]byte
	out, in chan byte
}

var ops = map[byte]func(*VM){
	'>': (*VM).pr,
	'<': (*VM).pl,
	'+': (*VM).add,
	'-': (*VM).sub,
	',': (*VM).i,
	'.': (*VM).o,
	'[': (*VM).jmp0,
	']': (*VM).jmpn0,
}

// ============================================================
// Constructor
// ============================================================

// NewVM instantiates a new instance
// of the brainfuck virtual machine with a loaded program.
func NewVM(prog []byte) *VM {
	return &VM{prog: prog, out: make(chan byte), in: make(chan byte)}
}

// ============================================================
// Methods
// ============================================================

// Exec runs the virtual machine.
// And error may be returned on syntax validation.
func (vm *VM) Exec() error {
	if err := ValidateSyntax(vm.prog); err != nil {
		return err
	}
	for {
		if vm.instr >= len(vm.prog) {
			break
		}
		op := ops[vm.prog[vm.instr]]
		op(vm)
		vm.instr++
	}
	return nil
}

// LoadProg sets the VM's program memory
// with a given program.
func (vm *VM) LoadProg(prog []byte) {
	vm.prog = prog
	vm.out = make(chan byte)
	vm.in = make(chan byte)
}

func (vm *VM) Read() byte {
	return <-vm.out
}

func (vm *VM) Write(b byte) {
	vm.in <- b
}

// ============================================================
// Operation Methods
// ============================================================

func (vm *VM) pr() {
	vm.p++
	if vm.p >= len(vm.tape) {
		vm.p = 0
	}
}

func (vm *VM) pl() {
	vm.p--
	if vm.p < 0 {
		vm.p = len(vm.tape) - 1
	}
}

func (vm *VM) add() {
	vm.tape[vm.p]++
}

func (vm *VM) sub() {
	vm.tape[vm.p]--
}

func (vm *VM) o() {
	vm.out <- vm.tape[vm.p]
}

func (vm *VM) i() {
	vm.tape[vm.p] = <-vm.in
}

func (vm *VM) jmp0() {
	if vm.tape[vm.p] != 0 {
		return
	}
	var depth int
	vm.instr++
	for depth > 0 || vm.prog[vm.instr] != ']' {
		switch vm.prog[vm.instr] {
		case '[':
			depth++
		case ']':
			depth--
		}
		vm.instr++
	}
}

func (vm *VM) jmpn0() {
	var depth int
	vm.instr--
	for depth > 0 || vm.prog[vm.instr] != '[' {
		switch vm.prog[vm.instr] {
		case '[':
			depth--
		case ']':
			depth++
		}
		vm.instr--
	}
	vm.instr--
}
