package template_basic2

// https://blog.csdn.net/weixin_42117918/article/details/115466786

import (
	"fmt"
)

// =========================== abstract class ===============================================
type HandleInterface interface {
	// Handle this is a general action to call for outside
	Handle(uri string)
}

type ExecutorInterface interface {
	download()
	save()
}

type Template struct {
	ExecutorInterface
	uri string
}

func (t *Template) Handle(uri string) {
	fmt.Println("Prepare for download...", uri)
	t.ExecutorInterface.download()
	fmt.Println("Prepare for save...", uri)
	t.ExecutorInterface.save()
}

func (t *Template) save() {
	fmt.Println("default save...")
}

func newTemp(impl ExecutorInterface) *Template {
	t := &Template{
		uri:               "",
		ExecutorInterface: impl,
	}
	return t
}

// =========================== sub1 class ===============================================
type HttpHandler struct {
	*Template
}

func (hh *HttpHandler) download() {
	fmt.Println("http download...", hh.uri)
}

func (hh *HttpHandler) save() {
	fmt.Println("http save...")
}

// =========================== sub 2 class ===============================================
type FtpHandler struct {
	*Template
}

func (hh *FtpHandler) download() {
	fmt.Println("ftp download...", hh.uri)
}

func (hh *FtpHandler) save() {
	fmt.Println("ftp save...")
}
