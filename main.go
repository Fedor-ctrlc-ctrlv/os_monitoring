package main

import(
	"fmt"
	"time"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"net/http"
)

func getCPUload() float64{
	percent,err:=cpu.Percent(time.Second,false )
	if err!=nil{
		return 0
	}
	return percent[0]
}

func monCpu(cpuChan chan <-string){
	for{
		load:= getCPUload()
		msg:= fmt.Sprintf(" CPU load is %.2f%% ",load)
		cpuChan<-msg
		time.Sleep(time.Second*5)
	}
}

func getRamusage() float64{
    ram,err:=mem.VirtualMemory()
	if err != nil{
		return 0
	}
	return ram.UsedPercent
}

func monRam(ramchan chan<-string){
	for{
		usageram:=getRamusage()
		msg:=fmt.Sprintf("Ram usage is %.2f%%", usageram)
		ramchan<-msg
		time.Sleep(time.Second *5)
	}
}

func metricsHandler(w http.ResponseWriter, r *http.Request){
	cpu:=getCPUload()
	ram:=getRamusage()
	fmt.Fprintf(w, "# HELP os_monitor_cpu_load Current CPU load percentage\n")
    fmt.Fprintf(w, "# TYPE os_monitor_cpu_load gauge\n")
    fmt.Fprintf(w, "os_monitor_cpu_load %.2f\n", cpu)

    fmt.Fprintf(w, "# HELP os_monitor_ram_usage Current RAM usage percentage\n")
    fmt.Fprintf(w, "# TYPE os_monitor_ram_usage gauge\n")
    fmt.Fprintf(w, "os_monitor_ram_usage %.2f\n", ram)
}


func main(){
	cpuChan:=make(chan string)
	ramchan:=make(chan string)
	go monCpu(cpuChan)
	go monRam(ramchan)

	http.HandleFunc("/metrics", metricsHandler)
	go func(){
		if err:=http.ListenAndServe(":8080",nil);err!=nil{
			fmt.Printf("Error starting server: %s\n", err)
		}

	}()


	for{
		select{
		case msg:=<-cpuChan:
			curent_time:=time.Now().Format("15.03.2006 15:01:05")
			fmt.Printf("[%s] %s \n", curent_time,msg)
		case msg:=<-ramchan:
			curent_time:=time.Now().Format("15.03.2006 15:01:05")
			fmt.Printf("[%s] %s \n", curent_time,msg)
		}
	}
}