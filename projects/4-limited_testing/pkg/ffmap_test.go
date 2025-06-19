package pkg

import (
	"os"
	"testing"

	"github.com/go-analyze/flat-file-map/ffmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeTestMap(t *testing.T) (string, *ffmap.KeyValueCSV) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "testm.*.csv")
	require.NoError(t, err)
	m, err := Open(tmpfile.Name())
	require.NoError(t, err)
	return tmpfile.Name(), m
}

type TestStruct struct {
	Value     string
	ID        int
	Float     float64
	Bool      bool
	Map       map[string]TestStruct
	MapIntKey map[int]string
	Bytes     []byte
}

func TestSetCommitStruct(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	require.NoError(t, SetCSV(m, "zero", TestStruct{
		Map:       map[string]TestStruct{"foo": {}},
		MapIntKey: map[int]string{0: "foo", 1: "1", 2: ""},
	}))
	require.NoError(t, CommitCSV(m))

	_, err := os.Stat(tmpFile)
	require.NoError(t, err)
}

func TestInterface(t *testing.T) {
	t.Parallel()

	t.Run("set_get", func(t *testing.T) {
		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		_, ok := GetInterface(m, "foo")
		assert.False(t, ok)

		require.NoError(t, SetInterface(m, "foo", "bar"))

		assert.Equal(t, 1, SizeCSV(m))

		_, ok = GetInterface(m, "foo")
		assert.True(t, ok)
	})
}
