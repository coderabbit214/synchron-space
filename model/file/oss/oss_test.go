package oss

import (
	"testing"
)

func Test_getService(t *testing.T) {
	GetService()
}

func TestOssService_Upload(t *testing.T) {
	service := GetService()
	service.Upload("/Users/Mr_J/Downloads/【补充协议二】中金财富私享5777号FOF单一资产管理计划资产管理合同之补充协议二.docx", "aaa.docx")
}
