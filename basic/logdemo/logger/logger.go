package logger

/**
  @Author:      He Bao Jing
  @Date:        6/27/2023 9:32 AM
  @Description:
*/

import (
	"fmt"
	"io"
	"path/filepath"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
	logger "github.com/sirupsen/logrus"
)

var log = logger.New()

// init init logs
func init() {
	path := "./"

	writer, err := rotate.New(
		filepath.Join(path, fmt.Sprintf("test-%s.log", "%Y%m%d")),
		rotate.WithLinkName(filepath.Join(path, "test.log")),
		rotate.WithRotationCount(7),
		rotate.WithRotationTime(time.Hour*24),
	)
	if err == nil {
		log.SetOutput(io.MultiWriter(writer /*, os.Stdout*/))
	}
}

// New new Logger
func New() *logger.Logger {
	return log
}

// NewRotatedLogger new Logger
func NewRotatedLogger(path string, t int, a int) (io.WriteCloser, error) {
	return rotate.New(
		fmt.Sprintf("%s.%s", path, "%Y%m%d%H%M"),
		rotate.WithLinkName(path),
		rotate.WithMaxAge(time.Minute*10),
		rotate.WithRotationTime(time.Minute),
	)
}
