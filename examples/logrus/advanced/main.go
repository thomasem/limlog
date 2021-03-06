package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jar-o/limlog"
	"github.com/sirupsen/logrus"
)

func main() {
	// Get an instance of the logger
	l := limlog.NewLimlogrusInstance()
	inst := l.L.GetLogger().(*logrus.Logger)

	// Change settings for the instance
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		inst.Out = file
	} else {
		l.Info("Failed to log to file, using default stderr")
	}
	inst.Formatter = &logrus.JSONFormatter{}
	inst.Level = logrus.DebugLevel

	// Use the logger as you normally would
	l.SetLimiter("limiter1", 4, 1*time.Second, 6)
	l.Info("You don't have to limit if you don't want.")
	l.Debug("It's true.")
	for i := 0; i <= 10000000; i++ {
		l.ErrorL("limiter1", fmt.Sprintf("%d", i))
		l.WarnL("limiter1", fmt.Sprintf("%d", i), logrus.Fields{"helo": "wrld"})
		l.TraceL("limiter1", fmt.Sprintf("%d", i)) // Won't log. At DebugLevel
		l.InfoL("limiter1", fmt.Sprintf("%d", i))
		l.DebugL("limiter1", fmt.Sprintf("%d", i))
		// l.Debug(i) // <--- This will spew every i
	}
}
