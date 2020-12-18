package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const jsonVal = `
{"a" : 1,
"b": 2,
"c":3}
`

type TestStruct struct {
	A int
	B int
	C int
}

func TestPrettyPrint(t *testing.T) {
	var val = TestStruct{}

	json.Unmarshal([]byte(jsonVal), &val)
	sval, err := PrettyJSON(val)
	assert.True(t, err == nil)
	t.Logf("Value: %s", sval)
}
