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

const (
	// MyFiles is the files' status. [my_files] replace ace_upload
	MyFiles = "my_files"
	// Multimodal is the multimodal status. [multimodal]
	Multimodal = "multimodal"
)

// UploadUseCase is the upload use case.
// multimodal is the GPT-4 multimodal.
func UploadUseCase(str string) string {
	list := []string{MyFiles, Multimodal}
	for _, s := range list {
		if s == str {
			return str
		}
	}
	panic("not support this use case")
}

// IsMultimodal is the multimodal use case. [只支持图片]
func IsMultimodal(s string) bool {
	if s == Multimodal {
		return true
	}
	return false
}
