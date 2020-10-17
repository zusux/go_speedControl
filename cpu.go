package main

import (
	"fmt"
	"speedControl/config"
	"strconv"
	"strings"
)

type cpuInfo struct{
	CpuNum int
	CpuCores int
}
var Cpu *cpuInfo

func init()  {
	Cpu = &cpuInfo{}
	Cpu.GetCpuNum()
	Cpu.GetCpuCores()
}

func (c *cpuInfo)GetCpuNum() int {
	cpuNumStr := execCommand("cat /proc/cpuinfo |grep \"physical id\" | sort | uniq | wc -l")
	cpuNumStr = strings.Replace(cpuNumStr,"\n","",-1)
	cpuNum,_ :=strconv.Atoi(cpuNumStr)
	c.CpuNum = cpuNum
	return cpuNum
}

func (c *cpuInfo)GetCpuCores() int {
	cpuCoresStr := execCommand("cat /proc/cpuinfo | grep \"cores\" | uniq|awk '{print $4}'")
	cpuCoresStr = strings.Replace(cpuCoresStr,"\n","",-1)
	cpuCores,_ :=strconv.Atoi(cpuCoresStr)
	c.CpuCores = cpuCores
	return cpuCores
}

func getTemperature(pos int) float64  {
	cmdStr := fmt.Sprintf("sensors | sed -n '%dp' |awk '{print $4}'|sed 's/\\+//g'|sed 's/°C//g'",pos)
	t := execCommand(cmdStr)
	t = strings.Replace(t,"\n","",-1)
	config.Log.Info("温度为"+t)
	float, _ := strconv.ParseFloat(t, 64)
	return float
}



