package common

import (
	"bytes"
	"sort"
)

//----------------------------------------
// KVPair

/*
Defined in types.proto

type KVPair struct {
	Key   []byte
	Value []byte
}
*/

type KVPairs []KVPair

// Sorting
func (kvs KVPairs) Len() int { return len(kvs) }
func (kvs KVPairs) Less(i, j int) bool {
	switch bytes.Compare(kvs[i].Key, kvs[j].Key) {
	case -1:
		return true
	case 0:
		return bytes.Compare(kvs[i].Value, kvs[j].Value) < 0
	case 1:
		return false
	default:
		panic("invalid comparison result")
	}
}
func (kvs KVPairs) Swap(i, j int) { kvs[i], kvs[j] = kvs[j], kvs[i] }
func (kvs KVPairs) Sort()         { sort.Sort(kvs) }
