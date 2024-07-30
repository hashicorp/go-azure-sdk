package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type File struct {
	Hashes             *Hashes `json:"hashes,omitempty"`
	MimeType           *string `json:"mimeType,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	ProcessingMetadata *bool   `json:"processingMetadata,omitempty"`
}
