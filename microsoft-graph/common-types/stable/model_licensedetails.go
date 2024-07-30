package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LicenseDetails struct {
	Id            *string            `json:"id,omitempty"`
	ODataType     *string            `json:"@odata.type,omitempty"`
	ServicePlans  *[]ServicePlanInfo `json:"servicePlans,omitempty"`
	SkuId         *string            `json:"skuId,omitempty"`
	SkuPartNumber *string            `json:"skuPartNumber,omitempty"`
}
