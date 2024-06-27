package billingsubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingSubscriptionSplitRequest struct {
	BillingFrequency    *string `json:"billingFrequency,omitempty"`
	Quantity            *int64  `json:"quantity,omitempty"`
	TargetProductTypeId *string `json:"targetProductTypeId,omitempty"`
	TargetSkuId         *string `json:"targetSkuId,omitempty"`
	TermDuration        *string `json:"termDuration,omitempty"`
}
