package billingsubscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RenewalTermDetails struct {
	BillingFrequency *string `json:"billingFrequency,omitempty"`
	ProductTypeId    *string `json:"productTypeId,omitempty"`
	Quantity         *int64  `json:"quantity,omitempty"`
	SkuId            *string `json:"skuId,omitempty"`
	TermDuration     *string `json:"termDuration,omitempty"`
}
