package credits

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AmountWithExchangeRate struct {
	Currency          *string  `json:"currency,omitempty"`
	ExchangeRate      *float64 `json:"exchangeRate,omitempty"`
	ExchangeRateMonth *float64 `json:"exchangeRateMonth,omitempty"`
	Value             *float64 `json:"value,omitempty"`
}
