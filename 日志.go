package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

type CustomLogger struct {
	writer    io.Writer
	logLevel  LogLevel
	timeStamp bool
}

func NewCustomLogger(writer io.Writer, logLevel LogLevel, timeStamp bool) *CustomLogger {
	return &CustomLogger{writer: writer, logLevel: logLevel, timeStamp: timeStamp}
}

func (l *CustomLogger) Write(p []byte) (n int, err error) {
	logMsg := string(p)

	if l.timeStamp {
		logMsg = fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), logMsg)
	}

	if l.logLevel <= Info {
		_, err = l.writer.Write([]byte(logMsg))
		if err != nil {
			return 0, err
		}
	}

	return len(p), nil
}

func main() {
	var logFilePath string
	var logLevelStr string
	var customFormat bool

	flag.StringVar(&logFilePath, "log", "logfile.txt", "日志文件路径")
	flag.StringVar(&logLevelStr, "level", "info", "日志级别: debug, info, warning, error")
	flag.BoolVar(&customFormat, "customFormat", false, "启用自定义日志格式")

	flag.Parse()

	logLevel := parseLogLevel(logLevelStr)

	file, err := openLogFile(logFilePath)
	if err != nil {
		log.Fatalf("打开日志文件出错: %v", err)
	}
	defer file.Close()

	var multiWriter io.Writer
	if customFormat {
		multiWriter = io.MultiWriter(os.Stdout, NewCustomLogger(file, logLevel, true))
	} else {
		multiWriter = io.MultiWriter(os.Stdout, file)
	}

	log.SetOutput(multiWriter)
	log.SetFlags(0)

	fmt.Fprintln(multiWriter, "开始记录日志...")

	// 模拟用户操作并记录日志
	fmt.Fprintln(multiWriter, "用户登录")
	time.Sleep(2 * time.Second)
	fmt.Fprintln(multiWriter, "用户执行操作A")
	time.Sleep(1 * time.Second)
	fmt.Fprintln(multiWriter, "用户执行操作B")

	fmt.Println("日志记录完成.")

	// 文件同步工具
	sourceDir := "./source"
	destinationDir := "./destination"

	err = syncFiles(sourceDir, destinationDir)
	if err != nil {
		log.Fatalf("同步文件出错: %v", err)
	}
}

func openLogFile(logFilePath string) (*os.File, error) {
	dir := filepath.Dir(logFilePath)

	// 如果日志目录不存在，则创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("创建日志目录失败: %v", err)
		}
	}

	// 以追加模式打开日志文件，如果不存在则创建
	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("打开日志文件失败: %v", err)
	}

	return file, nil
}

func parseLogLevel(levelStr string) LogLevel {
	switch strings.ToLower(levelStr) {
	case "debug":
		return Debug
	case "info":
		return Info
	case "warning":
		return Warning
	case "error":
		return Error
	default:
		return Info
	}
}

func syncFiles(sourceDir, destinationDir string) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				relativePath, _ := filepath.Rel(sourceDir, path)
				destinationPath := filepath.Join(destinationDir, relativePath)

				if info.IsDir() {
					return os.MkdirAll(destinationPath, os.ModePerm)
				}

				sourceModTime := info.ModTime()
				destinationInfo, err := os.Stat(destinationPath)

				if os.IsNotExist(err) || destinationInfo.ModTime().Before(sourceModTime) {
					fmt.Printf("同步中: %s\n", relativePath)
					return copyFile(path, destinationPath)
				}

				return nil
			})

			if err != nil {
				return err
			}
		}
	}
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
