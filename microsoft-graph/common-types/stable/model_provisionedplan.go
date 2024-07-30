package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionedPlan struct {
	CapabilityStatus   *string `json:"capabilityStatus,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	ProvisioningStatus *string `json:"provisioningStatus,omitempty"`
	Service            *string `json:"service,omitempty"`
}
