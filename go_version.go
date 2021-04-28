package vers

import (
	"fmt"
	"runtime"
)

func Vers() {
	fmt.Println(runtime.Version())
}

func Cpu()  {

	fmt.Println(runtime.NumCPU())

}