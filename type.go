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
	"fmt"
	"log"

	"github.com/gabriel-vasile/mimetype"
)

const (
	// MyFiles is the files' status. [my_files] replace ace_upload
	MyFiles = "my_files"
	// Multimodal is the multimodal status. [multimodal]
	Multimodal = "multimodal"
	// Gizmo is the gizmo status. [gizmo]
	Gizmo = "gizmo"
)

// UploadFileType is use gizmo upload file type.
type UploadFileType string

const (
	// Image is the image type.
	Image UploadFileType = "image"
	// File is the file type.
	File UploadFileType = "file"
	// Unknown is the unknown type.
	Unknown UploadFileType = "unknown"
)

// UploadUseCase is the upload use case.
// multimodal is the GPT-4 multimodal.
func UploadUseCase(str string) string {
	list := []string{MyFiles, Multimodal, Gizmo}
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

// OpenAISupportFileAndImageType is the my files use case. [支持图片、视频、音频、文档]
func OpenAISupportFileAndImageType(filename string, data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover from a panic: %v", r)
			err = fmt.Errorf("not support this file type")
		}
	}()

	mm := mimetype.Detect(data)

	switch mm.String() {
	case "image/webp", "image/jpeg", "image/gif", "image/png":
		return
	default:
		DetectFileType(filename, data)
		return
	}
}

// GetGizmoUploadFileType is the gizmo upload file type.
func GetGizmoUploadFileType(filename string, data []byte) UploadFileType {
	mm := mimetype.Detect(data)
	switch mm.String() {
	case "image/webp", "image/jpeg", "image/gif", "image/png":
		return Image
	default:
		ft := GizmoDetectFileType(filename, data)
		if ft == "" {
			return Unknown
		}
		return File
	}
}
