package shortmsg

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	asm := NewAliyunShortMsg()
	asm.Init("LTAI5t9h8KXfcmfrpFdmfkaF", "tWYPlgdjyVlp2iyRY3aRmDUxClhYhu", "cn-hangzhou", "https", "Hollow抒发树洞", "SMS_247945313")
	fmt.Println(asm.PostMsg("99999999999", "123456"))

}
