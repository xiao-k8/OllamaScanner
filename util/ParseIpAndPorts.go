package util

import (
	"fmt"
	"github.com/malfunkt/iprange"
	"net"
	"sort"
	"strconv"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

// Generate IP and Port map to return
func GenerateTask(ipList []net.IP, ports []int) ([]map[string]int, int) {
	tasks := []map[string]int{}
	for _, ip := range ipList {
		for _, port := range ports {
			ipPort := map[string]int{ip.String(): port}
			tasks = append(tasks, ipPort)
		}
	}
	return tasks, len(tasks)
}

// Analyze user input ip
func GetIpList(ips string) ([]net.IP, error) {
	addressRangeList, err := iprange.ParseList(ips)
	if err != nil {
		return nil, err
	}
	IpList := addressRangeList.Expand()
	return IpList, nil
}

// Analyze user input ports
func ParsePorts(portStr string) ([]int, error) {
	portMap := make(map[int]struct{})

	if portStr == "" {
		return nil, nil
	}

	parts := strings.Split(portStr, ",")
	for _, part := range parts {
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("无效的端口范围格式: %s", part)
			}

			// 分别解析起始和结束端口
			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return nil, fmt.Errorf("无效的起始端口: %s", rangeParts)
			}

			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return nil, fmt.Errorf("无效的结束端口: %s", rangeParts)
			}

			if start > end {
				return nil, fmt.Errorf("无效的端口范围: %d-%d", start, end)
			}

			for p := start; p <= end; p++ {
				if p < 1 || p > 65535 {
					return nil, fmt.Errorf("端口号超出范围 (1-65535): %d", p)
				}
				portMap[p] = struct{}{}
			}
		} else {
			port, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("无效的端口号: %s", part)
			}

			if port < 1 || port > 65535 {
				return nil, fmt.Errorf("端口号超出范围 (1-65535): %d", port)
			}
			portMap[port] = struct{}{}
		}
	}

	ports := make([]int, 0, len(portMap))
	for p := range portMap {
		ports = append(ports, p)
	}
	sort.Ints(ports)
	return ports, nil
}
