package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileSystemInfo struct {
	CreatedDateTime      *string `json:"createdDateTime,omitempty"`
	LastAccessedDateTime *string `json:"lastAccessedDateTime,omitempty"`
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
}
