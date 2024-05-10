package template_basic3

import "fmt"

// =========================== abstract class ===============================================
type RepoManageInterface interface {
	Delete(param string)
	List() string
}
type OperationsInterface interface {
	PreHandleDel()
	HandleDel()
	PostHandleDel()

	PreHandleList()
	HandleList()
	PostHandleList()
}

type Client struct {
	RepoUrl    string `json:"repoUrl"`
	RepoName   string `json:"repoName"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	ConfigPath string `json:"configPath"`
}

type Template struct {
	OperationsInterface
	RepoManageInterface
}

func (t *Template) Delete(param string) {
	t.OperationsInterface.PreHandleDel()
	t.OperationsInterface.HandleDel()
	t.OperationsInterface.PostHandleDel()
}

func (t *Template) List() string {
	t.OperationsInterface.PreHandleList()
	t.OperationsInterface.HandleList()
	t.OperationsInterface.PostHandleList()
	return "list"
}

func (t *Template) PreHandleDel() {
	fmt.Println("default PreHandleDel")
}

func (t *Template) HandleDel() {
	fmt.Println("default HandleDel")
}

func (t *Template) PostHandleDel() {
	fmt.Println("default PostHandleDel")
}

func (t *Template) PreHandleList() {
	fmt.Println("default PreHandleList")
}

func (t *Template) HandleList() {
	fmt.Println("default HandleList")
}

func (t *Template) PostHandleList() {
	fmt.Println("default PostHandleList")
}

func newTemplate(impl OperationsInterface) *Template {
	return &Template{
		OperationsInterface: impl,
	}
}

// =========================== sub1 class ===============================================
type PypiRepoManager struct {
	*Template
	Client
}

func (handle *PypiRepoManager) PreHandleDel() {
	handle.Template.PreHandleDel()
	//fmt.Println("Pypi repo is PreHandleDel")
}

func (handle *PypiRepoManager) HandleDel() {
	fmt.Println("Pypi repo is HandleDel")
}

func (handle *PypiRepoManager) PreHandleList() {
	fmt.Println("Pypi repo is PreHandleList")
}

func (handle *PypiRepoManager) HandleList() {
	fmt.Println("Pypi repo is HandleList")
}

func NewPypiRepoManager() *PypiRepoManager {
	manager := &PypiRepoManager{
		Client:   Client{},
		Template: &Template{},
	}
	template := newTemplate(manager)
	manager.Template = template
	return manager
}

// =========================== sub 2 class ===============================================
type DebRepoManager struct {
	*Template
	Client
}

func (handle *DebRepoManager) PreHandleDel() {
	fmt.Println("Deb repo is PreHandleDel")
}

func (handle *DebRepoManager) HandleDel() {
	fmt.Println("Deb repo is HandleDel")
}

func (handle *DebRepoManager) PreHandleList() {
	fmt.Println("Deb repo is PreHandleList")
}

//func (handle *DebRepoManager) HandleList() {
//	fmt.Println("Deb repo is HandleList")
//}

func NewDebRepoManager() *DebRepoManager {
	manager := &DebRepoManager{
		Client:   Client{},
		Template: &Template{},
	}
	template := newTemplate(manager)
	manager.Template = template
	return manager
}
