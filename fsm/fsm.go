package fsm

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/dgraph-io/badger/v3"
	"github.com/hashicorp/raft"
)

type Command struct {
	Op    string
	Key   string
	Value []byte
}

type Response struct {
	Error error
	Data  interface{}
}

type badgerFSM struct {
	db *badger.DB
}

func (b badgerFSM) get(key string) ([]byte, error) {
	var txn = b.db.NewTransaction(false)
	defer func() {
		_ = txn.Commit()
	}()

	item, err := txn.Get([]byte(key))
	if err != nil {
		return nil, err
	}

	var value = make([]byte, 0)
	err = item.Value(func(val []byte) error {
		value = append(value, val...)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (b badgerFSM) set(key string, value []byte) error {
	var txn = b.db.NewTransaction(true)
	var err = txn.Set([]byte(key), value)
	if err != nil {
		return err
	}
	return txn.Commit()
}

func (b badgerFSM) delete(key string) error {
	var keyByte = []byte(key)
	var txn = b.db.NewTransaction(true)
	var err = txn.Delete(keyByte)
	if err != nil {
		return err
	}
	return txn.Commit()
}

func (b badgerFSM) Apply(log *raft.Log) interface{} {
	switch log.Type {
	case raft.LogCommand:
		var cmd = &Command{}
		err := json.Unmarshal(log.Data, cmd)
		if err != nil {
			return nil
		}
		var op = strings.ToUpper(strings.TrimSpace(cmd.Op))
		switch op {
		case "SET":
			return &Response{
				Error: b.set(cmd.Key, cmd.Value),
				Data:  cmd.Value,
			}
		case "GET":
			data, err := b.get(cmd.Key)
			return &Response{
				Error: err,
				Data:  data,
			}
		case "DELETE":
			return &Response{
				Error: b.delete(cmd.Key),
				Data:  nil,
			}
		}
	}

	return nil
}

func (b badgerFSM) Snapshot() (raft.FSMSnapshot, error) {
	return newSnapshot(), nil
}

func (b badgerFSM) Restore(closer io.ReadCloser) error {
	defer func() {
		var err = closer.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var decoder = json.NewDecoder(closer)
	for decoder.More() {
		var data = &Command{}
		var err = decoder.Decode(data)
		if err != nil {
			return err
		}

		err = b.set(data.Key, data.Value)
		if err != nil {
			return err
		}
	}

	_, err := decoder.Token()
	if err != nil {
		return err
	}

	return nil
}

func New(db *badger.DB) raft.FSM {
	return badgerFSM{db: db}
}
