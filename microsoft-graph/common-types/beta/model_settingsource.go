package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingSource struct {
	DisplayName *string                  `json:"displayName,omitempty"`
	Id          *string                  `json:"id,omitempty"`
	ODataType   *string                  `json:"@odata.type,omitempty"`
	SourceType  *SettingSourceSourceType `json:"sourceType,omitempty"`
}
