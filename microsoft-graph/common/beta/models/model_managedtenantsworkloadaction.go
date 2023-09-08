package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadAction struct {
	ActionId    *string                               `json:"actionId,omitempty"`
	Category    *ManagedTenantsWorkloadActionCategory `json:"category,omitempty"`
	Description *string                               `json:"description,omitempty"`
	DisplayName *string                               `json:"displayName,omitempty"`
	Licenses    *[]string                             `json:"licenses,omitempty"`
	ODataType   *string                               `json:"@odata.type,omitempty"`
	Service     *string                               `json:"service,omitempty"`
	Settings    *[]ManagedTenantsSetting              `json:"settings,omitempty"`
}
