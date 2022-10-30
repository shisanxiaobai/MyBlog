package middleware

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// <------标准库打印日志设置，同时打印到控制台和文件中，适合开发时调试使用------>

func init() {
	//Getwd获得当前目录的根目录
	path, _ := os.Getwd()
	f, err := os.OpenFile(path+"/log/testlog.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		log.Panic("testlog日志文件创建失败")
	}
	//io包中MultiWriter，可以把文件流和控制台标准输出流整合到一个io.Writer上，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)
}

// <----------------用于gin的日志中间件------------------------>

func Logger() gin.HandlerFunc {
	//在log目录下创建一个永久的gin日志，此处只是个花式写法0.0随意就好
	dir, _ := os.Getwd()
	filePath := "log"
	fileName := path.Join(dir, filePath, "gin-api.log")
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		log.Panic("gin-api日志文件创建失败")
	}

	linkName := "latest_log.log"
	//实例化logrus
	logger := logrus.New()
	//输出路径
	logger.Out = f
	//log等级debug级别
	logger.SetLevel(logrus.DebugLevel)

	logPath := dir + "\\log\\"
	logWriter, _ := retalog.New(
		logPath+"%Y%m%d.log",                   //gin日志循环输出地址
		retalog.WithMaxAge(7*24*time.Hour),     //文件最大保存时间7天168h
		retalog.WithRotationTime(24*time.Hour), // 文件最大保存时间24h
		retalog.WithLinkName(linkName),         //为最新的一份日志文件建立软链接
	)

	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	//以text文本格式输出内容
	Hook := lfshook.NewHook(writerMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//钩子函数
	logger.AddHook(Hook)

	return func(c *gin.Context) {

		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		//结束时间，since返回从启动算起经过的时间
		stopTime := time.Since(startTime)
		//花费时间单位换算为ms
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		//主机名
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknow"
		}
		//http请求的响应码
		statusCode := c.Writer.Status()
		//客户端ip
		clientIP := c.ClientIP()
		//用户代理
		UserAgent := c.Request.UserAgent()
		//返回的字节数
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		//请求方法
		method := c.Request.Method
		//url请求路由
		path := c.Request.URL.Path

		//以下面的格式结构化日志目录，看起来更加美观
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIP,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"UserAgent": UserAgent,
			"method":    method,
			"path":      path,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
