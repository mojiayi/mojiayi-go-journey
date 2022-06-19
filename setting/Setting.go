package setting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/go-eden/routine"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	cfg            *ini.File
	WebSetting     = &WebConfig{}
	logOutSetting  = &LogOutputConfig{}
	MyLogger       = &logrus.Logger{}
	MetadataLogger *logrus.Logger
	localTraceId   = routine.NewLocalStorage()
	mySQLSetting   = &MySQLConfig{}
	DB             *gorm.DB
)

func Setup() {
	var err error
	cfg, err = ini.Load("setting/my.ini")
	if err != nil {
		fmt.Println("failed while load setting file setting/my.ini,err: ", err)
	}

	mapToConfig("web", WebSetting)

	mapToConfig("log", logOutSetting)

	mapToConfig("mysql", mySQLSetting)

	setupLogOutput()

	setupMySQL()
}

func mapToConfig(section string, value interface{}) {
	err := cfg.Section(section).MapTo(value)
	if err != nil {
		fmt.Println("failed while cfg.MapTo "+section+",err: ", err)
	}
}

type WebConfig struct {
	Port        int
	ContextPath string
}

type MyLogFormatter struct {
}

type LogOutputConfig struct {
	Dir                string
	MyLogPattern       string
	MetadataLogPattern string
}

func (m *MyLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	var requestMetadata = make(map[string]interface{})
	for k, v := range entry.Data {
		requestMetadata[k] = v
	}
	str, _ := json.Marshal(requestMetadata)

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog = fmt.Sprintf("%s|%s|%s|%s|%s\n", timestamp, entry.Level, GetTraceId(), entry.Message, string(str))
	buffer.WriteString(newLog)
	return buffer.Bytes(), nil
}

func setupLogOutput() {
	// 打印请求中业务日志
	MyLogger = initLog(logOutSetting.Dir, "access.log")
	// 打印请求的元数据信息
	MetadataLogger = initLog(logOutSetting.Dir, "metadata.log")
}

func initLog(path string, filename string) *logrus.Logger {
	log := logrus.New()
	log.Formatter = &MyLogFormatter{}

	filepath := path + filename
	writer, err := rotatelogs.New(
		filepath+".%Y%m%d",
		rotatelogs.WithLinkName(filepath),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	if err != nil {
		fmt.Println("fail to open log file " + filepath)
	}

	log.SetOutput(writer)
	log.Level = logrus.InfoLevel

	return log
}

func PutTraceIdIntoLocalStorage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		GetTraceId()

		ctx.Next()
	}
}

func GetTraceId() string {
	var traceId = ""
	if localTraceId.Get() == nil {
		traceId = strings.ReplaceAll(uuid.New(), "-", "")
		localTraceId.Set(traceId)
	} else {
		traceId = localTraceId.Get().(string)
	}
	return traceId
}

func RemoveTraceId() {
	localTraceId.Del()
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

func setupMySQL() {
	var err error
	var dbUrl = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mySQLSetting.User,
		mySQLSetting.Password,
		mySQLSetting.IP,
		mySQLSetting.Port,
		mySQLSetting.Database)
	DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		MyLogger.Error("models setup err:", err)
	}
}
