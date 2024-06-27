package reservationorders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Price struct {
	Amount       *float64 `json:"amount,omitempty"`
	CurrencyCode *string  `json:"currencyCode,omitempty"`
}
