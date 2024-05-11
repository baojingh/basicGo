package template_basic4

import "testing"

/**
  Author     : He Bao Jing
  Date       : 5/10/2024 3:08 PM
  Description:
*/

func TestHello(t *testing.T) {
	dh := &DepositExecutor{userVip: false}
	bbe := NewBankExecutor(dh)
	bbe.ExecuteBankBusiness()
}
