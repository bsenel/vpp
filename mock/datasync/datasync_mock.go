// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datasync

import (
	"github.com/golang/protobuf/proto"
	"github.com/ligato/cn-infra/datasync"
	"strings"
)

// MockDataSync can be used to generate datasync events from provided data.
type MockDataSync struct {
	data     map[string]*ProtoData
	anyError error
}

// ProtoData is used to store proto message with revision.
type ProtoData struct {
	val proto.Message
	rev int64
}

// MockKeyVal implements KeyVal interface.
type MockKeyVal struct {
	key string
	val proto.Message
	rev int64
}

// MockChangeEvent implements ChangeEvent interface.
type MockChangeEvent struct {
	mds *MockDataSync
	MockKeyVal
	prevVal proto.Message
}

// MockResyncEvent implements ResyncEvent interface.
type MockResyncEvent struct {
	mds         *MockDataSync
	keyPrefixes []string
}

// MockKeyValIterator implements KeyValIterator interface.
type MockKeyValIterator struct {
	mds    *MockDataSync
	keys   []string
	cursor int
}

//// Data Sync ////

// NewMockDataSync is a constructor for MockDataSync.
func NewMockDataSync() *MockDataSync {
	return &MockDataSync{
		data: make(map[string]*ProtoData),
	}
}

// Put allows to put a new value under the given key and to get the corresponding
// data change event.
func (mds *MockDataSync) Put(key string, value proto.Message) datasync.ChangeEvent {
	var prevValue proto.Message
	if value == nil {
		return mds.Delete(key)
	}
	if _, modify := mds.data[key]; modify {
		prevValue = mds.data[key].val
		mds.data[key].val = value
		mds.data[key].rev++
	} else {
		mds.data[key] = &ProtoData{
			val: value,
			rev: 0,
		}
	}
	return &MockChangeEvent{
		mds: mds,
		MockKeyVal: MockKeyVal{
			key: key,
			val: value,
			rev: mds.data[key].rev,
		},
		prevVal: prevValue,
	}
}

// Delete allows to remove value under the given key and to get the corresponding
// data change event.
func (mds *MockDataSync) Delete(key string) datasync.ChangeEvent {
	if _, found := mds.data[key]; !found {
		return nil
	}
	rev := mds.data[key].rev
	delete(mds.data, key)
	return &MockChangeEvent{
		mds: mds,
		MockKeyVal: MockKeyVal{
			key: key,
			rev: rev,
		},
	}
}

// Resync returns resync event corresponding to a given list of key prefixes
// and the current state of the mocked data store.
func (mds *MockDataSync) Resync(keyPrefix ...string) datasync.ResyncEvent {
	return &MockResyncEvent{
		mds:         mds,
		keyPrefixes: keyPrefix,
	}
}

// AnyError returns non-nil if any data change or resync event was processed
// unsuccessfully.
func (mds *MockDataSync) AnyError() error {
	return mds.anyError
}

//// Key-Value ////

// GetValue returns the associated value.
func (mkv *MockKeyVal) GetValue(value proto.Message) error {
	tmp, err := proto.Marshal(mkv.val)
	if err != nil {
		return err
	}
	return proto.Unmarshal(tmp, value)
}

// GetRevision returns the associated revision.
func (mkv *MockKeyVal) GetRevision() (rev int64) {
	return mkv.rev
}

// GetKey returns the associated key.
func (mkv *MockKeyVal) GetKey() string {
	return mkv.key
}

//// Change Event ////

// Done stores non-nil error to MockDataSync.
func (mche *MockChangeEvent) Done(err error) {
	if err != nil {
		mche.mds.anyError = err
	}
}

// GetChangeType returns either "Put" or "Delete".
func (mche *MockChangeEvent) GetChangeType() datasync.PutDel {
	if mche.val == nil {
		return datasync.Delete
	}
	return datasync.Put
}

// GetPrevValue returns the previous value.
func (mche *MockChangeEvent) GetPrevValue(prevValue proto.Message) (prevValueExist bool, err error) {
	if mche.prevVal == nil {
		return false, nil
	}
	tmp, err := proto.Marshal(mche.prevVal)
	if err != nil {
		return true, err
	}
	return true, proto.Unmarshal(tmp, prevValue)
}

//// Resync Event ////

// Done stores non-nil error to MockDataSync.
func (mche *MockResyncEvent) Done(err error) {
	if err != nil {
		mche.mds.anyError = err
	}
}

// GetValues returns map "key-prefix->iterator".
func (mche *MockResyncEvent) GetValues() map[ /*keyPrefix*/ string]datasync.KeyValIterator {
	values := make(map[string]datasync.KeyValIterator)
	for _, prefix := range mche.keyPrefixes {
		var keys []string
		for key := range mche.mds.data {
			if strings.HasPrefix(key, prefix) {
				keys = append(keys, key)
			}
		}
		if len(keys) > 0 {
			values[prefix] = &MockKeyValIterator{
				mds:  mche.mds,
				keys: keys,
			}
		}
	}

	return values
}

//// Key Value Iterator ////

// GetNext returns the next item in the list.
func (mkvi *MockKeyValIterator) GetNext() (kv datasync.KeyVal, allReceived bool) {
	key := mkvi.keys[mkvi.cursor]
	kv = &MockKeyVal{
		key: key,
		val: mkvi.mds.data[key].val,
		rev: mkvi.mds.data[key].rev,
	}
	mkvi.cursor++
	return kv, mkvi.cursor == len(mkvi.keys)
}
