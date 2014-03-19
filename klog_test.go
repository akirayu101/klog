package klog

import (
	"os"
	"testing"
)

var l *Logger = newLogger()

func TestNewFileLogger(t *testing.T) {
	filename := "testdata/tmpfile"
	err := l.SetLoggerBackend(TFile, filename)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		os.Remove(filename)
	}()
	l.Info("hello")
	fd, err := os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	stat, _ := fd.Stat()
	if stat.Size() <= 0 {
		t.Error("expect write into something, but klog write nothing")
	}
}

func TestAll(t *testing.T) {
	l.SetLoggerBackend(TStdout, "")
	l.SetLevel(LDebug)
	l.Debug("this is debug")
	l.Info("this is info")
	l.Warn("this is warn")
	l.Error("this is error")
	//K.Fatal("msg:fatal")
}

func TestNoColor(t *testing.T) {
	flags := l.Flags()
	l.SetFlags(flags & ^Fcolor)
	l.Info("this info msg has no color")
}

func TestSetLevel(t *testing.T) {
	l.SetLevel(LInfo)
	l.Debug("dddd")
	l.Info("iiii")
}
