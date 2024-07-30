package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OmaSettingInteger struct {
	Description            *string `json:"description,omitempty"`
	DisplayName            *string `json:"displayName,omitempty"`
	IsEncrypted            *bool   `json:"isEncrypted,omitempty"`
	IsReadOnly             *bool   `json:"isReadOnly,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
	OmaUri                 *string `json:"omaUri,omitempty"`
	SecretReferenceValueId *string `json:"secretReferenceValueId,omitempty"`
	Value                  *int64  `json:"value,omitempty"`
}
