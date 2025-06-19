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

func ContainsCSV(kv *ffmap.KeyValueCSV, key string) bool {
	return kv.ContainsKey(key)
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

func ContainsSetCSV(kv *ffmap.KeyValueCSV, key string, value any) (bool, error) {
	contains := ContainsCSV(kv, key)
	return contains, kv.Set(key, value)
}

func RecursiveCSVSet(kv *ffmap.KeyValueCSV, key string, value any, count int) error {
	if count > 0 {
		return RecursiveCSVSet(kv, key, value, count-1)
	}
	return kv.Set(key, value)
}
