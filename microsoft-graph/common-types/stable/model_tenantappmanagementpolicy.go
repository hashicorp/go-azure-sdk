package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantAppManagementPolicy struct {
	ApplicationRestrictions      *AppManagementConfiguration `json:"applicationRestrictions,omitempty"`
	DeletedDateTime              *string                     `json:"deletedDateTime,omitempty"`
	Description                  *string                     `json:"description,omitempty"`
	DisplayName                  *string                     `json:"displayName,omitempty"`
	Id                           *string                     `json:"id,omitempty"`
	IsEnabled                    *bool                       `json:"isEnabled,omitempty"`
	ODataType                    *string                     `json:"@odata.type,omitempty"`
	ServicePrincipalRestrictions *AppManagementConfiguration `json:"servicePrincipalRestrictions,omitempty"`
}
