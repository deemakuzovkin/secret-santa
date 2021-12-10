package bitcask

import (
	"fmt"
	"github.com/prologic/bitcask"
	"os"
)

var (
	MAX_PART_CASH_SIZE = 1024 * 1024 * 100
)

type Data struct {
	db *bitcask.Bitcask
}

//Connect init data
func Connect(path string) (*Data, error) {
	tmpPath := fmt.Sprintf("%s/%s", os.TempDir(), path)
	base, err := bitcask.Open(tmpPath, bitcask.WithMaxDatafileSize(MAX_PART_CASH_SIZE), bitcask.WithAutoRecovery(true))
	if err != nil {
		return nil, err
	}
	return &Data{
		db: base,
	}, nil
}

//Add add value by key
func (data *Data) Add(key, value []byte) error {
	return data.db.Put(key, value)
}

//Get get value by key
func (data *Data) Get(key []byte) ([]byte, error) {
	return data.db.Get(key)
}