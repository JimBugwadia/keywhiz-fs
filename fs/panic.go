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

package fs

import (
	klog "github.com/square/keywhiz-fs/log"
)

var (
	logger *klog.Logger
)

// Helper function to panic on error
func panicOnError(err error) {
	if err != nil {
		logger.Errorf("panic: %v", err)
		panic(err)
	}
}
