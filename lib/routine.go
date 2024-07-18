package lib

import (
	"fmt"
	"runtime"
	"time"
)

func  print (till int, message string){
	for i := 0; i < till	; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println((i +1), message)
	}
}


func Rutin() {
	runtime.GOMAXPROCS(2)
	// cpus := runtime.NumCPU()

	go print(10, "halo")
	print(10, "apakabar")

	// println(cpus)

	var input string
	fmt.Scanln(&input)
}