/*
Copyright 2022 The deepauto-io LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package filestype

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDetectImageType(t *testing.T) {
	data, err := os.ReadFile("/Users/taoshumin_vendor/go/src/github.com/filestype/README.md")
	assert.NoError(t, err)

	ty := DetectImageType(data)
	t.Log(ty)
}
