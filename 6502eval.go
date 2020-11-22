package main

import (
	"flag"
	"fmt"
	"github.com/cespare/xxhash"
	"github.com/icepic/toy6502"
	"os"
)

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

func c64load(fname string, c *toy6502.CPU) uint16 {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	fi, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if fi.Size() < 3 {
		fmt.Println("Invalid small size")
		return 0
	}
	var addrtemp = make([]byte, 2)
	_, err = f.Read(addrtemp)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var newsize = fi.Size() - 2
	var addr int64 = int64(addrtemp[0]) + int64(addrtemp[1])*256
	//	fmt.Printf("C64 Load Addr: 0x%04x, end: 0x%04x\n", addr, addr+newsize)

	if (fi.Size() + addr - 2) > int64(len(toy6502.GetMemory(c))) {
		fmt.Printf(
			"Illegal size %04x, read would be beyond end of memory.\n",
			newsize)
		return 0
	}

	var temp = make([]byte, newsize)
	_, err = f.Read(temp)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	copy(toy6502.GetMemory(c)[addr:addr+newsize], temp[:])
	return c64findsys(addr, newsize, c)
}

func Loadfileat(fname string, c *toy6502.CPU, addr int64) bool {

	f, err := os.Open(fname)
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
	return true
}

func main() {

	var quiet bool
	c := toy6502.New()

	myFilename := flag.String("filename", "a.out", "c64.prg to load and run.")
	flag.BoolVar(&quiet, "q", false, "Quiet mode, only print the checksum.")
	flag.Parse()

	toy6502.SetPC(c, c64load(*myFilename, c))

	prevPC := uint16(0x00)

	if quiet == false {
		fmt.Println(" ** Starting emulation **")
	}
	var instructions uint64
	for {
		toy6502.ExecuteInstruction(c)
		instructions++
		if toy6502.GetPC(c) == prevPC {
			if toy6502.GetMemory(c)[toy6502.GetPC(c)] == 0x02 {
				if quiet == false {
					fmt.Printf("Exit on JAM instruction at PC 0x%04X.\n",
						toy6502.GetPC(c))
				}
			}
			if quiet == false {
				fmt.Printf("instructions run: %v cycles: %v\n",
					instructions,
					toy6502.GetCycles(c))
				fmt.Printf("Reading screen memory, ")
			}
			cksumScreen(c)
			return
		}

		prevPC = toy6502.GetPC(c)
	}
}
