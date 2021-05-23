package jsonpath

import (
	"encoding/json"
	"github.com/coinbase/step/jsonpath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JSONPath_Parse_Path(t *testing.T) {
	out, err := jsonpath.ParsePathString("$")
	assert.NoError(t, err)
	assert.Equal(t, len(out), 0)
}

func Test_JSONPath_Parse_PathLong(t *testing.T) {
	out, err := jsonpath.ParsePathString("$.a.b.c")

	assert.NoError(t, err)

	assert.Equal(t, len(out), 3)

	assert.Equal(t, out[0], "a")
	assert.Equal(t, out[1], "b")
	assert.Equal(t, out[2], "c")
}

func Test_JSONPath_NewPath(t *testing.T) {
	path, err := jsonpath.NewPath("$.a.b.c")

	assert.NoError(t, err)
	t.Logf("%+v", path)
}

func Test_JSONPath_Parsing(t *testing.T) {
	raw := []byte(`"$.a.b.c"`)

	var pathstr jsonpath.Path
	err := json.Unmarshal(raw, &pathstr)

	assert.NoError(t, err)
	t.Logf("%+v", pathstr)
}
