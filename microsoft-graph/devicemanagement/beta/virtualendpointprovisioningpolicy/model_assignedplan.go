package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedPlan struct {
	AssignedDateTime *string `json:"assignedDateTime,omitempty"`
	CapabilityStatus *string `json:"capabilityStatus,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	Service          *string `json:"service,omitempty"`
	ServicePlanId    *string `json:"servicePlanId,omitempty"`
}
