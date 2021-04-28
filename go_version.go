package vers

import (
	"fmt"
	"runtime"
)

func Vers() {
	fmt.Println(runtime.Version())
}

func Cpu()  {

	cpu := runtime.NumCPU()

	fmt.Println("Number of CPU is ", cpu)

}