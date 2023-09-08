package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMetadataAction struct {
	MetadataToAdd    *[]SecurityKeyValuePair `json:"metadataToAdd,omitempty"`
	MetadataToRemove *[]string               `json:"metadataToRemove,omitempty"`
	ODataType        *string                 `json:"@odata.type,omitempty"`
}
