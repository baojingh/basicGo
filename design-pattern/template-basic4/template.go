package template_basic4

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
*

	Author     : He Bao Jing
	Date       : 5/10/2024 2:59 PM
	Description:
*/

//======================== interface ====================================================

type BankHandler interface {
	// 排队拿号
	TakeRowNumber()
	// 等位
	WaitInHead()
	// 处理具体业务
	HandleConcreteBusiness()
	// 对服务作出评价
	Commentate()
	// 钩子方法，判断是不是VIP， VIP不用等位
	CheckVipIdentity() bool
}

type BankExecutor struct {
	handler BankHandler
}

// 模板方法，处理银行业务
func (b *BankExecutor) ExecuteBankBusiness() {
	// 适用于与客户端单次交互的流程
	// 如果需要与客户端多次交互才能完成整个流程，每次交互的操作去对应模板里定义的方法就好，并不需要一个调用所有方法的模板方法
	b.handler.TakeRowNumber()
	if !b.handler.CheckVipIdentity() {
		b.handler.WaitInHead()
	}
	b.handler.HandleConcreteBusiness()
	b.handler.Commentate()
}
func NewBankExecutor(bankHandler BankHandler) *BankExecutor {
	executor := &BankExecutor{handler: bankHandler}
	return executor
}

// ======================= default implement =====================================================

type DefaultExecutor struct {
}

func (*DefaultExecutor) TakeRowNumber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延三个安排")
}

func (dbh *DefaultExecutor) WaitInHead() {
	fmt.Println("排队等号中...")
	time.Sleep(1 * time.Second)
	fmt.Println("请去窗口xxx...")
}

func (*DefaultExecutor) Commentate() {
	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

func (*DefaultExecutor) CheckVipIdentity() bool {
	// 留给具体实现类实现
	return false
}

// ============================= sub class =====================================================================================

type DepositExecutor struct {
	*DefaultExecutor
	userVip bool
}

func (dh *DepositExecutor) HandleConcreteBusiness() {
	fmt.Println("账户存储很多万人民币...")
}
func (dh *DepositExecutor) CheckVipIdentity() bool {
	return dh.userVip
}

func (*DepositExecutor) Commentate() {
	fmt.Println("this comment is override by deposit")
}
