package collector

import (
	"fmt"
	"github.com/xixiwang12138/hlog/decode"
	"io/fs"
	"os"
	"time"
)

type FileConf struct {
	Folder     string
	FilePrefix string
}

type FileWriter struct {
	cf *FileConf
	fd *os.File
}

func NewFileWriter(cf *FileConf) *FileWriter {
	// 打开文件
	fw := &FileWriter{cf: cf}
	var err error
	_, err = os.Stat(cf.Folder)
	useFile := cf.Folder + "/" + cf.FilePrefix + "_" + Today()
	if err != nil {
		// Not Found Folder
		if os.IsNotExist(err) {
			err := os.Mkdir(cf.Folder, 0700)
			if err != nil {
				panic(err)
			}
			fw.fd, err = os.Create(useFile)
			err = fw.fd.Chmod(fs.ModeAppend)
			if err != nil {
				panic(err)
			}
			return fw
		} else {
			panic(err)
		}
	}
	// Exist
	if _, err = os.Stat(useFile); err == nil {
		fw.fd, err = os.OpenFile(useFile, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		return fw
	} else if os.IsNotExist(err) {
		fw.fd, err = os.Create(useFile)
		err = fw.fd.Chmod(fs.ModeAppend)
		if err != nil {
			panic(err)
		}
		return fw
	} else {
		panic(err)
	}
	return fw
}

func (f FileWriter) Output(l *decode.Log) {
	f.fd.Write(l.ToByteArray())
}

func (f FileWriter) Close() {
	f.fd.Close()
}

func Today() string {
	n := time.Now()
	return fmt.Sprintf("%d-%d-%d", n.Year(), n.Month(), n.Day())
}
