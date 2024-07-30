package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementAction struct {
	Category                 *ManagedTenantsManagementActionCategory `json:"category,omitempty"`
	Description              *string                                 `json:"description,omitempty"`
	DisplayName              *string                                 `json:"displayName,omitempty"`
	Id                       *string                                 `json:"id,omitempty"`
	ODataType                *string                                 `json:"@odata.type,omitempty"`
	ReferenceTemplateId      *string                                 `json:"referenceTemplateId,omitempty"`
	ReferenceTemplateVersion *int64                                  `json:"referenceTemplateVersion,omitempty"`
	WorkloadActions          *[]ManagedTenantsWorkloadAction         `json:"workloadActions,omitempty"`
}
