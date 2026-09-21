package main

import(
	"fmt"
	"time"
	"os"
	"strconv"
	"strings"
	"bufio"
)

func getCPUload() float64{
	file,err:=os.Open("/proc/stat")
	if err!=nil{
		return 0
	}
	defer file.Close()

	scanner:= bufio.NewScanner(file)
	scanner.Scan()
	line:=scanner.Text()
	fields:= strings.Fields(line)
	if len(fields)<5{
		return 0
	}

	var total,idle uint64
	for i:=1 ;i<len(fields);i++{
		val,_:= strconv.ParseUint(fields[i],10,64)
		total+=val
		if i==4{
			idle=val
		}
	}
	if total == 0{
		return 0
	}
	return float64(total - idle)/float64(total)*100
}


func main(){
	curent_time:=time.Now()
	cpu_load:=getCPUload()
	fmt.Printf("Time is - %s", curent_time.Format("15.03.2006 15:01:05"))
	fmt.Printf("\n CPU load is %.2f%% \n",cpu_load)
}