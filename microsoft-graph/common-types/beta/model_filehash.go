package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileHash struct {
	HashType  *FileHashHashType `json:"hashType,omitempty"`
	HashValue *string           `json:"hashValue,omitempty"`
	ODataType *string           `json:"@odata.type,omitempty"`
}
