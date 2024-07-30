package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStep struct {
	Description          *string                               `json:"description,omitempty"`
	Details              *DetailsInfo                          `json:"details,omitempty"`
	Name                 *string                               `json:"name,omitempty"`
	ODataType            *string                               `json:"@odata.type,omitempty"`
	ProvisioningStepType *ProvisioningStepProvisioningStepType `json:"provisioningStepType,omitempty"`
	Status               *ProvisioningStepStatus               `json:"status,omitempty"`
}
