package vers

import (
	"fmt"
	"runtime"
)

func Vers() {
	fmt.Println(runtime.Version())
	fmt.Println(runtime.NumCPU())

}
