package vers

import (
	"fmt"
	"runtime"
)

func Vers() {
	version := runtime.Version()

	fmt.Println("Go version is ", version)
}

func Cpu()  {

	cpu := runtime.NumCPU()

	fmt.Println("Number of CPU is ", cpu)

}