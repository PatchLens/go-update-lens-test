package pkg

import (
	"github.com/go-analyze/flat-file-map/ffmap"
)

func Open(path string) (*ffmap.KeyValueCSV, error) {
	return ffmap.OpenCSV(path)
}

func SetCSV(kv *ffmap.KeyValueCSV, key string, value any) error {
	return kv.Set(key, value)
}

func SetInterface(kv ffmap.MutableFFMap, key, value string) error {
	return kv.Set(key, value)
}

func GetStrCSV(kv *ffmap.KeyValueCSV, key string) (string, bool) {
	var value string
	ok := GetAnyCSV(kv, key, &value)
	return value, ok
}

func GetIntCSV(kv *ffmap.KeyValueCSV, key string) (int, bool) {
	var value int
	ok := GetAnyCSV(kv, key, &value)
	return value, ok
}

func GetAnyCSV(kv *ffmap.KeyValueCSV, key string, value any) bool {
	ok, err := kv.Get(key, value)
	if err != nil {
		panic(err) // bad practice, just example
	}
	return ok
}

func GetInterface(kv ffmap.MutableFFMap, key string) (string, bool) {
	var value string
	ok, err := kv.Get(key, &value)
	if err != nil {
		panic(err) // bad practice, just example
	}
	return value, ok
}

func SizeCSV(kv *ffmap.KeyValueCSV) int {
	return kv.Size()
}

func SizeInterface(kv ffmap.MutableFFMap) int {
	return kv.Size()
}

func CommitCSV(kv *ffmap.KeyValueCSV) error {
	return kv.Commit()
}

func CommitInterface(kv ffmap.MutableFFMap) error {
	return kv.Commit()
}

func OperateCSV(kv *ffmap.KeyValueCSV) {
	OperateWriteCSV(kv)
	OperateReadCSV(kv)
	if err := CommitCSV(kv); err != nil {
		panic(err)
	}
}

func OperateWriteCSV(kv *ffmap.KeyValueCSV) {
	if err := SetCSV(kv, "foo", "bar"); err != nil {
		panic(err)
	}
}

func OperateReadCSV(kv *ffmap.KeyValueCSV) {
	_, _ = GetStrCSV(kv, "foo")
	_ = SizeCSV(kv)
}

func OperateInterface(kv ffmap.MutableFFMap) {
	OperateWriteInterface(kv)
	OperateReadInterface(kv)
	if err := CommitInterface(kv); err != nil {
		panic(err)
	}
}

func OperateWriteInterface(kv ffmap.MutableFFMap) {
	if err := SetInterface(kv, "foo:i", "bar"); err != nil {
		panic(err)
	}
}

func OperateReadInterface(kv ffmap.MutableFFMap) {
	_, _ = GetInterface(kv, "foo:i")
	_ = SizeInterface(kv)
}
