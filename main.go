package main

import (
	"OllamaScaner/util"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	ipList := flag.String("ip", "127.0.0.1", "设置需要扫描的IP地址或IP段")
	portList := flag.String("p", "11434", "设置需要扫描的端口范围")
	maxConcurrency := flag.Int("t", 20, "设置扫描线程")
	showHelp := flag.Bool("h", false, "显示帮助信息")
	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	//Set task name
	TaskName := *ipList
	//Set concurrency and chan
	concurrencyCh := make(chan struct{}, *maxConcurrency)
	var wg sync.WaitGroup

	log, _, _ := util.WriteLog(TaskName)
	success, _, _ := util.WriteSuccess(TaskName)
	//Resolve IP and ports range
	ips, err := util.GetIpList(*ipList)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		os.Exit(0)
	}
	ports, err := util.ParsePorts(*portList)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		os.Exit(0)
	}

	//start
	for _, ip := range ips {
		for _, port := range ports {

			wg.Add(1)
			concurrencyCh <- struct{}{}
			go func(ip string, port int) {
				defer wg.Done()
				defer func() { <-concurrencyCh }()
				host := ip + ":" + strconv.Itoa(port)
				bytes, _ := util.CreateHttpClint(host, "")
				isOllama, err := util.CheckResposeBody(bytes)
				if err != nil {
					fmt.Printf("err:%v\n", err)
				}
				if isOllama {
					success(host + "   Ollama is running\n")
				}
				log(host + "   No Ollama\n")
			}(ip.String(), port)
		}
	}
	wg.Wait()
	close(concurrencyCh)
}
