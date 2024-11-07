package storage

import "github.com/martbul/locNot/types"

type Storage interface {
	Get(int) *types.User
}
