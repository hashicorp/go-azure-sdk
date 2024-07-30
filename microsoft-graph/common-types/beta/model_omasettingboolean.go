package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OmaSettingBoolean struct {
	Description            *string `json:"description,omitempty"`
	DisplayName            *string `json:"displayName,omitempty"`
	IsEncrypted            *bool   `json:"isEncrypted,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
	OmaUri                 *string `json:"omaUri,omitempty"`
	SecretReferenceValueId *string `json:"secretReferenceValueId,omitempty"`
	Value                  *bool   `json:"value,omitempty"`
}
