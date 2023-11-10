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

import "github.com/gabriel-vasile/mimetype"

// DetectImageType returns the image type of the data
func DetectImageType(data []byte) string {
	mm := mimetype.Detect(data)

	// todo OpenAI will add many image type in the future
	// "image/webp",
	// "image/jpeg",
	// "image/gif",
	// "image/png"

	switch mm.String() {
	case "image/webp":
		return mm.String()
	case "image/jpeg":
		return mm.String()
	case "image/gif":
		return mm.String()
	case "image/png":
		return mm.String()
	default:
		panic("unknown image type")
	}
}
