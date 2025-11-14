package creditsummaries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreditBalanceSummary struct {
	CurrentBalance                    *Amount                 `json:"currentBalance,omitempty"`
	EstimatedBalance                  *Amount                 `json:"estimatedBalance,omitempty"`
	EstimatedBalanceInBillingCurrency *AmountWithExchangeRate `json:"estimatedBalanceInBillingCurrency,omitempty"`
}
