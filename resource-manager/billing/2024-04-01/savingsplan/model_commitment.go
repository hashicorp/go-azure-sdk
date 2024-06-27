package savingsplan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Commitment struct {
	Amount       *float64         `json:"amount,omitempty"`
	CurrencyCode *string          `json:"currencyCode,omitempty"`
	Grain        *CommitmentGrain `json:"grain,omitempty"`
}
