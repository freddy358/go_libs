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

	fmt.Println(cpu)

}