package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LicenseAssignmentState struct {
	AssignedByGroup     *string   `json:"assignedByGroup,omitempty"`
	DisabledPlans       *[]string `json:"disabledPlans,omitempty"`
	Error               *string   `json:"error,omitempty"`
	LastUpdatedDateTime *string   `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string   `json:"@odata.type,omitempty"`
	SkuId               *string   `json:"skuId,omitempty"`
	State               *string   `json:"state,omitempty"`
}
