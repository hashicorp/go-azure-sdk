package credits

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreditBalanceSummary struct {
	CurrentBalance                    *Amount                 `json:"currentBalance"`
	EstimatedBalance                  *Amount                 `json:"estimatedBalance"`
	EstimatedBalanceInBillingCurrency *AmountWithExchangeRate `json:"estimatedBalanceInBillingCurrency"`
}
