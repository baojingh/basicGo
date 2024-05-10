package template_basic3

import "testing"

/**
  Author     : He Bao Jing
  Date       : 5/10/2024 9:27 AM
  Description:
*/

// =========================== Test class ===============================================
func TestHello(t *testing.T) {
	debmanager := NewPypiRepoManager()
	debmanager.Delete("del")
	//debmanager.List()

	//pypimanager := NewDebRepoManager()
	//pypimanager.Delete("del")
	//pypimanager.List()

}
