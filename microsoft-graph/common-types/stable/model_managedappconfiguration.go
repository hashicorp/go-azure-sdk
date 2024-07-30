package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppConfiguration struct {
	CreatedDateTime      *string         `json:"createdDateTime,omitempty"`
	CustomSettings       *[]KeyValuePair `json:"customSettings,omitempty"`
	Description          *string         `json:"description,omitempty"`
	DisplayName          *string         `json:"displayName,omitempty"`
	Id                   *string         `json:"id,omitempty"`
	LastModifiedDateTime *string         `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string         `json:"@odata.type,omitempty"`
	Version              *string         `json:"version,omitempty"`
}
