package invoices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PaymentProperties struct {
	Amount              *Amount              `json:"amount,omitempty"`
	Date                *string              `json:"date,omitempty"`
	PaymentMethodFamily *PaymentMethodFamily `json:"paymentMethodFamily,omitempty"`
	PaymentMethodType   *string              `json:"paymentMethodType,omitempty"`
	PaymentType         *string              `json:"paymentType,omitempty"`
}
