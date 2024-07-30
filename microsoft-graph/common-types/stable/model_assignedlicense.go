package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedLicense struct {
	DisabledPlans *[]string `json:"disabledPlans,omitempty"`
	ODataType     *string   `json:"@odata.type,omitempty"`
	SkuId         *string   `json:"skuId,omitempty"`
}
