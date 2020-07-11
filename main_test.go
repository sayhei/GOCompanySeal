package seal

import (
	"fmt"
	"testing"
)

func TestGenerateSeal(t *testing.T) {
	seal:=GenerateSeal("南京义途网络科技有限公司南京义途网络科技有限公司")
	fmt.Println(seal)
}