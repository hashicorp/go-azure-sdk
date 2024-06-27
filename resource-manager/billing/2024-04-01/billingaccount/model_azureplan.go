package billingaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzurePlan struct {
	ProductId      *string `json:"productId,omitempty"`
	SkuDescription *string `json:"skuDescription,omitempty"`
	SkuId          *string `json:"skuId,omitempty"`
}
