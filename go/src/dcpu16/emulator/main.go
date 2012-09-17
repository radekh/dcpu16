/*
 * $Header$
 * DCPU16 emulator.
 * Main file.
 * Copyright (c) 2012 Radek Hnilica
 *
 */
package main;

import "fmt";
import "dcpu16/cpu"

func main() {
	fmt.Println("d16e");

	// create new cpu
	cpu := cpu.New()

	// display contents of cpu the primitive way
	fmt.Printf("%#v\n", *cpu)

	// Load binary specified on command line
	//cpu.Load("testfile.bin")
	// Run the binary
	//cpu.Run()
}

/*
 *
 */