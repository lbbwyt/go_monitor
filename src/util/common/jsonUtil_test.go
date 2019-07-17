package common

import (
	"chief_operation/src/server/admin/handler"
	"fmt"
	"testing"
)

func TestObj2JsonStr(t *testing.T) {
	param := handler.Param{};
	param.AppName = "test"
	str := Obj2JsonStr(param)
	fmt.Println(str)
}

