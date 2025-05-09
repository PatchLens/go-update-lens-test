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

func TestOpen(t *testing.T) {
	t.Parallel()

	tmpFile, _ := makeTestMap(t)
	defer os.Remove(tmpFile)
}

func TestOpenCommit(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	require.NoError(t, m.Commit())
}

func TestCSV(t *testing.T) {
	t.Run("set", func(t *testing.T) {
		t.Parallel()

		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		require.NoError(t, SetCSV(m, "foo", "bar"))
	})
	t.Run("get", func(t *testing.T) {
		t.Parallel()

		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		_, ok := GetStrCSV(m, "foo")
		require.False(t, ok)
	})
	t.Run("size", func(t *testing.T) {
		t.Parallel()

		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		require.Equal(t, 0, SizeCSV(m))
	})
	t.Run("commit", func(t *testing.T) {
		t.Parallel()

		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		require.NoError(t, CommitCSV(m))
	})
	t.Run("set_get_zero_int", func(t *testing.T) {
		t.Parallel()

		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		err := SetCSV(m, "zero", 0)
		require.NoError(t, err)
		val, ok := GetIntCSV(m, "zero")
		assert.True(t, ok)
		assert.Equal(t, 0, val)
	})
	t.Run("set_get_zero_str", func(t *testing.T) {
		t.Parallel()

		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		err := SetCSV(m, "zero", "")
		require.NoError(t, err)
		val, ok := GetStrCSV(m, "zero")
		assert.True(t, ok)
		assert.Equal(t, "", val)
	})
	t.Run("set_get_zero_map", func(t *testing.T) {
		t.Parallel()

		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		zMap := map[string]int{"z": 0, "": 0}
		err := SetCSV(m, "zero", zMap)
		require.NoError(t, err)
		var val map[string]int
		ok := GetAnyCSV(m, "zero", &val)
		assert.True(t, ok)
		assert.Equal(t, zMap, val)
	})
	t.Run("set_get_zero_field_struct", func(t *testing.T) {
		t.Parallel()

		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		zeroStruct := TestStruct{
			Map:       map[string]TestStruct{"foo": {}},
			MapIntKey: map[int]string{0: "foo", 1: "1", 2: ""},
		}
		filledStruct := TestStruct{
			Value: "foo",
			ID:    1,
			Float: 1.0,
			Bool:  true,
			Bytes: []byte("foo"),
		}
		require.NoError(t, SetCSV(m, "zero", zeroStruct))
		require.NoError(t, SetCSV(m, "filled", filledStruct))
		var val TestStruct
		assert.True(t, GetAnyCSV(m, "filled", &val))
		assert.Equal(t, filledStruct, val)
		val = TestStruct{} // reset due to https://github.com/go-analyze/flat-file-map/issues/18
		assert.True(t, GetAnyCSV(m, "zero", &val))
		assert.Equal(t, zeroStruct, val)
	})
}

func TestUpdateFail(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	require.NoError(t, SetCSV(m, "zero", TestStruct{
		Map:       map[string]TestStruct{"foo": {}},
		MapIntKey: map[int]string{0: "foo", 1: "1", 2: ""},
	}))
	require.NoError(t, m.Commit())

	stat, err := os.Stat(tmpFile)
	require.NoError(t, err)
	assert.Equal(t, int64(254), stat.Size())
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

func TestOperateCSV(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	OperateCSV(m)
}

func TestOperateWriteCSV(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	OperateWriteCSV(m)
}

func TestOperateReadCSV(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	OperateReadCSV(m)
}

func TestInterface(t *testing.T) {
	t.Parallel()

	t.Run("set", func(t *testing.T) {
		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		require.NoError(t, SetInterface(m, "foo", "bar"))
	})
	t.Run("get", func(t *testing.T) {
		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		_, ok := GetInterface(m, "foo")
		require.False(t, ok)
	})
	t.Run("size", func(t *testing.T) {
		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		require.Equal(t, 0, SizeInterface(m))
	})
	t.Run("commit", func(t *testing.T) {
		tmpFile, m := makeTestMap(t)
		defer os.Remove(tmpFile)

		require.NoError(t, CommitInterface(m))
	})
}

func TestOperateInterface(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	OperateInterface(m)
}

func TestOperateWriteInterface(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	OperateWriteInterface(m)
}

func TestOperateReadInterface(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	OperateReadInterface(m)
}

// Tests for unique function structures

func TestContainsSetCSV(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	// isolated test due to unique double call behavior
	contains, err := ContainsSetCSV(m, "zero", TestStruct{
		Map:       map[string]TestStruct{"foo": {}},
		MapIntKey: map[int]string{0: "foo", 1: "1", 2: ""},
	})
	assert.False(t, contains)
	require.NoError(t, err)
}

func TestRecursiveCSVSet(t *testing.T) {
	t.Parallel()

	tmpFile, m := makeTestMap(t)
	defer os.Remove(tmpFile)

	// isolated test due to unique double call behavior
	err := RecursiveCSVSet(m, "zero", TestStruct{
		Map:       map[string]TestStruct{"foo": {}},
		MapIntKey: map[int]string{0: "foo", 1: "1", 2: ""},
	}, 10)
	require.NoError(t, err)
}
