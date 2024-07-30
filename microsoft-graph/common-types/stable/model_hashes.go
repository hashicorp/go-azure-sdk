package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Hashes struct {
	Crc32Hash    *string `json:"crc32Hash,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	QuickXorHash *string `json:"quickXorHash,omitempty"`
	Sha1Hash     *string `json:"sha1Hash,omitempty"`
	Sha256Hash   *string `json:"sha256Hash,omitempty"`
}
