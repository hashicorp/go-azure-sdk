package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementIntentInfo struct {
	ManagementIntentDisplayName *string                                         `json:"managementIntentDisplayName,omitempty"`
	ManagementIntentId          *string                                         `json:"managementIntentId,omitempty"`
	ManagementTemplates         *[]ManagedTenantsManagementTemplateDetailedInfo `json:"managementTemplates,omitempty"`
	ODataType                   *string                                         `json:"@odata.type,omitempty"`
}
