package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePlanInfo struct {
	AppliesTo          *string `json:"appliesTo,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	ProvisioningStatus *string `json:"provisioningStatus,omitempty"`
	ServicePlanId      *string `json:"servicePlanId,omitempty"`
	ServicePlanName    *string `json:"servicePlanName,omitempty"`
}
