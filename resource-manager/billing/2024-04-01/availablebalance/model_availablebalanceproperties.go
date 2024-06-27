package availablebalance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableBalanceProperties struct {
	Amount                 *Amount             `json:"amount,omitempty"`
	PaymentsOnAccount      *[]PaymentOnAccount `json:"paymentsOnAccount,omitempty"`
	TotalPaymentsOnAccount *Amount             `json:"totalPaymentsOnAccount,omitempty"`
}
