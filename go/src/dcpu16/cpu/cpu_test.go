/*
 * $Header$
 * Unit tests.
 * Copyright (c) 2012 Radek Hnilica
 * License: GPLv3 or at your opinion higher.
 */
package cpu

import "testing"

func Test_Register_A(t *testing.T) {
	c := New()
	c.regA = 0
	c.regA = 0xffff
	t.Log("A register is OK")
}

func Test_Registers_1(t *testing.T) {
	t.Log("test passed")
}



// *** Testing instruction operand decoding using instruction SET ******


/*
 PRIV: http://fasm.elasticbeanstalk.com/?proj=v4m9zf
 PUB:  http://fasm.elasticbeanstalk.com/?proj=xdyc62
 *
0x0000:                     ;*** Testing immediate value operands nad all registers using SET
0x0000: 8c01                    set a,2
0x0001: 9021                    set b,3
0x0002: 8041                    set c,-1
0x0003: f061                    set x,27
0x0004: f481                    set y,28
0x0005: f8a1                    set z,29
0x0006: 80c1                    set i,-1
0x0007: fce1                    set j,30
0x0008:                     ; Checkpoint. cycle=8
 */
func Test_Operands_using_SET(t *testing.T) {
	c := New(); c.Reset()
	c.memory[0] = 0x8c01
	c.memory[1] = 0x9021
	c.memory[2] = 0x8041
	c.memory[3] = 0xf061
	c.memory[4] = 0xf481
	c.memory[5] = 0xf8a1
	c.memory[6] = 0x80c1
	c.memory[7] = 0xfce1
	for i:=0; i<8; i++ {c.Step()}
	if c.cycle != 8 {t.Errorf("Fail: cycle is %d and should be 8\n", c.cycle)}
	if c.regPC != 8 {t.Errorf("Fail: PC is %#.4x and should be 0x0008\n", c.regPC)}
	if c.regA  != 2 {t.Errorf("Fail: A  is %#.4x and should be 0x0002\n", c.regA)}
	if c.regB  != 3 {t.Errorf("Fail: B  is %#.4x and should be 0x0003\n", c.regB)}
	if c.regC  != 0xffff {t.Errorf("Fail: C  is %#.4x and should be 0xffff\n", c.regC)}
	if c.regX  != 27 {t.Errorf("Fail: X  is %#.4x and should be 0x001b\n", c.regX)}
	if c.regY  != 28 {t.Errorf("Fail: Y  is %#.4x and should be 0x001c\n", c.regY)}
	if c.regZ  != 29 {t.Errorf("Fail: Z  is %#.4x and should be 0x001d\n", c.regZ)}
	if c.regI  != 0xffff {t.Errorf("Fail: I  is %#.4x and should be 0xffff\n", c.regI)}
	if c.regJ  != 30 {t.Errorf("Fail: J  is %#.4x and should be 0x001e\n", c.regJ)}
}

/* Test register operands using SET
 PRIV: http://fasm.elasticbeanstalk.com/?proj=f0wmc2
 PUB:  http://fasm.elasticbeanstalk.com/?proj=nphs6y
 *
0x0000:                     ; Test Instruction SET b,a
0x0000: 0021                    set b,a
 */
func Test_Instruction_SET_B_A(t *testing.T) {
	c := New(); c.Reset()
	// Load memory
	c.memory[0] = 0x0021
	c.regA = 1; c.regB = 0	// Set registers
	c.Step()
	if c.regPC != 1 {t.Errorf("Fail: PC is %#.4x and should be 0x0001\n", c.regPC)}
	if c.regB  != 1 {t.Errorf("Fail: B  is %#.4x and should be 0x0001\n", c.regB)}
	if c.cycle != 1 {t.Errorf("Fail: cycle is %d and should be 1\n", c.cycle)}
}

/* RFH: Test Register Indirect operand
 PRIV: http://fasm.elasticbeanstalk.com/?proj=1jl8mt
 PUB:  http://fasm.elasticbeanstalk.com/?proj=7hljh8
 *
0x0000:                     ; Test register indirect operand using SET
0x0000: 0541                    set [c],b
 */	     
func Test_Instruction_SET_iC_B(t *testing.T) {
	c := New(); c.Reset()
	c.regB = 2; c.regC = 100
	c.memory[0] = 0x0541	// set [c],b
	c.Step()
	if c.memory[100] != 2 {t.Errorf("Fail: [100] is %#.4x and should be 0x0002\n", c.memory[100])}
	if c.cycle != 1 {t.Errorf("Fail: cycle is %d and should be 1\n", c.cycle)}
}


// 0x0000:                     ;Testing SET [x+2],c
// 0x0000: 9041                        set c, 3
// 0x0001: 7c61 0064                   set x, 100
// 0x0003: 0a61 0002                   set [x+2],c
func Test_Instruction_SET_iX2_C(t *testing.T) {
	c := New(); c.Reset()
	c.regC = 3; c.regX = 100
	c.memory[0] = 0x0a61	// set [x+2],c
	c.memory[1] = 0x0002
	c.Step()
	if c.cycle != 2 {t.Errorf("Fail: cycle is %d and should be 2\n", c.cycle)}
	if c.memory[102] != 3 {t.Errorf("Fail: [102] is %#.4x and should be %#.4x\n", c.memory[102], 3)}
}


// 0x0000:                     ;Testing PUSH
// 0x0000: 7f61 0040                   set SP, 0x40
// 0x0002: 7c81 2341                   set Y, 0x2341
// 0x0004: 1301                        set PUSH,Y

func Test_Instruction_SET_PUSH_Y(t *testing.T) {
	c := New(); c.Reset()
	c.regY = 0x2341; c.regSP = 0x40
	c.memory[0] = 0x1301	// set push,y
	c.Step()
	if c.cycle != 1 {t.Errorf("Fail: cycle is %d and should be 1\n", c.cycle)}
	if c.memory[0x3f] != 0x2341 {t.Errorf("Fail: [0x3f] is %#.4x and should be 0x2341\n", c.memory[0x3f])}
}


/*
 PRIV: http://fasm.elasticbeanstalk.com/?proj=xz84cd
 PUB:  http://fasm.elasticbeanstalk.com/?proj=kqvv6q
 *
0x0000:                     ; Testing PEEK
0x0000: 7f61 0040               set SP, 0x40    ;cyc=2
0x0002: 64a1                    set z,peek      ;cyc=1
0x0003:                     ; Checkpoint: cycle=3
 */
func Test_Instruction_SET_Z_PEEK(t *testing.T) {
	c := New(); c.Reset()
	c.memory[0] = 0x7f61; c.memory[1] = 0x0040 // set sp,0x40
	c.memory[2] = 0x64a1			   // set z,peek
	c.memory[0x40] = 0x3415
	c.Step();c.Step()
	if c.cycle != 3 {t.Errorf("Fail: cycle is %d and should be 3\n", c.cycle)}
	if c.regZ  != 0x3415 {t.Errorf("Fail: Z  is %#.4x and should be 0x3415\n", c.regZ)}
	if c.regSP != 0x0040 {t.Errorf("Fail: SP is %#.4x and should be 0x0040\n", c.regSP)}
}

/*
0x0000:                     ; Testing PICK n argument using SET instruction:
0x0000: 7fc1 6347 0043              set [0x43], 0x6347
0x0003: 7f61 0040                   set sp, 0x40
0x0005: 6801 0003                   set a,[sp+3]
 */
func Test_Instruction_SET_A_SP3(t *testing.T) {
	c := New(); c.Reset()
	c.memory[0x43] = 0x6347
	c.regSP = 0x40
	c.memory[0] = 0x6801; c.memory[1] = 0x0003 // set a,[sp+3]
	c.Step()
	if c.cycle != 2 {t.Errorf("Fail: cycle is %d and should be 2\n", c.cycle)}
	if c.regA  != 0x6347 {t.Errorf("Fail: A  is %#.4x and should be 0x6347\n", c.regA)}
}


/*
0x0000:                     ; Testing SP as an source and destination argument using SET instruction:
0x0000: df61                        set sp, 0x16
0x0001: 6c01                        set a, sp
*/
func Test_Instruction_SET_SP(t *testing.T) {
	c := New()
	c.regSP = 0x40
	c.memory[0] = 0xdf61	// set sp,0x16
	c.memory[1] = 0x6c01	// set a,sp
	c.regPC = 0
	c.Step()
	if c.regSP != 0x0016 {
		t.Error("Fail: SP is", c.regSP, "and should be 0x0016")
	} else {
		t.Log("Pass: SET SP,0x16")
	}
	c.Step()
	if c.regA != 0x0016 {
		t.Error("Fail: A is", c.regA, "and should be 0x0016")
	} else {
		t.Log("Pass: SET A,SP")
	}
}

/*
0x0000:                     ; Testing PC as an source and destination argument using SET instruction:
0x0000: 7001                        set a,pc
0x0001: 8781                        set pc,0
*/
func Test_Instruction_SET_PC(t *testing.T) {
	c := New()
	c.memory[0] = 0x7001	// set a,pc
	c.memory[1] = 0x8781	// set pc,0
	c.regPC = 0
	c.Step()
	if c.regA != 0x0001 {
		t.Error("Fail: A is", c.regA, "and should be 0x0001")
	} else {
		t.Log("Pass: SET A,PC")
	}
	c.Step()
	if c.regPC != 0x0000 {
		t.Error("Fail: PC is", c.regPC, "and should be 0x0000")
	} else {
		t.Log("Pass: SET PC,0")
	}
}

/*
0x0000:                     ; Testing EX as an source and destination argument using SET instruction:
0x0000: 9fa1                        set ex,6
0x0001: 7401                        set a,ex
 */
func Test_Istruction_SET_EX(t *testing.T) {
	c := New()
	c.memory[0] = 0x9fa1	// set ex,6
	c.memory[1] = 0x7401	// set a,ex
	c.regPC = 0
	c.Step()
	if c.regEX != 0x0006 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x0006")
	} else {
		t.Log("Pass: SET EX,6")
	}
	c.Step()
	if c.regA != 0x0006 {
		t.Error("Fail: A is", c.regA, "and should be 0x0006")
	} else {
		t.Log("Pass: SET A,EX")
	}
}


/*
0x0000:                     ; Testing [] as an source and destination argument using SET instruction:
0x0000: 7801 0002                   set a,[0x0002]
0x0002: 03c1 0010                   set [0x0010],a
*/
func Test_Instruction_SET_long_address(t *testing.T) {
	c := New()
	c.memory[0] = 0x7801
	c.memory[1] = 0x0002
	c.memory[2] = 0x03c1
	c.memory[3] = 0x0010
	c.regPC = 0
	c.Step()
	if c.regA != 0x03c1 {
		t.Error("Fail: A is", c.regA, "and should be 0x03c1")
	} else {
		t.Log("Pass: SET A,[0x0002]")
	}
	c.Step()
	if c.memory[0x0010] != 0x03c1 {
		t.Error("Fail: [0x0010] is", c.memory[0x0010], "and should be 0x03c1")
	} else {
		t.Log("Pass: SET [0x0010],a")
	}
}

/*
0x0000:                     ; Testing:
0x0000: 7bc1 0000 0010              set [0x0010],[0x0000]
*/

func Test_Instruction_SET_addr_addr(t *testing.T) {
	c := New()
	c.memory[0] = 0x7bc1
	c.memory[1] = 0x0000
	c.memory[2] = 0x0010
	c.regPC = 0
	c.Step()
	if c.memory[0x0010] != 0x7bc1 {
		t.Error("Fail: [0x0010] is", c.memory[0x0010], "and should be 0x7bc1")
	} else {
		t.Log("Pass: SET [0x0010],[0x0000]")
	}
}


/*	     
0x0000:                     ; Testing ADD:
0x0000: 8ba1                        set ex,1
0x0001: 7822 0000                   ADD b,[0x0000]
0x0003: 0422                        ADD b,b
*/
func Test_Instruction_ADD(t *testing.T) {
	c := New()
	c.memory[0] = 0x8ba1
	c.memory[1] = 0x7822
	c.memory[2] = 0x0000
	c.memory[3] = 0x0422
	c.regPC = 0
	c.Step()
	c.Step()
	if c.regB != 0x8ba1 {
		t.Error("Fail: B is", c.regB, "and should be 0x8ba1")
	}
	if c.regEX != 0 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x0000")
	}
	c.Step()
	if c.regB != 0x1742 {
		t.Error("Fail: B is", c.regB, "and should be 0x1742")
	}
	if c.regEX != 1 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x0001")
	}
}

/*
0x0000:                     ; Testing SUB
0x0000: 8ba1                    set ex,1
0x0001: 7841 0000               set c, [0]
0x0003: 7843 0001               sub c, [1]
0x0005: 7843 0003               sub c, [3]
*/
func Test_Instruction_SUB(t *testing.T) {
	c := New()
	c.memory[0] = 0x8ba1
	c.memory[1] = 0x7841
	c.memory[2] = 0x0000
	c.memory[3] = 0x7843
	c.memory[4] = 0x0001
	c.memory[5] = 0x7843
	c.memory[6] = 0x0003
	c.regPC = 0
	c.Step()
	c.Step()
	if c.regC != 0x8ba1 {
		t.Error("Fail: C is", c.regC, "and should be 0x8ba1")
	}
	if c.regEX != 1 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x0001")
	}
	c.Step()
	if c.regC != 0x1360 {
		t.Error("Fail: C is", c.regC, "and should be 0x1360")
	}
	if c.regEX != 0 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x0000")
	}
	c.Step()
	if c.regC != 0x9b1d {
		t.Error("Fail: C is", c.regC, "and should be 0x9b1d")
	}
	if c.regEX != 0xffff {
		t.Error("Fail: EX is", c.regEX, "and should be 0xffff")
	}
}
 

// 2	00100	MUL b,a	b*a->b, unsigned; EX
/*	     
0x0000:                     ; Testing MUL
0x0000: 8ba1                    set ex,1
0x0001: 7c61 1234               set x, 0x1234
0x0003: 7c64 0100               mul x, 0x0100
*/
func Test_Instruction_MUL_1(t *testing.T) {
	c := New()
	c.memory[0] = 0x8ba1
	c.memory[1] = 0x7c61
	c.memory[2] = 0x1234
	c.memory[3] = 0x7c64
	c.memory[4] = 0x0100
	c.regPC = 0
	c.Step()
	c.Step()
	c.Step()
	if c.regX != 0x3400 {
		t.Error("Fail: X is", c.regX, "and should be 0x3400")
	}
	if c.regEX != 0x0012 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x0012")
	}
}

/*
0x0000:                     ; Testing MUL
0x0000: 8ba1                    set ex,1
0x0001: 7c61 fffe               set x, -2
0x0003: 7c64 fffd               mul x, -3
*/
func Test_Instruction_MUL_2(t *testing.T) {
	c := New()
	c.memory[0] = 0x8ba1
	c.memory[1] = 0x7c61
	c.memory[2] = 0xfffe
	c.memory[3] = 0x7c64
	c.memory[4] = 0xfffd
	c.regPC = 0
	c.Step()
	c.Step()
	c.Step()
	if c.regX != 0x0006 {
		t.Error("Fail: X is", c.regX, "and should be 0x0006")
	}
	if c.regEX != 0xfffb {
		t.Error("Fail: EX is", c.regEX, "and should be 0xfffb")
	}
}

/*
0x0000:                     ; Testing MLI
0x0000: 7c61 fffe               set x, -2
0x0002: 7c65 fffd               mli x, -3
0x0004:                     ; F1DE: b=6, ex=0
0x0004:                     ; dcpu.ru: b=6, ex=0
*/
func Test_Instruction_MLI_m2_m3(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c61
	c.memory[1] = 0xfffe
	c.memory[2] = 0x7c65
	c.memory[3] = 0xfffd
	c.regPC = 0
	c.Step(); c.Step()
	if c.regX != 0x0006 {
		t.Error("Fail: X is", c.regX, "and should be 0x0006")
	}
	if c.regEX != 0x0000 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x0000")
	}
}

/*
0x0000:                     ; Testing DIV 10/4
0x0000: ac21                    set b, 10
0x0001: 9426                    div b, 4
0x0002:                     ; b=2, ex=0x8000
*/
func Test_Instruction_DIV_10_4(t *testing.T) {
	c := New()
	c.memory[0] = 0xac21
	c.memory[1] = 0x9426
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 2 {
		t.Error("Fail: B is", c.regB, "and should be 2")
	}
	if c.regEX != 0x8000 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x8000")
	}
}

/*
0x0000:                     ; Testing DIV
0x0000: f821                    set b, 29
0x0001: 9826                    div b, 5
*/
func Test_Instruction_DIV_29_5(t *testing.T) {
	c := New()
	c.memory[0] = 0xf821
	c.memory[1] = 0x9826
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0x0005 {
		t.Error("Fail: B is", c.regB, "and should be 0x0005")
	}
	if c.regEX != 0xcccc {
		t.Error("Fail: EX is", c.regEX, "and should be 0xcccc")
	}
}

/*
0x0000:                     ; Testing DIV 65001/100
0x0000: 7c21 fde9               set b, 65001
0x0002: 7c26 0064               div b, 100
0x0004:                     ; F1DE: b=0x28a=650, ex=0xa667
0x0004:                     ; dcpu.ru: b=0x28a, ex=0x28f
*/
func Test_Instruction_DIV_65001_100(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0xfde9
	c.memory[2] = 0x7c26
	c.memory[3] = 0x0064
	c.regPC = 0
	c.Step()
	c.Step()
	if c.regB != 650 {
		t.Error("Fail: B is", c.regB, "and should be 650")
	}
	if c.regEX != 0x28f {
		t.Error("Fail: EX is", c.regEX, "and should be 0x28f")
	}
}

/*
0x0000:                     ; Testing DIV 3/0
0x0000: 9021                    set b, 3
0x0001: 8426                    div b, 0
0x0002:                     ; F1DE: b=0, ex=0
0x0002:                     ; dcpu.ru: b=0, ex=0
*/
func Test_Instruction_DIV_3_0(t *testing.T) {
	c := New()
	c.memory[0] = 0x9021
	c.memory[1] = 0x8426
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0 {
		t.Error("Fail: B is", c.regB, "and should be 0")
	}
	if c.regEX != 0 {
		t.Error("Fail: EX is", c.regEX, "and should be 0")
	}
}



/*
0x0000:                     ; Testing DVI
0x0000: f821                    set b, 29
0x0001: 9027                    dvi b, 3
0x0002:                     ; b=9, ex=0xaaaa
*/
func Test_Instruction_DVI_29_3(t *testing.T) {
	c := New()
	c.memory[0] = 0xf821
	c.memory[1] = 0x9027
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 9 {
		t.Error("Fail: B is", c.regB, "and should be 9")
	}
	if c.regEX != 0xaaaa {
		t.Error("Fail: EX is", c.regEX, "and should be 0xaaaa")
	}
}


/*
0x0000:                     ; Testing DVI -33/8
0x0000: 7c21 ffdf               set b, -33  ;=0xffdf=0b1...1 1101 1111
0x0002: a427                    dvi b, 8
0x0003:                     ; F1DE: b=0xfffb=-5, ex=0xe000
0x0003:                     ; dcpu.ru: b=0xfffc=-4, ex=0xe000
*/
func Test_Instruction_DVI_m33_8(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0xffdf
	c.memory[2] = 0xa427
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0xfffc {
		t.Error("Fail: B is", c.regB, "and should be 0xfffc")
	}
	if c.regEX != 0xe000 {
		t.Error("Fail: EX is", c.regEX, "and should be 0xe000")
	}
}


/*
0x0000:                     ; Testing DVI
0x0000: 7c21 ffde               set b, -34
0x0002: 7c27 fffc               dvi b, -4
0x0004:                     ; b=8, ex=0x8000
*/
func Test_Instruction_DVI_m34_m4(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0xffde
	c.memory[2] = 0x7c27
	c.memory[3] = 0xfffc
	c.regPC = 0
	c.Step()
	c.Step()
	if c.regB != 8 {
		t.Error("Fail: B is", c.regB, "and should be 8")
	}
	if c.regEX != 0x8000 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x8000")
	}
}


/*
0x0000:                     ; Testing DVI -6/4
0x0000: 7c21 fffa               set b, -6
0x0002: 9427                    dvi b, 4
0x0003:                     ; F1DE: b=0xfffe=-2, ex=0x8000
0x0003:                     ; dcpu.ru: b=0xffff=-1, ex=0x8000
*/
func Test_Instruction_DVI_m6_4(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0xfffa
	c.memory[2] = 0x9427
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0xffff {
		t.Error("Fail: B is", c.regB, "and should be 0xffff")
	}
	if c.regEX != 0x8000 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x8000")
	}
}

/*
0x0000:                     ; Testing DVI -6/3
0x0000: 7c21 fffa               set b,-6
0x0002: 9027                    dvi b, 3
*/
func Test_Instruction_DVI_m6_3(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0xfffa
	c.memory[2] = 0x9027
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0xfffe {
		t.Error("Fail: B is", c.regB, "and should be 0xfffe")
	}
	if c.regEX != 0x0000 {
		t.Error("Fail: EX is", c.regEX, "and should be 0x0000")
	}
}


/*
0x0000:                     ; Testing DVI 2/0
0x0000: 8c21                    set b,2
0x0001: 8427                    dvi b,0
*/
func Test_Instruction_DVI_2_0(t *testing.T) {
	c := New()
	c.memory[0] = 0x8c21
	c.memory[1] = 0x8427
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0 {
		t.Error("Fail: B is", c.regB, "and should be 0")
	}
	if c.regEX != 0 {
		t.Error("Fail: EX is", c.regEX, "and should be 0")
	}
}


/*
0x0000:                     ; Testing MOD 12/5
0x0000: b421                    set b,12
0x0001: 9828                    mod b,5
0x0002:                     ; F1DE:  b=2
*/
func Test_Instruction_MOD_12_5(t *testing.T) {
	c := New()
	c.memory[0] = 0xb421
	c.memory[1] = 0x9828
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 2 {
		t.Error("Fail: B is", c.regB, "and should be 2")
	}
}

/*
0x0000:                     ; Testing MOD 12/0
0x0000: b421                    set b,12
0x0001: 8428                    mod b,0
0x0002:                     ; F1DE:  b=0
*/
func Test_Instruction_MOD_12_0(t *testing.T) {
	c := New()
	c.memory[0] = 0xb421
	c.memory[1] = 0x8428
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0 {
		t.Error("Fail: B is", c.regB, "and should be 0")
	}
}


/*
0x0000:                     ; Testing instruction MDI 12,0
0x0000: b421                    set b,12
0x0001: 8429                    mdi b,0
0x0002:                     ; F1DE:  b=0
*/
func Test_Instruction_MDI_12_0(t *testing.T) {
	c := New()
	c.memory[0] = 0xb421
	c.memory[1] = 0x8429
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0 {
		t.Error("Fail: B is", c.regB, "and should be 0")
	}
}

/*
0x0000:                     ; Testing instruction MDI -7,4
0x0000: 7c21 fff9               set b,-7
0x0002: 9429                    mdi b,4
0x0003:                     ; F1DE:  b=0xfffd=-3
*/
func Test_Instruction_MDI_m7_4(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0xfff9
	c.memory[2] = 0x9429
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0xfffd {
		t.Error("Fail: B is", c.regB, "and should be 0xfffd")
	}
}

/*
0x0000:                     ; Testing instruction MDI -7,-4
0x0000: 7c21 fff9               set b,-7
0x0002: 7c29 fffc               mdi b,-4
0x0004:                     ; F1DE:     b=0xfffd=-3
0x0004:                     ; lag.net:  b=0xfffd=-3
0x0004:                     ; aws.johnmccann.me: b=0xfffd=-3
*/
func Test_Instruction_MDI_m7_m4(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0xfff9
	c.memory[2] = 0x7c29
	c.memory[3] = 0xfffc
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0xfffd {
		t.Error("Fail: B is", c.regB, "and should be 0xfffd")
	}
}

/*
0x0000:                     ; Testing instruction AND 0xacac, 0xcaca
0x0000: 7c21 acac               set b,0xacac
0x0002: 7c2a caca               and b,0xcaca
0x0004:                     ; b=0x8888
0x0004:                     ; F1DE:     b=0x8888
*/
func Test_Instruction_AND_acac_caca(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0xacac
	c.memory[2] = 0x7c2a
	c.memory[3] = 0xcaca
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0x8888 {
		t.Error("Fail: B is", c.regB, "and should be 0x8888")
	}
}

/*
0x0000:                     ; Testing instruction BOR 0xacac, 0xcaca
0x0000: 7c21 acac               set b,0xacac
0x0002: 7c2b caca               bor b,0xcaca
0x0004:                     ; b=0xeeee
0x0004:                     ; F1DE:     b=0xeeee
*/
func Test_Instruction_BOR_acac_caca(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0xacac
	c.memory[2] = 0x7c2b
	c.memory[3] = 0xcaca
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0xeeee {
		t.Errorf("Fail: B is %#x and should be 0xeeee\n", c.regB)
	}
}


/*
0x0000:                     ; Testing instruction XOR 0xacac, 0xcaca
0x0000: 7c21 acac               set b,0xacac
0x0002: 7c2c caca               xor b,0xcaca
0x0004:                     ; b=0x6666
0x0004:                     ; F1DE:     b=0x6666
*/
func Test_Instruction_XOR_acac_caca(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0xacac
	c.memory[2] = 0x7c2c
	c.memory[3] = 0xcaca
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB != 0x6666 {
		t.Errorf("Fail: B is %#x and should be 0x6666\n", c.regB)
	}
}

/*
 * b= b>>>a; sets EX to ((b<<16)>>a)&0xffff (logical shift)
0x0000:                     ; Testing instruction SHR 0x1234, 8
0x0000: 7c21 1234               set b,0x1234
0x0002: a42d                    shr b,8
0x0003:                     ; result:   b=0x0012, ex=3400
0x0003:                     ; F1DE:     b=0x0012, ex=2400
*/
func Test_Instruction_SHR_x1234_8(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0x1234
	c.memory[2] = 0xa42d
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB  != 0x0012 {t.Errorf("Fail: B  is %#.4x and should be 0x0012\n", c.regB)}
	if c.regEX != 0x3400 {t.Errorf("Fail: EX is %#.4x and should be 0x3400\n", c.regEX)}
}

/*
0x0000:                     ; Testing instruction SHR 0x8765, 8
0x0000: 7c21 8765               set b,0x8765
0x0002: a42d                    shr b,8
0x0003:                     ; result:   b=0x0087, ex=6500
0x0003:                     ; F1DE:     b=0x0087, ex=6500
*/
func Test_Instruction_SHR_x8765_8(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0x8765
	c.memory[2] = 0xa42d
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB  != 0x0087 {t.Errorf("Fail: B  is %#.4x and should be 0x0087\n", c.regB)}
	if c.regEX != 0x6500 {t.Errorf("Fail: EX is %#.4x and should be 0x6500\n", c.regEX)}
}


/* RFH: Instruction Testing ASR
RW http://fasm.elasticbeanstalk.com/?proj=rr61qj
RO http://fasm.elasticbeanstalk.com/?proj=rrlxxm
 *
0x0000:                     ; Testing instruction ASR 0x1235, 8
0x0000: 7c21 1235               set b,0x1235
0x0002: a42e                    asr b,8
0x0003:                     ; result:   b=0x0012, ex=0x3500
0x0003:                     ; F1DE      b=0x0012, ex=0x3500
0x0003: 7c21 8765               set b,0x8765
0x0005: a42e                    asr b,8
0x0006:                     ; result:   b=0xff87, ex=6500
0x0006:                     ; F1DE:     b=0xff87, ex=6500
 */
func Test_Instruction_ASR(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0x1235
	c.memory[2] = 0xa42e
	c.memory[3] = 0x7c21
	c.memory[4] = 0x8765
	c.memory[5] = 0xa42e
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB  != 0x0012 {t.Errorf("Fail: B  is %#.4x and should be 0x0012\n", c.regB)}
	if c.regEX != 0x3500 {t.Errorf("Fail: EX is %#.4x and should be 0x3500\n", c.regEX)}
	c.Step(); c.Step()
	if c.regB  != 0xff87 {t.Errorf("Fail: B  is %#.4x and should be 0xff87\n", c.regB)}
	if c.regEX != 0x6500 {t.Errorf("Fail: EX is %#.4x and should be 0x6500\n", c.regEX)}
}


/* Testing SHL
RW http://fasm.elasticbeanstalk.com/?proj=jvy2b7
RO http://fasm.elasticbeanstalk.com/?proj=dgyfqj
 *
0x0000:                     ; Testing instruction SHL
0x0000: 7c21 9235               set b,0x9235
0x0002: 942f                    shl b,4
0x0003:                     ; result:   b=0x2350, ex=0x0009
0x0003:                     ; FID1:     b=0x2350, ex=0x0009
 */
func Test_Instruction_SHL_x9235_4(t *testing.T) {
	c := New()
	c.memory[0] = 0x7c21
	c.memory[1] = 0x9235
	c.memory[2] = 0x942f
	c.regPC = 0
	c.Step(); c.Step()
	if c.regB  != 0x2350 {t.Errorf("Fail: B  is %#.4x and should be 0x2350\n", c.regB)}
	if c.regEX != 0x0009 {t.Errorf("Fail: EX is %#.4x and should be 0x0009\n", c.regEX)}
}


/* Testing IFB
 PRIV: http://fasm.elasticbeanstalk.com/?proj=90kfw4
 PUB:  http://fasm.elasticbeanstalk.com/?proj=jh233v
 *
0x0000:                     ; Testing Instruction IFB
0x0000:                     ; True -- NoSkip
0x0000: 9c21                    set b,0x0006        ;cyc=1
0x0001: 8c30                    ifb b,0x0002        ;cyc=2+0
0x0002: 8801                    set a,1 ;must be executed; cyc=1
0x0003:                     ; result:   a=1
0x0003:                     ; F1DE:     a=1
0x0003:                     ; False -- Skip
0x0003: a430                    ifb b,0x0008        ;cyc=2+1
0x0004: 7c01 2222               set a,0x2222 ;must NOT be executed; cyc=0
0x0006:                     ; result:   a=1
0x0006:                     ; F1DE:     a=1
0x0006: 9041                    set c,3             ;cyc=1
 */
func Test_Instruction_IFB(t *testing.T) {
	c := New(); c.Reset()
	c.memory[0] = 0x9c21
	c.memory[1] = 0x8c30
	c.memory[2] = 0x8801
	c.memory[3] = 0xa430
	c.memory[4] = 0x7c01; c.memory[5] = 0x2222
	c.memory[6] = 0x9041
	c.Step(); c.Step()	// set,ifb
	if c.cycle != 3 {t.Errorf("Fail: cycle is %d and should be 3\n", c.cycle)}
	if c.regPC != 2 {t.Errorf("Fail: PC is %#.4x and should be 0x0002\n", c.regPC)}
	c.Step()		// set a,1
	if c.cycle != 4 {t.Errorf("Fail: cycle is %d and should be 4\n", c.cycle)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
	c.Step()		// ifb b,0x0008
	if c.cycle != 7 {t.Errorf("Fail: cycle is %d and should be 7\n", c.cycle)}
	if c.regPC != 6 {t.Errorf("Fail: PC is %#.4x and should be 0x0006\n", c.regPC)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
	c.Step()		// set c,3
	if c.cycle != 8 {t.Errorf("Fail: cycle is %d and should be 8\n", c.cycle)}
	if c.regPC != 7 {t.Errorf("Fail: PC is %#.4x and should be 0x0007\n", c.regPC)}
	if c.regC  != 3 {t.Errorf("Fail: C  is %#.4x and should be 0x0003\n", c.regC)}
}



/* Testing IFC
 PRIV: http://fasm.elasticbeanstalk.com/?proj=r1f09n
 PUB:  http://fasm.elasticbeanstalk.com/?proj=yrzx7p
 *      
0x0000:                     ; Testing Instruction IFC
0x0000:                     ; True -- NoSkip
0x0000: 9c21                    set b,0x0006        ;cyc=1
0x0001: a431                    ifc b,0x0008        ;cyc=2+0
0x0002: 8801                    set a,1 ;must be executed; cyc=1
0x0003:                     ; result:   a=1
0x0003:                     ; F1DE:     a=1
0x0003:                     ; False -- Skip
0x0003: 9431                    ifc b,0x0004        ;cyc=2+1
0x0004: 7c01 2222               set a,0x2222 ;must NOT be executed; cyc=0
0x0006:                     ; result:   a=1
0x0006:                     ; F1DE:     a=1
0x0006: 9041                    set c,3             ;cyc=1
 */
func Test_Instruction_IFC(t *testing.T) {
	c := New(); c.Reset()
	c.memory[0] = 0x9c21
	c.memory[1] = 0xa431
	c.memory[2] = 0x8801
	c.memory[3] = 0x9431
	c.memory[4] = 0x7c01; c.memory[5] = 0x2222
	c.memory[6] = 0x9041
	c.Step(); c.Step()	// set,ifc
	if c.cycle != 3 {t.Errorf("Fail: cycle is %d and should be 3\n", c.cycle)}
	if c.regPC != 2 {t.Errorf("Fail: PC is %#.4x and should be 0x0002\n", c.regPC)}
	c.Step()		// set a,1
	if c.cycle != 4 {t.Errorf("Fail: cycle is %d and should be 4\n", c.cycle)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
	c.Step()		// ifc b,0x0008
	if c.cycle != 7 {t.Errorf("Fail: cycle is %d and should be 7\n", c.cycle)}
	if c.regPC != 6 {t.Errorf("Fail: PC is %#.4x and should be 0x0006\n", c.regPC)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
	c.Step()		// set c,3
	if c.cycle != 8 {t.Errorf("Fail: cycle is %d and should be 8\n", c.cycle)}
	if c.regPC != 7 {t.Errorf("Fail: PC is %#.4x and should be 0x0007\n", c.regPC)}
	if c.regC  != 3 {t.Errorf("Fail: C  is %#.4x and should be 0x0003\n", c.regC)}
}

/*
 PRIV: http://fasm.elasticbeanstalk.com/?proj=r0qg96
 PUB:  http://fasm.elasticbeanstalk.com/?proj=rwcszg
 *
0x0000:                     ; Testing Instruction IFE
0x0000:                     ; True -- NoSkip
0x0000: 9c21                    set b,6        ;cyc=1
0x0001: 9c32                    ife b,6        ;cyc=2+0
0x0002: 8801                    set a,1 ;must be executed; cyc=1
0x0003:                     ; Checkpoint: a=1
0x0003: 9832                    ife b,5        ;cyc=2+1
0x0004: 8c01                    set a,2 ;must NOT be executed; cyc=0
0x0005:                     ; Checkpoint: a=1
 */
func Test_Instruction_IFE(t *testing.T) {
	c := New(); c.Reset()
	c.memory[0] = 0x9c21
	c.memory[1] = 0x9c32
	c.memory[2] = 0x8801
	c.memory[3] = 0x9832
	c.memory[4] = 0x8c01
	c.Step(); c.Step(); c.Step()	// set,ife,set
	if c.cycle != 4 {t.Errorf("Fail: cycle is %d and should be 4\n", c.cycle)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
	c.Step() // ife, skip set
	if c.cycle != 7 {t.Errorf("Fail: cycle is %d and should be 7\n", c.cycle)}
	if c.regPC != 5 {t.Errorf("Fail: PC is %#.4x and should be 0x0005\n", c.regPC)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
}


/*
 PRIV: http://fasm.elasticbeanstalk.com/?proj=y70c28
 PUB:  http://fasm.elasticbeanstalk.com/?proj=09zfn5
 *
0x0000:                     ; Testing Instruction IFN
0x0000:                     ; True -- NoSkip
0x0000: 9c21                    set b,6        ;cyc=1
0x0001: 9833                    ifn b,5        ;cyc=2+1
0x0002: 8801                    set a,1 ;must be executed; cyc=
0x0003:                     ; Checkpoint: a=1
0x0003: 9c33                    ifn b,6        ;cyc=2+1
0x0004: 8c01                    set a,2 ;must NOT be executed; cyc=0
0x0005:                     ; Checkpoint: a=1
 */
func Test_Instruction_IFN(t *testing.T) {
	c := New(); c.Reset()
	c.memory[0] = 0x9c21
	c.memory[1] = 0x9833
	c.memory[2] = 0x8801
	c.memory[3] = 0x9c33
	c.memory[4] = 0x8c01
	c.Step(); c.Step(); c.Step()	// set,ifn,set
	if c.cycle != 4 {t.Errorf("Fail: cycle is %d and should be 4\n", c.cycle)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
	c.Step() // ifn, skip set
	if c.cycle != 7 {t.Errorf("Fail: cycle is %d and should be 7\n", c.cycle)}
	if c.regPC != 5 {t.Errorf("Fail: PC is %#.4x and should be 0x0005\n", c.regPC)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
}


/*
 PRIV: http://fasm.elasticbeanstalk.com/?proj=js7gx6
 PUB:  http://fasm.elasticbeanstalk.com/?proj=2tnbrd
 *
0x0000:                     ; Testing Instruction IFG
0x0000:                     ; True -- NoSkip
0x0000: a821                    set b,9        ;cyc=1
0x0001: a434                    ifg b,8        ;cyc=2+0
0x0002: 8801                    set a,1 ;must be executed; cyc=1
0x0003:                     ; Checkpoint: a=1
0x0003: a834                    ifg b,9        ;cyc=2+1
0x0004: 8c01                    set a,2 ;must NOT be executed; cyc=0
0x0005:                     ; Checkpoint a=1
 */
func Test_Instruction_IFG(t *testing.T) {
	c := New(); c.Reset()
	c.memory[0] = 0xa821
	c.memory[1] = 0xa434
	c.memory[2] = 0x8801
	c.memory[3] = 0xa834
	c.memory[4] = 0x8c01
	c.Step(); c.Step(); c.Step()	// set,ifg,set
	if c.cycle != 4 {t.Errorf("Fail: cycle is %d and should be 4\n", c.cycle)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
	c.Step() // ifg, skip set
	if c.cycle != 7 {t.Errorf("Fail: cycle is %d and should be 7\n", c.cycle)}
	if c.regPC != 5 {t.Errorf("Fail: PC is %#.4x and should be 0x0005\n", c.regPC)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
}


/*
 PRIV: http://fasm.elasticbeanstalk.com/?proj=jz7qhb
 PUB:  http://fasm.elasticbeanstalk.com/?proj=w4zscr
 *
0x0000:                     ; Testing Instruction IFA
0x0000:                     ; True -- NoSkip
0x0000: 7c21 fff9               set b,-7        ;cyc=2
0x0002: 7c35 fff8               ifa b,-8        ;cyc=3+0
0x0004: 8801                    set a,1 ;must be executed; cyc=1
0x0005:                     ; Checkpoint: a=1
0x0005: 7c35 fff9               ifa b,-7        ;cyc=3+1
0x0007: 8c01                    set a,2 ;must NOT be executed; cyc=0
0x0008:                     ; Checkpoint a=1
 */
func Test_Instruction_IFA(t *testing.T) {
	c := New(); c.Reset()
	c.memory[0] = 0x7c21; c.memory[1] = 0xfff9
	c.memory[2] = 0x7c35; c.memory[3] = 0xfff8
	c.memory[4] = 0x8801
	c.memory[5] = 0x7c35; c.memory[6] = 0xfff9
	c.memory[7] = 0x8c01
	c.Step(); c.Step(); c.Step()	// set,ifa,set
	if c.cycle != 6 {t.Errorf("Fail: cycle is %d and should be 6\n", c.cycle)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
	c.Step() // ifg, skip set
	if c.cycle != 10 {t.Errorf("Fail: cycle is %d and should be 10\n", c.cycle)}
	if c.regPC != 8 {t.Errorf("Fail: PC is %#.4x and should be 0x0008\n", c.regPC)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
}

/*
 PRIV: http://fasm.elasticbeanstalk.com/?proj=9wmtns
 PUB:  http://fasm.elasticbeanstalk.com/?proj=ps7xk9
 *
0x0000:                     ; Testing Instruction IFL
0x0000:                     ; True -- NoSkip
0x0000: b021                    set b,11        ;cyc=1
0x0001: b436                    ifl b,12        ;cyc=2+0
0x0002: 8801                    set a,1 ;must be executed; cyc=1
0x0003:                     ; Checkpoint: a=1
0x0003: b036                    ifl b,11        ;cyc=2+1
0x0004: 8c01                    set a,2 ;must NOT be executed; cyc=0
0x0005:                     ; Checkpoint a=1
 */
func Test_Instruction_IFL(t *testing.T) {
	c := New(); c.Reset()
	c.memory[0] = 0xb021
	c.memory[1] = 0xb436
	c.memory[2] = 0x8801
	c.memory[3] = 0xb036
	c.memory[4] = 0x8c01
	c.Step(); c.Step(); c.Step()	// set,ifl,set
	if c.cycle != 4 {t.Errorf("Fail: cycle is %d and should be 4\n", c.cycle)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
	c.Step() // ifl, skip set
	if c.cycle != 7 {t.Errorf("Fail: cycle is %d and should be 7\n", c.cycle)}
	if c.regPC != 5 {t.Errorf("Fail: PC is %#.4x and should be 0x0005\n", c.regPC)}
	if c.regA  != 1 {t.Errorf("Fail: A  is %#.4x and should be 0x0001\n", c.regA)}
}
