package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/cespare/xxhash"
	"github.com/icepic/toy6502"
	"os"
	"path/filepath"
)

//go:embed zp.bin
var zpbuf []byte

//go:embed buff.bin
var buf []byte

func c64format(pet byte) rune {
	return Petmap(pet)
}

func cksumScreen(c *toy6502.CPU) {
	xxsum := xxhash.Sum64(toy6502.GetMemory(c)[0x0400:0x07f7])
	fmt.Printf("checksum:  %x\n", xxsum)
}

func c64findsys(addr int64, size int64, c *toy6502.CPU) uint16 {
	var i, j, sys int64
	for i = addr; i < addr+size; i++ {
		if toy6502.GetMemory(c)[i] != 0x9e {
			continue
		} else {
			for j = 0; j < 5; j++ {
				var value = toy6502.GetMemory(c)[i+j+1]
				if value == 0x20 {
					continue
				}
				if value > 0x2f && value < 0x3a {
					sys = sys * 10
					sys = sys + int64(value-0x30)
				} else {
					break
				}
			}
			//	fmt.Printf("Calculated value: 0x%04x\n", sys)
			if sys <= 65535 {
				return uint16(sys) // the happy case
			} else {
				fmt.Println("bogus value")
				return 0
			}
		}
	}
	return 0
}

func c64load(fname string, c *toy6502.CPU) (uint16, bool) {
	f, err := os.Open(filepath.Clean(fname))
	if err != nil {
		fmt.Println(err)
		return 0, false
	}
	fi, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return 0, false
	}
	if fi.Size() < 3 {
		fmt.Println("Invalid small size")
		return 0, false
	}
	var addrtemp = make([]byte, 2)
	_, err = f.Read(addrtemp)
	if err != nil {
		fmt.Println(err)
		return 0, false
	}

	var newsize = fi.Size() - 2
	var addr int64 = int64(addrtemp[0]) + int64(addrtemp[1])*256
	//	fmt.Printf("C64 Load Addr: 0x%04x, end: 0x%04x\n", addr, addr+newsize)

	if (fi.Size() + addr - 2) > int64(len(toy6502.GetMemory(c))) {
		fmt.Printf(
			"Illegal size %04x, read would be beyond end of memory.\n",
			newsize)
		return 0, false
	}

	var temp = make([]byte, newsize)
	_, err = f.Read(temp)
	if err != nil {
		fmt.Println(err)
		return 0, false
	}
	copy(toy6502.GetMemory(c)[addr:addr+newsize], temp[:])
	return c64findsys(addr, newsize, c), true
}

func loadKernalBasic(c *toy6502.CPU, kb bool, zp bool, scr bool) {
	if kb == true {
		Loadfileat("basic", c, 0xa000)
		Loadfileat("kernal", c, 0xe000)
	}
	if zp == true {
		Loadfileat("zp.bin", c, 0x00)
		Loadfileat("buff.bin", c, 0x200)
	}
	if scr == true {
		Loadfileat("screen.bin", c, 0x400)
	}
}

func Loadfileat(fname string, c *toy6502.CPU, addr int64) bool {

	f, err := os.Open(filepath.Clean(fname))
	if err != nil {
		fmt.Println(err)
		return false
	}
	fi, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if fi.Size() > int64(len(toy6502.GetMemory(c))) {
		fmt.Println("invalid ram image size")
		return false
	}
	if (fi.Size() + addr) > int64(len(toy6502.GetMemory(c))) {
		fmt.Printf(
			"Illegal size %x, read would end up beyond end of memory.\n",
			fi.Size())
		return false
	}
	var temp = make([]byte, fi.Size())
	//	_, err = f.Read(c.memory)
	_, err = f.Read(temp)
	if err != nil {
		fmt.Println(err)
		return false
	}
	copy(toy6502.GetMemory(c)[addr:addr+fi.Size()], temp[:])
	fmt.Printf("Loaded %v at %x\n", fname, addr)
	return true
}

func main() {

	var shouldexit bool = false
	var quiet bool
	var kerbas bool
	var zpbuff bool
	var loadscreen bool
	var disasm bool
	var showstack bool
	var printscreen bool
	var printscreenhex bool
	var maxinst uint64

	c := toy6502.New()

	myFilename := flag.String("filename", "a.out", "c64.prg to load and run.")
	flag.BoolVar(&quiet, "q", false, "Quiet mode, only print the checksum.")
	flag.BoolVar(&kerbas, "k", false, "Load files 'basic' and 'kernal' into $a000 and $e000")
	flag.BoolVar(&zpbuff, "z", false, "Load files 'zp.bin' and 'buff.bin' into $00 and $200")
	flag.BoolVar(&loadscreen, "sc", false, "Load files 'screen.bin' into $400")
	flag.BoolVar(&disasm, "d", false, "Disassemble each instruction during run")
	flag.BoolVar(&showstack, "ss", false, "Hexdump stack after exit")
	flag.BoolVar(&printscreen, "ps", false, "Print simple screen rendering after exit")
	flag.BoolVar(&printscreenhex, "psx", false, "Print top part of screen in hex after exit")
	flag.Uint64Var(&maxinst, "max", 1000000000,
		"Maximum number of instructions to execute before aborting.")

	flag.Parse()

	if kerbas == true || zpbuff == true || loadscreen == true {
		loadKernalBasic(c, kerbas, zpbuff, loadscreen)
	}
	startaddr, ok := c64load(*myFilename, c)

	if quiet == false {
		fmt.Printf("Start address: 0x%04X\n", startaddr)
	}

	if ok == false {
		fmt.Println("Failed to load files or find start address, exiting")
		flag.PrintDefaults()
		return
	}

	toy6502.SetPC(c, startaddr)

	prevPC := uint16(0x100)

	if quiet == false {
		fmt.Println(" ** Starting emulation **")
	}
	var instructions uint64
	for shouldexit == false {

		if disasm == true {
			mnem, _ := c.Disassemble(toy6502.GetPC(c))
			fmt.Printf("Addr: 0x%04X %s\n",
				toy6502.GetPC(c), mnem)
		}

		toy6502.ExecuteInstruction(c)
		instructions++

		if instructions > maxinst {
			shouldexit = true
			fmt.Println(" Max # of insts reached")
		}
		if toy6502.GetPC(c) == prevPC || toy6502.GetPC(c) == 0x0000 {
			shouldexit = true
			mnem, _ := c.Disassemble(toy6502.GetPC(c))
			fmt.Printf("CPU stuck on 0x%04X Inst: %s\n",
				toy6502.GetPC(c), mnem)
		}
		if toy6502.GetMemory(c)[toy6502.GetPC(c)] == 0x02 {
			shouldexit = true
			if quiet == false {
				fmt.Printf("Exit on JAM instruction at PC 0x%04X.\n",
					toy6502.GetPC(c))
			}
		}
		if toy6502.GetMemory(c)[toy6502.GetPC(c)] == 0x00 {
			shouldexit = true
			if quiet == false {
				fmt.Printf("Exit on BRK instruction at PC 0x%04X.\n",
					toy6502.GetPC(c))
			}
		}
		prevPC = toy6502.GetPC(c)

	}
	if quiet == false {
		fmt.Printf("instructions run: %v cycles: %v\n",
			instructions, toy6502.GetCycles(c))
		fmt.Printf("Reading screen memory, ")
	}
	cksumScreen(c)
	if printscreen == true {
		var mem []byte
		var petscii byte
		mem = toy6502.GetMemory(c)
		fmt.Println("+----------------------------------------+")
		for y := 0; y < 25; y++ {
			fmt.Printf("|")
			for x := 0; x < 40; x++ {
				petscii = mem[0x0400+x+(y*40)]
				fmt.Printf("%s", string(c64format(petscii)))
			}
			fmt.Printf("|\n")
		}
		fmt.Println("+----------------------------------------+")
	}
	if printscreenhex == true {
		var mem []byte
		mem = toy6502.GetMemory(c)
		fmt.Println("+--------------------+")
		for y := 0; y < 25; y++ {
			fmt.Printf("|")
			for x := 0; x < 40; x++ {
				fmt.Printf("%02x ", mem[0x0400+x+(y*40)])
			}
			fmt.Printf("|\n")
		}
		fmt.Println("+--------------------+")
	}
	if showstack == true {
		var stack []byte
		stack = toy6502.GetMemory(c)
		fmt.Println("+-- STACK CONTENTS --+")
		fmt.Println("+--------------------+")
		for y := 0; y < 16; y++ {
			fmt.Printf("|")
			for x := 0; x < 16; x++ {
				fmt.Printf("%02x ", stack[0x0100+x+(y*16)])
			}
			fmt.Printf("|\n")
		}
		fmt.Println("+--------------------+")
	}
}
