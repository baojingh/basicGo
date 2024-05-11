package template_basic3

import (
	"fmt"
	"testing"
)

/**
  Author     : He Bao Jing
  Date       : 5/10/2024 9:27 AM
  Description:
*/

// =========================== Test class ===============================================
func TestHello(t *testing.T) {
	executor := NewPypiRepoManager()
	executor.Delete("del")
	executor.List()

	fmt.Println("==========================")
	debExecutor := NewDebRepoManager()
	debExecutor.Delete("del")
	debExecutor.List()

}
