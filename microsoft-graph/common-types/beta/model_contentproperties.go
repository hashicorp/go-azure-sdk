package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentProperties struct {
	Extensions           *[]string        `json:"extensions,omitempty"`
	LastModifiedBy       *string          `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string          `json:"lastModifiedDateTime,omitempty"`
	Metadata             *ContentMetadata `json:"metadata,omitempty"`
	ODataType            *string          `json:"@odata.type,omitempty"`
}
