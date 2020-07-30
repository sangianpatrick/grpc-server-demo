package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/sangianpatrick/grpc-service-demo/pkg/util"
	"github.com/sangianpatrick/grpc-service-demo/src/service"

	_ "github.com/joho/godotenv/autoload" // for development
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	mariadbHost := os.Getenv("MARIADB_HOST")
	mariadbPort := os.Getenv("MARIADB_PORT")
	mariadbUser := os.Getenv("MARIADB_USER")
	mariadbPassword := os.Getenv("MARIADB_PASSWORD")
	mariadbName := os.Getenv("MARIADB_DATABASE")

	// set logger with apm hook
	logger := logrus.New()
	logger.SetFormatter(getLogFormatter())
	logger.SetReportCaller(true)
	logger.AddHook(&apmlogrus.Hook{
		LogLevels: logrus.AllLevels,
	})

	// set mariadb
	mariadbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mariadbUser, mariadbPassword, mariadbHost, mariadbPort, mariadbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", mariadbConnectionString, val.Encode())
	mariadb, err := apmsql.Open("mysql", dsn)
	if err != nil {
		logger.Fatal(err)
	}
	err = mariadb.Ping()
	if err != nil {
		logger.Fatal(err)
	}
	mariadb.SetMaxOpenConns(100)
	mariadb.SetMaxIdleConns(5)
	mariadb.SetConnMaxLifetime(time.Minute * 5)

	server := grpc.NewServer()

	// initiate service
	service.InitializeUser(logger, server, mariadb)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Serve(listener))
}

func getLogFormatter() *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		TimestampFormat: util.MySQLFormat,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	}
}
