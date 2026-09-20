package main

import(
	"fmt"
	"time"
)

func main(){
	curent_time:=time.Now()
	cpu_load:=0
	fmt.Printf("Time is - %s", curent_time.Format("15.03.2005 15:01:05"))
	fmt.Printf("\n CPU load is %d \n",cpu_load)
}