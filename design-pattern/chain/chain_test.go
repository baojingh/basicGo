package chain

import (
	"fmt"
	"testing"
)

/**
  @Author:      He Bao Jing
  @Date:        6/2/2023 5:56 PM
  @Description:
*/

func TestStartChain(t *testing.T) {
	chain := &SensitiveWordFilterChain{}

	political := &PoliticalSensitiveFilter{}
	ad := &AdSensitiveFilter{}
	force := &ForceSensitiveFilter{}

	chain.AddFilter(political)
	chain.AddFilter(ad)
	chain.AddFilter(force)

	isOK := chain.StartFilter("this is hello")
	fmt.Printf("Sensitive words? %t", isOK)
}
