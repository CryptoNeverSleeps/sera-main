package model

import "github.com/seracoin/serad/domain/consensus/model/externalapi"

// BlockIterator is an iterator over blocks according to some order.
type BlockIterator interface {
	First() bool
	Next() bool
	Get() (*externalapi.DomainHash, error)
	Close() error
}
