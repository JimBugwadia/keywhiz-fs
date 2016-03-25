// Copyright 2015 Square Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOwnershipInvalidUser(t *testing.T) {
	ownership := NewOwnership("invalid", "invalid")
	assert.NotNil(t, ownership, "should never return nil")
}

func TestGroupFileParsingValid(t *testing.T) {
	file, err := ioutil.TempFile("", "keywhiz-fs-test")
	panicOnError(err)

	_, err = file.WriteString("test:x:1234:test0,test1\n")
	panicOnError(err)
	file.Sync()
	file.Seek(0, 0)

	gid, err := lookupGidInFile("test", file)
	assert.Nil(t, err)
	assert.EqualValues(t, 1234, gid)
}
