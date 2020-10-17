package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sort"
	"speedControl/config"
	"strconv"
	"strings"
	"time"
)



func main()  {
	//读取配置
	idrac := config.NewIdrac()
	//画出趋势图
	var xAsix , yAsix []float64
	xAsix = make([]float64,0,100)
	yAsix = make([]float64,0,100)
	for i :=0.0;i<=100 ;i++  {
		var p float64
		switch idrac.Mod {
		 case "slow":
		 	p = Mod1(i)
		case "faster":
			p = Mod2(i)
		default:
			p = Mod1(i)
		}
		xAsix = append(xAsix,i)
		yAsix = append(yAsix,p)
	}
	image(xAsix,yAsix)
	//return
	//循环监控温度
	for{
		idrac = config.NewIdrac()
		max := getMaxTemp()
		var p float64
		switch idrac.Mod {
		case "slow":
			p = Mod1(max)
		case "faster":
			p = Mod2(max)
		default:
			p = Mod1(max)
		}
		fmt.Println(p)
		c := getCurrent()
		percent := int(p)
		// 如果温度转换的百分比和数据库中的不一致 则改变转速
		if c != percent{
			h16 := hex(percent)
			cmdStr := fmt.Sprintf("ipmitool -I lanplus -H %s -U %s -P %s raw 0x30 0x30 0x02 0xff 0x%s",idrac.Host,idrac.Username,idrac.Password,h16)
			//执行改变转速
			//config.Log.Info(cmdStr)
			config.Log.Info("目前转速:"+strconv.Itoa(percent))
			res := execCommand(cmdStr)
			res = strings.Replace(res,"\n","",-1)
			if len(res) == 0{
				config.Log.Info("执行改变转速成功")
				//把当前转速写入数据库
				setCurrent(percent)
			}
		}
		time.Sleep(time.Second * 10)
	}
}



func getMaxTemp() float64 {
	temps := make([]float64,0,4)
	for i:=0;i<Cpu.CpuNum;i++{
		pos := i * (Cpu.CpuCores+1) + 3 * (i+1)
		t := getTemperature(pos)
		temps = append(temps,t)
	}
	sort.Float64s(temps)
	max := temps[len(temps)-1]
	return max
}

func hex(i int) string {
	h := fmt.Sprintf("%x",i)
	return h
}

func execCommand(execStr string) string {
	cmd := exec.Command("bash", "-c", execStr)
	//显示运行的命令
	//fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		config.Log.Error(err)
		return ""
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	var result string
	for {
		line, err2 := reader.ReadString('\n')
		result += line
		if err2 != nil || io.EOF == err2 {
			break
		}
	}
	cmd.Wait()
	return result
}
