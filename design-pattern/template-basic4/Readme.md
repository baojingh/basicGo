首先，让我们回顾一下 `NewBankExecutor` 函数的定义：
```go
func NewBankExecutor(bankHandler BankHandler) *BankExecutor {
	executor := &BankExecutor{handler: bankHandler}
	return executor
}
```
这个函数接收一个类型为 `BankHandler` 的参数 `bankHandler`，并返回一个指向 `BankExecutor` 结构体的指针。
由于 `BankHandler` 是一个接口，因此任何实现了该接口的类型都可以作为参数传递给 `NewBankExecutor` 函数。


接下来，我们看 `DepositExecutor` 结构体：
```go
type DepositExecutor struct {
	*DefaultExecutor
	userVip bool
}
// ... 其他方法 ...
```

`DepositExecutor` 结构体中嵌入了一个 `*DefaultExecutor` 的指针，并且它自己也实现了 `BankHandler` 接口中的 `HandleConcreteBusiness` 
和 `CheckVipIdentity` 方法（同时也重写了 `Commentate` 方法）。

DepositExecutor 还实现了 BankHandler 接口中未被 DefaultExecutor 覆盖的方法 HandleConcreteBusiness 和 CheckVipIdentity。
由于 DepositExecutor 提供了 BankHandler 接口所需的所有方法，因此它满足该接口的要求，
可以被视为该接口的一个实现。当 DepositExecutor 结构体实例化时，它自动继承了 DefaultExecutor 中为 BankHandler 接口定义的所有方法（除了被重写的方法）。


由于 `DefaultExecutor` 已经为 `BankHandler` 接口中的其他方法提供了默认实现，而 `DepositExecutor` 又覆盖了部分方法，
因此 `DepositExecutor` 结构体类型完整地实现了 `BankHandler` 接口。

在 Go 中，如果一个类型实现了接口中定义的所有方法（包括接口中嵌入的其他接口的方法），那么这个类型就被认为实现了该接口。在 Go 中，接口是隐式实现的.
由于 `DepositExecutor` 实现了 `BankHandler` 接口，因此它可以被赋值给 `BankHandler` 类型的变量，并作为参数传递给 `NewBankExecutor` 函数。

因为 `DepositExecutor` 结构体类型实现了 `BankHandler` 接口，所以 `NewBankExecutor` 函数可以接收 `DepositExecutor` 类型的变量作为参数。