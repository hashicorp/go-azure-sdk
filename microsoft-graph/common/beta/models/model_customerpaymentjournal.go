package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomerPaymentJournal struct {
	Account                *Account           `json:"account,omitempty"`
	BalancingAccountId     *string            `json:"balancingAccountId,omitempty"`
	BalancingAccountNumber *string            `json:"balancingAccountNumber,omitempty"`
	Code                   *string            `json:"code,omitempty"`
	CustomerPayments       *[]CustomerPayment `json:"customerPayments,omitempty"`
	DisplayName            *string            `json:"displayName,omitempty"`
	Id                     *string            `json:"id,omitempty"`
	LastModifiedDateTime   *string            `json:"lastModifiedDateTime,omitempty"`
	ODataType              *string            `json:"@odata.type,omitempty"`
}
