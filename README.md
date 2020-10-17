# go_speedControl
dell r720 风扇风速静音控制系统

系统:CentOS Linux release 7.8.2003 (Core)

ipmitool 设置风速驱动(服务器开机控制面板中设置)
######1.设置 ipmitool ip:192.168.0.120  账号:root(默认) 密码:calvin(默认)

######2.安装和检查驱动是否有用 
```
yum  -y install epel-release
yum  -y install ipmitool

ipmitool -I lanplus -H 192.168.0.120 -U root -P calvin chassis power status
Chassis Power is on
```

######3.输入下面的命令把风扇转速设置为手动的(重启系统需要再次执行)
```cassandraql
ipmitool -I lanplus -H 192.168.0.120 -U root -P calvin raw 0x30 0x30 0x01 0x00
```

######4.安装linux温度监控软件 
```cassandraql
yum install -y lm_sensors
```

######5.配置用户名密码
my.ini 文件可以配置ip,用户名,密码
