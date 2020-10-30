package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gopkg.in/sohlich/elogrus.v7"
	_ "gopkg.in/sohlich/elogrus.v7"
	"path"
	"runtime"
	"strconv"
	"study_gin/common/library/logger"
	"time"
)

// 日志记录到文件
func HandlerLoggerToFile() gin.HandlerFunc {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, filename := path.Split(f.File)
			return f.Function, filename + "/line:" + strconv.Itoa(f.Line)
		},
		ForceColors: true,
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.AddHook(logger.NewFileHook())

	// 是否转发到es
	// 如果es节点异常,会影响api接口性能.有延迟.todo 后面会兼容节点异常
	esMap := viper.GetStringMap("es")
	if _, ok := esMap["open"]; ok {
		if esMap["open"] == true {
			// 转发到es
			client, err := elastic.NewClient(elastic.SetSniff(false),
				elastic.SetURL(cast.ToString(esMap["hostport"])))
			if err != nil {
				logrus.Fatalf("启动连接es失败:[%v]", err)
			}
			// es hook
			hookEs, err := elogrus.NewElasticHook(client, cast.ToString(esMap["localhost"]), logrus.DebugLevel, cast.ToString(esMap["index"]))
			if err != nil {
				logrus.Fatal("es hook启动失败:[%v]", err)
			}
			logrus.AddHook(hookEs)
		}
	}

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logrus.Infof("| %v | %s | %s | %s |",
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
