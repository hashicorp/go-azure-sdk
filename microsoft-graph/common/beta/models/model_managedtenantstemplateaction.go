package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTemplateAction struct {
	Description      *string                  `json:"description,omitempty"`
	DisplayName      *string                  `json:"displayName,omitempty"`
	Licenses         *LicenseDetails          `json:"licenses,omitempty"`
	ODataType        *string                  `json:"@odata.type,omitempty"`
	Service          *string                  `json:"service,omitempty"`
	Settings         *[]ManagedTenantsSetting `json:"settings,omitempty"`
	TemplateActionId *string                  `json:"templateActionId,omitempty"`
}
