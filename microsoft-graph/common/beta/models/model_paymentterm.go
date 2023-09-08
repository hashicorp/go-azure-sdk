package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PaymentTerm struct {
	CalculateDiscountOnCreditMemos *bool    `json:"calculateDiscountOnCreditMemos,omitempty"`
	Code                           *string  `json:"code,omitempty"`
	DiscountDateCalculation        *string  `json:"discountDateCalculation,omitempty"`
	DiscountPercent                *float64 `json:"discountPercent,omitempty"`
	DisplayName                    *string  `json:"displayName,omitempty"`
	DueDateCalculation             *string  `json:"dueDateCalculation,omitempty"`
	Id                             *string  `json:"id,omitempty"`
	LastModifiedDateTime           *string  `json:"lastModifiedDateTime,omitempty"`
	ODataType                      *string  `json:"@odata.type,omitempty"`
}
