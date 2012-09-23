Title:		DCPU-16
Author:		Radek Hnilica
RCSId:		$Id: dcpu16.md,v 1.2 2012/09/17 17:07:47 radek Exp radek $
css:		documentation.css
css:			local.css
HTML header:	<meta http-equiv="refresh" content="120"/>
HTML header:	<script type="text/javascript" src="http://cdn.mathjax.org/mathjax/latest/MathJax.js?config=default"></script>


DCPU-16
=======

> $Id: dcpu16.md,v 1.2 2012/09/17 17:07:47 radek Exp radek $

> by Radek Hnilica

> This document is in very early stage, and some parts were not yet
> translated to english.

This document is about DCPU-16 computer as introduced by Markus
Persson also known as Notch.  Markus created description of this
computer as crucial part of his new computer game 0x10c.  My work
started by collecting informations about this machine.

Links:

* [Mojang 0x10<sup>C</sup>](http://0x10c.com)
* [Markus Persson twitter](https://twitter.com/#!/notch)
* [0x10<sup>c</sup> reddit](http://www.reddit.com/r/0x10c/)
* [Forum](http://0x10cforum.com/)
* [DCPU-16 Tools](http://www.redpointsoftware.com.au/dcpu/)
* [noname22/dtools](https://github.com/noname22/dtools)
* [DCPU-16 Assembler & Emulator](http://denull.ru/dcpu/dcpu.htm)
* [DCPU World](http://dcpuworld.com)
* [An Introduction to the DCPU-16](http://www.trillek.net/blog/2012/4/13/article-an-introduction-to-the-dcpu-16.html) by R. Horning, 2012-04-13
* Wikipedia article [0x10c](http://en.wikipedia.org/wiki/0x10c)

Known DCPU-16 specifications:

* [Version 1.7](http://pastebin.com/raw.php?i=Q4JvQvnM)

In this document I want stay as actual as possible.  So I use last
specification known to me.

Note

    gcc -I ../third-party/libtcod-1.5.0/include/ -L ../third-party/libtcod-1.5.0/ *.cpp -o dcpu -ltcod -ltcodgui -ltcodxx

    c++ *.cpp -I <dir containing libtcod.h> /path/to/libtcod.so -Wl,-rpath <dir containing libtcod.so> -o emu


Web based tools:

* [DCPU-16 Assembler & Emulator](http://fingswotidun.com/dcpu16/)
* [deNULL's Web Assembler, Emulator & Disassembler](http://dcpuworld.com/2012/04/denulls-web-assembler-emulator-disassembler/)
* [DCPU-16 Assembler, Emulator & Disassembler](http://dcpu.ru/)
* github: mappum / [DCPU-16](https://github.com/mappum/DCPU-16) -- see 0x10co.de
* [dcpu16-lcc](http://dcpu16-lcc.bugs3.com)
* [DCPU-16 Workbench](http://dcpu16_workbench.whoatemydomain.co.za)
* [0x10<sup>c</sup>Programs](http://0x10c-programs.com/dcpu.php)
* [iammer](http://www.iammer.com/dcpu/dcpu16.html)
* [elasticbeanstalk](http://fasm.elasticbeanstalk.com/?proj=21rnsl)


For Linux:

* github: Adbook / [DCPU-16-tool](https://github.com/Adbook/DCPU-16-tool), C++
* bitbucket: benedek / [DCPU-16](https://bitbucket.org/benedek/dcpu-16/), C++

* CPAN: [CPU-Emulator-DCPU16](http://search.cpan.org/~simonw/CPU-Emulator-DCPU16/), Perl
* github: [dcpu16-go](https://github.com/dcpu16/dcpu16-go), go


Nástroje
--------

* github: Wilduck / [dcpu16-mode](https://github.com/Wilduck/dcpu16-mode) -- An Emacs major mode for editing DCPU-16 Assembly
* [Benedek's Domain - DCPU-16 Emulator and Assembler](http://n.ethz.ch/~vartokb/dcpu.html) -- 
* github: Janiczek / [cfs-dcpu16](https://github.com/Janiczek/cfs-dcpu16)


Vývojové nástroje
-----------------

### Benedek's DCPU-16

Simulátor a assembler.

Kompiluje se bez problémů.  Je třeba nainstalovat knihovnu libsdl.

    # aptitude install libsdl-dev
    $ make
    
### IDE

#### DCpuToolchain

* [DCPU-16 Toolchain](http://dcputoolcha.in/)
* [Automated Linux Builds](http://irc.lysdev.com/dcputoolchain/builds/linux/?C=M;O=D)

### Webová IDE

* [DCPU-16 Assembler, Emulator & Disassembler](http://lag.net/dcpu/dcpu.htm) by deNULL
* [F1DE](http://fasm.elasticbeanstalk.com/)
* [aws.johnmccann](http://aws.johnmccann.me/)
* [dcpubin](http://www.dcpubin.com/)  spec 1.1?
* [0x10co.de](http://0x10co.de/)

### Assemblery

* [F1DE](http://fasm.elasticbeanstalk.com/) -- webový assembler a simulátor

### Překladače jazyka C

* [llvm-dcpu16](https://github.com/llvm-dcpu16/llvm-dcpu16) -- llvm backend for dcpu16

### Forth

* [hellige's goForth](https://github.com/hellige/dcpu)
* [CamelForth-16 for DCPU](https://github.com/dsmvwld/CamelForth-16)

Programy a knihovny
-------------------

* Mappum’s web emulator [0x10co.de](http://0x10co.de/) -- Web Library of programs
* [Your DCPU](http://your-dcpu.com) -- Web library and development
* [CALCULON-3000](http://0x10co.de/dw8km)
* [DCPU-16 APPS](http://www.dcpu16apps.com/)
* github: jlongster / [dcpu-lisp](https://github.com/jlongster/dcpu-lisp)
* [dcpubin](http://www.dcpubin.com/) -- pastebin for dcpu


Programmers model of DCPU16
---------------------------

DCPU-16 is a sixteen bit processor having sixteen bit wide registers,
data bus and also address bus.

### Registers

There are twelve 16 bit register accessible to the programer.  Eight
registers are universal ones, named A, B, C, X, Y, Z, I, J, and four
register has special functions.  These are named PC, SP, EX/OV, IA.
In following table are names of the register, short description, and
number.  The number corresponds with part of the instruction word.

| Name | Number | Description
|:----:|:-------|:------------
| IA   |        | Interrupt Address
| [PC] | 0x1C   | program counter / instruction pointer
| SP   | 0x1B   | stack pointer
| EX/OV| 0x1D   | overflow register
| A    | 0x00   | general purpose register
| B    | 0x01   | general purpose register
| C    | 0x02   | general purpose register
| X    | 0x03   | general purpose register
| Y    | 0x04   | general purpose register
| Z    | 0x05   | general purpose register
| I    | 0x06   | general purpose register, used as a counter by [STI] and [STD] instructions
| J    | 0x07   | general purpose register, used as a counter by [STI] and [STD] instructions
[DCPU-16 Registers]

The instruction set is mostly orthogonal.  That means that for all the
instruction you can use any combination of general purpose registers.
The exception from this are instructions [STI] ans [STD] whose use
implict registers I and J.

> Special registers has special functions!


#### Register PC [PC]

This register stores address of executed instruction.  If we want
break the linear sequence of instruction executed, we can assign a new
valut to this register and thus create something we call JUMP.  What
wil happen is that after finishing actual instruction `SET PC,0x200`,
the next executed instruction is that on the address 0x200.

> *FIXME*: Check if the PC is incremented in early stage of instruction decoding or in later stage.

> *FIXME*: Create an example of "JUMP".

#### Stack pointer SP [SP]

Another special register is Stack Pointer, SP.  It is necessary
register for stack operation PUSH, POP, JSR, and RET.  It allow us to
make subroutines / functions.



Memory
------

Memory module has sixteen bit address bus and sixteen bit data bus.
Memory consist of 0x1000 (65536) cells each one sixteen bit wide.  So
in terms of other units it has 1048576 bits or 131072 octets.  Because
of sixteen bit width it's not safe to speeak in term of Bytes.  In
this document I'm using term Word because it fits best.


### Memory map

Actually I do not know how the memory map will look about.  And
because of the way hardware devices are connected, the memory map will
be at the opinion of the programer.  In the olderspecifications there
was an memory map which looks following

| Address Range | Description
|:-------------:|:-----
| 0000 - 7FFF   | RAM
| 8000 - 817F   | VRAM 32*12 cells
| 8180 - 827F   | CHRAM Character RAM
| 8280          | Video screen border collor
| 9000 - 900F   | Keyboard circular buffer
| 9010 - 9FFF   | reserved
| A000 - FFFF   | Stack

As is seen on the map, first half of the map is dedicated to operating
system and application(s).  Above application RAM are the hardware
devices.  The last part of the ram is dedicated for stack, and
probably heap.  But because there is no fixed hadware mapping, you can
map devices which need some memory as you wish.  The stack and heap
are also matter of programmer choice.

Instructions
------------

DCPU-16 instruction has from one to three words.  First word describes
the instruction itself and following optional words contains data.

There are two basic form of the instruction word.

     15  14  13  12  11  10   9   8   7   6   5   4   3   2   1   0
    +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
    |      SOURCE (a)       |  DESTINATION (b)  |    OPCODE != 0    |
    +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+

If OPCODE is zero, then the instruction is one operand instruction.

    +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
    |      OPERAND (a)      |      OPCODE       | 0   0   0   0   0 |
    +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+

In the source operand is usually refered as *a* and the destination
operand as *b*.


[Operand encoding, 5 or 6 bits]
| C |   code      | popis       |
|:-:|------------:|:------------|
| 0 | `0x00-0x07` | registry A, B, .... v tomto pořadí |
| 0 | `0x08-0x0f` | [registr] -- v registru je adresa  |
| 1 | `0x10-0x17` | [register + next word ]            |
| 0 | `0x18`      | (PUSH / \[--SP]) if in b, or (POP / \[SP++]) if in a
| 0 | `0x19`      | \[SP] / PEEK
| 1 | `0x1a`      | [SP + next word ] / PICK n
| 0 | `0x1b`      | SP
| 0 | `0x1c`      | PC
| 0 | `0x1d`      | EX
| 1 | `0x1e`      | [next word]
| 1 | `0x1f`      | next word (literal)
| 0 | `0x20-0x3f` | litaral value 0xffff-0x1e (-1..30) (literal) only for a

Operand encodings is same for both, source and destination, operands.  Exceptions are:

* literal values, code 0x20-0x3f, are possible only in source operand.  Because of this source operand has six bits and destination only five.
* stack operation, operand code 0x18, is POP in source operand and means PUSH in destination operand


Instruction Set
---------------

| C |opcode | mnemonic | popis
|:-:|:------|:---------|:---------------------
|   | 00000 |          | _special and one operand instructions_
| 1 | 00001 | SET b,a  | a->b
| 2 | 00010 | ADD b,a  | b+a->b; update Overflow
| 2 | 00011 | SUB b,a  | b-a->b; update Overflow
| 2 | 00100 | MUL b,a  | b*a->b, unsigned; EX
| 2 | 00101 | MLI b,a  | like MUL but a and b signed
| 3 | 00110 | DIV b,a  | sets b to b/a, sets EX to ((b<<16)/a)&0xffff. if a==0, sets b and EX to 0 instead. (treats b, a as unsigned)
| 3 | 00111 | DVI b,a  | like DIV, but treat b, a as signed. Rounds towards 0
| 3 | 01000 | MOD b,a  | sets b to b%a. if a==0, sets b to 0 instead.
| 3 | 01001 | MDI b,a  | like MOD, but treat b, a as signed. (MDI -7, 16 == -7)
| 1 | 01010 | AND b,a  | b&a->b
| 1 | 01011 | BOR b,a  | b\|a->b
| 1 | 01100 | XOR b,a  | b^a->b
| 2 | 01101 | SHR b,a  | b= b>>>a; sets EX to ((b<<16)>>a)&0xffff (logical shift)
| 1 | 01110 | ASR b,a  | b= b>>a; sets EX to ((b<<16)>>>a)&0xffff (arithmetic shift) (treats b as signed)
| 1 | 01111 | SHL b,a  | b &larr; b<<a, sets EX to ((b<<a)>>16)&0xffff
|2+ | 10000 | IFB b,a  | performs next instruction only if (b&a)!=0
|2+ | 10001 | IFC b,a  | performs next instruction only if (b&a)==0
|2+ | 10010 | IFE b,a  | performs next instruction only if b==a
|2+ | 10011 | IFN b,a  | performs next instruction only if b!=a
|2+ | 10100 | IFG b,a  | performs next instruction only if b>a
|2+ | 10101 | IFA b,a  | performs next instruction only if b>a (signed)
|2+ | 10110 | IFL b,a  | performs next instruction only if b<a
|2+ | 10111 | IFU b,a  | performs next instruction only if b<a (signed)
|   | 11000 |          |
|   | 11001 |          |
| 3 | 11010 | ADX b,a  | b+a+EX->b, sets EX to 0x0001 if overflow, 0x0000 otherwise
| 3 | 11011 | SBX b,a  | b-a+EX->b, sets EX to 0xFFFF if underflow, 0x0001 if overflow, 0x0000 otherwise
|   | 11100 |          |
|   | 11101 |          |
| 2 | 11110 | STI b,a  | a->b, I++, J++
| 2 | 11111 | STD b,a  | a->b, I--, J--
[Table of two operand instructions]


[Table of one operand instructions]
| C | code  | mnemo | description                                    |
|:-:|:-----:|:------|:-----------------------------------------------|
|   | 00000 |       | _reserved for future expansion_                |
| 3 | 00001 | JSR a | jump to subroutine: PUSH PC, PC=a              |
|   | 00010 |       |                                                |
|   | 00011 |       |                                                |
|   | 00100 |       |                                                |
|   | 00101 |       |                                                |
|   | 00111 |       |                                                |
| 4 | 01000 | INT a | triggers a software interrupt with message a   |
| 1 | 01001 | IAG a | IA->a                                          |
| 1 | 01010 | IAS a | a->IA                                          |
| 3 | 01011 | RFI a |                                                |
| 2 | 01100 | IAQ a |                                                |
|   | 01101 |       |                                                |
|   | 01110 |       |                                                |
|   | 01111 |       |                                                |
| 2 | 10000 | HWN a | sets a to number of connected hardware devices |
| 4 | 10001 | HWQ a | sets A,B,C,X,Y to information about hardware   |
|4+ | 10010 | HWI a | sends an interrupt to hardware a               |
|   | 10011 |       |                                                |
|   | 10100 |       |                                                |
|   | 10101 |       |                                                |
|   | 10110 |       |                                                |
|   | 10111 |       |                                                |
|   | 11000 |       |                                                |
|   | 11001 |       |                                                |
|   | 11010 |       |                                                |
|   | 11011 |       |                                                |
|   | 11100 |       |                                                |
|   | 11101 |       |                                                |
|   | 11110 |       |                                                |
|   | 11111 |       |                                                |


### STI

    STI b,a

> a->b, I++, J++

### STD

    STD b,a

> b->a, I--, J--

Interrupts
----------

There is possibility for hardware devices to send attention signals to
main processor.  This signal is known as interrupt.  Between each
instruction execution there can be processed the interrupt signals, at
most one.

The interrupt is serviced by interrupt routine, and therefore we have
special register IA.  If IA contains zero, it means that the interrupt
processing is disabled.  If we want enable processing of interrupts we
need to set the value of IA to starting address of interrupt routine.

> Interrupt handler should end with [RFI] instruction.

microcode of interrupt

    push PC
    push A
    IA -> PC
    interrupt_message -> A

microcode of RFI

    disable interrupt queueing
    pop A
    pop PC

See also:

* [RFI]
* [IAQ]


Hardware Devices
----------------

DCPU-16 supports up to 65535 connected hardware devices.

For working with hardware there are three instructions: HWN, HWQ and HWI.


### Jednoduchý sčítací program

    ; Basic DCPU adder by Fred
    SET A, 72           ; Store 72 in A
    SET B, 54           ; Store 54 in B
    ADD A, B            ; Add B to A
    :end
    SET PC, end

### Základní aritmetické instrukce

    ADD A, B        ; adds B to A
    SUB A, B
    MUL A, B
    DIV A, B
    MOD A, B

### Binární operace

    SHL A, B        ; shift A left by B bits
    SHR A, B        ; shift A right by B bits
    AND A, B        ; A and B -> A
    BOR A, B        ; A or  B -> A
    XOR A, B        ; A xor B -> A


Hardware
--------


### HMD2043 Harold Media Drive ID=0x74fa4cae, version 0x07c2

* [HIT_HMD2043](https://gist.github.com/2495578)

Řadič připojuje 1.44 MB 3.5" disketovou jednotku.

[Příkazy]
|  A   | název                  | popis |
|:----:|:-----------------------|:------|
|0x0000| QUERY_MEDIA_PRESENT    |       |
|0x0001| QUERY_MEDIA_PARAMETERS |       |
|0x0002| QUERY_DEVICE_FLAGS     |       |
|0x0003| UPDATE_DEVICE_FLAGS    |       |
|0x0004| QUERY_INTERRUPT_TYPE   |       |
|0x0005| SET_INTERRUPT_MESSAGE  |       |
|0x0010| [READ_SECTORS]         |       |
|0x0011| WRITE_SECTORS          |       |
|0xFFFF| QUERY_MEDIA_QUALITY    |       |

#### 0x0010 -- READ_SECTORS [READ_SECTORS]

    B = počáteční sektor
    C = počet sektorů ke čtění
    X = adresa paměti

#### 0x0011 -- WRITE_SECTORS

    B = počáteční sektor
    C = počet sektorů k zápisu
    X = adresa paměti
   

Code reactor
------------

Různé kousky kódu a úvahy.

### Objektové programování

* [OOdevel](http://fasm.elasticbeanstalk.com/?proj=qrctfc)

Volání virtuální funkce podle [ratchetfreak].

    ;set X,*obj ;X <- objekt
    set X,[X+vtableOffset]      ; X <- tabulka metod
    set X,[X+funcOffset]        ; X <- adresa metody
    jsr X


[ratchetfreak]: http://0x10cforum.com/profile/2070695

<!--
Local Variables:
mode: markdown
coding: utf-8
End:
-->
