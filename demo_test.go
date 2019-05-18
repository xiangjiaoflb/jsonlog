package jsonlog

import (
	"testing"
)

func TestNew(t *testing.T) {
	jslog := NewJSONLog("cslog", "conf")

	Warn(jslog).Str("cs", "cccc").Msg("")

}
