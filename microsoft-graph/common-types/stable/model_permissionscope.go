package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionScope struct {
	AdminConsentDescription *string `json:"adminConsentDescription,omitempty"`
	AdminConsentDisplayName *string `json:"adminConsentDisplayName,omitempty"`
	Id                      *string `json:"id,omitempty"`
	IsEnabled               *bool   `json:"isEnabled,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
	Origin                  *string `json:"origin,omitempty"`
	Type                    *string `json:"type,omitempty"`
	UserConsentDescription  *string `json:"userConsentDescription,omitempty"`
	UserConsentDisplayName  *string `json:"userConsentDisplayName,omitempty"`
	Value                   *string `json:"value,omitempty"`
}
