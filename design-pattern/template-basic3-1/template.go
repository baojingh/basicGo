package template_basic3

import "fmt"

// =========================== abstract class ===============================================
type RepoManagerHandler interface {
	Delete(param string)
	List() string
}
type OperationsHandler interface {
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

type RepoExecutor struct {
	RepoManagerHandler
	OperationsHandler
}

func (t *RepoExecutor) Delete(param string) {
	t.OperationsHandler.PreHandleDel()
	t.OperationsHandler.HandleDel()
	t.OperationsHandler.PostHandleDel()
}

func (t *RepoExecutor) List() string {
	t.OperationsHandler.PreHandleList()
	t.OperationsHandler.HandleList()
	t.OperationsHandler.PostHandleList()
	return "list"
}

func NewRepoExecutor(operationHandler OperationsHandler) *RepoExecutor {
	return &RepoExecutor{
		OperationsHandler: operationHandler,
	}
}

// ================ default =========================================

type DefaultManager struct {
}

func (t *DefaultManager) PreHandleDel() {
	fmt.Println("default PreHandleDel")
}

func (t *DefaultManager) HandleDel() {
	fmt.Println("default HandleDel")
}

func (t *DefaultManager) PostHandleDel() {
	fmt.Println("default PostHandleDel")
}

func (t *DefaultManager) PreHandleList() {
	fmt.Println("default PreHandleList")
}

func (t *DefaultManager) HandleList() {
	fmt.Println("default HandleList")
}

func (t *DefaultManager) PostHandleList() {
	fmt.Println("default PostHandleList")
}

// =========================== sub1 class ===============================================
type PypiRepoManager struct {
	*DefaultManager
	Client
}

func (handle *PypiRepoManager) PreHandleDel() {
	fmt.Println("Pypi repo is PreHandleDel")
}

func (handle *PypiRepoManager) PreHandleList() {
	fmt.Println("Pypi repo is PreHandleList")
}

func NewPypiRepoManager() *RepoExecutor {
	manager := &PypiRepoManager{
		Client:         Client{},
		DefaultManager: &DefaultManager{},
	}
	executor := NewRepoExecutor(manager)
	return executor
}

// =========================== sub 2 class ===============================================
type DebRepoManager struct {
	*DefaultManager
	Client
}

func (handle *DebRepoManager) PreHandleDel() {
	fmt.Println("Deb repo is PreHandleDel")
}

func (handle *DebRepoManager) PreHandleList() {
	fmt.Println("Deb repo is PreHandleList")
}

func NewDebRepoManager() *RepoExecutor {
	manager := &DebRepoManager{
		Client:         Client{},
		DefaultManager: &DefaultManager{},
	}
	executor := NewRepoExecutor(manager)
	return executor
}
