package template_basic2

import "testing"

/**
  Author     : He Bao Jing
  Date       : 5/10/2024 9:27 AM
  Description:
*/

// =========================== Test class ===============================================
func TestHello(t *testing.T) {
	dl := &FtpHandler{}
	template := newTemp(dl)
	dl.Template = template
	dl.Handle("baidu.com")

	dl2 := &HttpHandler{}
	template2 := newTemp(dl2)
	dl2.Template = template2
	dl2.Handle("google.com")
}
