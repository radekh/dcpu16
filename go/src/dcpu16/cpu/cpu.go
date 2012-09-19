/*
 * DCPU16 emulator.
 * Main file.
 * Copyright (c) 2012 Radek Hnilica
 * Licence: GPLv3
 */
package cpu;

//import "fmt"

type Word   uint16
type SWord  int16

type Cpu struct {
	// Eight general purpose registers
	regA, regB, regC    Word
	regX, regY, regZ    Word
	regI, regJ          Word

	regSP   Word		// Stack pointer
	regPC   Word	        // program counter / instruction pointer
	regEX   Word		// overflow register
	regIA   Word		// Interrupt Address

	skip    bool
	memory  [0x10000]Word	// Computer memory
	cycle   int64		// The biggest possible integer type
}

// Factory method creating new instance of Cpu
func New() *Cpu {
	c := new(Cpu)
	return c
}


/*
 * Definition of object methods.
 */

/*
 * Reset CPU does basic CPU settings.  It initializes all internal
 * circuits, flip-flops, and registers.
 */
func (c *Cpu) Reset() {
	c.skip = false
	// Because all values are zero by GoLang definitions, we do not set.
}


/*
 * Step method process one instruction and updates according to it the
 * state of the Cpu and memory.
 */
func (c *Cpu) Step() {
	// *** FETCH ***
	instruction := c.memory[c.regPC]
	c.regPC++	
	// c.cycle is not updated here

	// *** DECODE ***
	opcode := instruction & 0x1f
	var (
		source Word
		destination *Word
	)
	/*
	 * There are two basic instruction forms.  Two operand
	 * instruction, where opcode >0, and single operand (special)
	 * instruction where opcode ==0.
	 */
	if opcode != 0 {
		source      = c.decode_source((instruction >> 10) & 0x3f)
		destination = c.decode_destination((instruction >> 5) & 0x1f)
	} else {
		//FIXME: decode operand for special instruction
		source, destination = c.decode_operand((instruction >> 10) & 0x3f)
	}
	
	/*
	 * The instruction has decoded operands and we are prepared to
	 * execute.  This is the right tim to check if c.skip is set.
	 * If it is, we should not execute but skip.  We should also
	 * clear the skip flag.
	 */
	if c.skip {
		c.skip = false
		return
	}

	// *** EXECUTE ***
	switch opcode {
	case 1: // SET: C=1
		c.cycle++
		*destination = source
	case 2: // ADD, update EX
		result := int32(*destination) + int32(source)
		*destination = Word(result)
		c.regEX = Word(result >>16)
	case 3: // SUB, update EX
		result := int32(*destination) - int32(source)
		*destination = Word(result)
		c.regEX = Word(result >>16)
	case 4: // MUL, 
		result := uint32(*destination) * uint32(source)
		*destination = Word(result)
		c.regEX = Word(result >>16)
	case 5: // MLI, 
		result := int32(SWord(*destination)) * int32(SWord(source))
		*destination = Word(result)
		c.regEX = Word(result >>16)
	case 6: // DIV
		if source == 0 {
			// If divisor is zero then by definition result is zero!
			*destination = 0; c.regEX = 0
		} else {
			result := (uint32(*destination) << 16) / uint32(source)
			*destination = Word(result >> 16)
			c.regEX = Word(result)
		}
	case 7: // DVI
		if source == 0 {
			// If divisor is zero then by definition result is zero!
			*destination = 0; c.regEX = 0
		} else {
			result := (int32(SWord(*destination)) << 16) / int32(SWord(source))
			*destination = Word(result >> 16)
			c.regEX = Word(result)
			// negative result correction
			if result < 0 && c.regEX >0 {*destination++}
		}
	case 8: // MOD
		if source == 0 {
			*destination = 0
		} else {
			result := *destination % source
			*destination = result
		}
	case 9: // MDI
		if source == 0 {
			*destination = 0
		} else {
			result := SWord(*destination) % SWord(source)
			*destination = Word(result)
		}
	case 10: // AND
		result := *destination & source
		*destination = result
	case 11: // BOR
		result := *destination | source
		*destination = result
	case 12: // XOR
		result := *destination ^ source
		*destination = result
	case 13: // SHR -- logical shift right
		// Because >> is arithmetic shift the MSB must be 0 to behave like logic shift.
		result := (int64(*destination) << 16) >> source
		*destination = Word(result >> 16)
		c.regEX = Word(result)
	case 14: // ASR -- Arithmetic shift.  Note we do not need int64 as in SHR
		result := (int32(*destination) << 16) >> source
		*destination = Word(result >> 16)
		c.regEX = Word(result)
	case 15: // SHL -- Arithmetic shift left
		result := (int32(*destination)) << source
		*destination = Word(result)
		c.regEX = Word(result >> 16)

	case 16: // IFB
		c.cycle += 2
		c.skip = ! (*destination & source != 0)
	case 17: // IFC
		c.cycle += 2
		c.skip = ! (*destination & source == 0)
	case 18: // IFE
		c.cycle += 2
		c.skip = ! (*destination == source)
	case 19: // IFN
		c.cycle += 2
		c.skip = ! (*destination != source)
	case 20: // IFG
		c.cycle += 2
		c.skip = ! (*destination > source)
	case 21: // IFA
		c.cycle += 2
		c.skip = ! (SWord(*destination) > SWord(source))
	case 22: // IFL
		c.cycle += 2
		c.skip = ! (*destination < source)
	}


	/*
         * The instruction was executed and is finished.  Except if it
         * was IFx instruction.  In this case we must test if skip has
         * to be performed.
         */
	//"Execute" next instruction in skip mode.
	if c.skip {
		c.Step()
		c.cycle++
	}
}


/*
 * decode_source decodes source operand and returns it's value.
 */
func (c *Cpu) decode_source(code Word) (value Word) {
	switch {
	// POP [SP++]
	case code == 0x18: value = c.memory[c.regSP]; c.regSP++;
	case code >= 0x20: // LITERAL
		value = code - 33 // literal value
	default:
		value = *c.decode_destination(code)
	}
	return
}

/*
 * decode_destination decodes destination operand and return pointers
 * to it.  This code is inspired by Scott Fergusons
 * [dcpu16-go](git://github.com/dcpu16/dcpu16-go.git).
 */
func (c *Cpu) decode_destination(code Word) (refference *Word) {
	switch code {
	// Register value: reg
	case 0x00: refference = &c.regA
	case 0x01: refference = &c.regB
	case 0x02: refference = &c.regC
	case 0x03: refference = &c.regX
	case 0x04: refference = &c.regY
	case 0x05: refference = &c.regZ
	case 0x06: refference = &c.regI
	case 0x07: refference = &c.regJ
	// Register refference:  [reg]
	case 0x08: refference = &c.memory[c.regA]
	case 0x09: refference = &c.memory[c.regB]
	case 0x0a: refference = &c.memory[c.regC]
	case 0x0b: refference = &c.memory[c.regX]
	case 0x0c: refference = &c.memory[c.regY]
	case 0x0d: refference = &c.memory[c.regZ]
	case 0x0e: refference = &c.memory[c.regI]
	case 0x0f: refference = &c.memory[c.regJ]
	// Register + literar ref:  [reg+lit]
	case 0x10: refference = &c.memory[c.regA+c.memory[c.regPC]]; c.regPC++
	case 0x11: refference = &c.memory[c.regB+c.memory[c.regPC]]; c.regPC++
	case 0x12: refference = &c.memory[c.regC+c.memory[c.regPC]]; c.regPC++
	case 0x13: refference = &c.memory[c.regX+c.memory[c.regPC]]; c.regPC++
	case 0x14: refference = &c.memory[c.regY+c.memory[c.regPC]]; c.regPC++
	case 0x15: refference = &c.memory[c.regZ+c.memory[c.regPC]]; c.regPC++
	case 0x16: refference = &c.memory[c.regI+c.memory[c.regPC]]; c.regPC++
	case 0x17: refference = &c.memory[c.regJ+c.memory[c.regPC]]; c.regPC++
        // Secial registers and/or operations
	// PUSH: [--SP]
	case 0x18: c.regSP--; refference = &c.memory[c.regSP] // PUSH [--SP]
	case 0x19: refference = &c.memory[c.regSP] // [SP] / PEEK
	case 0x1a: refference = &c.memory[c.regSP+c.memory[c.regPC]]; c.regPC++
	case 0x1b: refference = &c.regSP
	case 0x1c: refference = &c.regPC
	case 0x1d: refference = &c.regEX
	case 0x1e: refference = &c.memory[c.memory[c.regPC]]; c.regPC++
	case 0x1f: refference = &c.memory[c.regPC]; c.regPC++
	}
	
	// *** Update c.cycle, but only if not in skip mode ***
	if ! c.skip {
		if (0x10 <= code && code <= 0x17) || code == 0x1a || code == 0x1e || code == 0x1f {
			c.cycle++
		}
	}
	return
}

/*
 * decode_operand decodes operand in case of special instruction.
 * This operand could be either source or destination.  Because of
 * this we need to return both the value and if it is possible the
 * refference.
 */
func (c *Cpu) decode_operand(code Word) (value Word, refference *Word) {
	if code < 0x20 {	// register or memory
		refference = c.decode_destination(code)
		value = *refference
	} else {		// >=0x20 -- LITERAL
		value = code - 33 // literal value
	}
	return
}
