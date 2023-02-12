package _test

import (
	"fmt"
	"github.com/xixiwang12138/hlog/conf"
	"testing"
)

type HTTPReq struct {
	Name string
	Id   string
}

func (H *HTTPReq) GetRequestId() string {
	return H.Id
}

func (H *HTTPReq) GetUserFlag() string {
	return H.Name
}

func TestInsert(t *testing.T) {
	_ = &conf.MongoDBConfig{
		UserName: "admin-all",
		Password: "B8ZC7DaEONQknCv",
		Host:     "111.230.227.84",
		Port:     "27016",
		DataBase: "Homi",
	}
	req := &HTTPReq{
		Name: "openid-3",
		Id:   "a543b54ad",
	}
	fmt.Println(req)
}
