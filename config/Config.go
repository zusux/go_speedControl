package config

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"path"
	"sync"
	"time"
)
type Idrac struct {
	Host string `ini:"host"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Mod string `ini:"mod"`
}
var idrac *Idrac

var once sync.Once

var Log *logrus.Logger

//初始化日志文件系统
func init()  {
	Log = logrus.New()
	ConfigLogger("./logs","log",time.Hour*24,time.Hour*24)
}
//NewIdrac 返回idrac 配置结构体
//配置文件只读取一次 如果读取异常则使用默认值 host:192.168.0.120 root calvin
func NewIdrac() *Idrac {

	idrac = &Idrac{
		Host:"192.168.0.120",
		Username:"root",
		Password:"calvin",
		Mod:"slow",
	}
	cfg, err := ini.Load("my.ini")
	if err != nil {
		Log.Error(fmt.Sprintf("Fail to read file: %v", err))
	}else{
		err = cfg.Section("idrac").MapTo(idrac)
		if err != nil{
			Log.Error(fmt.Sprintf("Fail to map my.ini %v",err))
		}
	}

	//once.Do(func() {
	//
	//})
	return idrac
}

func ConfigLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge), // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	},&logrus.JSONFormatter{})
	Log.AddHook(lfHook)
}
