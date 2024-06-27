package transaction

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransactionSummary struct {
	AzureCreditApplied               *float64 `json:"azureCreditApplied,omitempty"`
	BillingCurrency                  *string  `json:"billingCurrency,omitempty"`
	ConsumptionCommitmentDecremented *float64 `json:"consumptionCommitmentDecremented,omitempty"`
	SubTotal                         *float64 `json:"subTotal,omitempty"`
	Tax                              *float64 `json:"tax,omitempty"`
	Total                            *float64 `json:"total,omitempty"`
}
