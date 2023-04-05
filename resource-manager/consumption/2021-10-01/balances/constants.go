package balances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingFrequency string

const (
	BillingFrequencyMonth   BillingFrequency = "Month"
	BillingFrequencyQuarter BillingFrequency = "Quarter"
	BillingFrequencyYear    BillingFrequency = "Year"
)

func PossibleValuesForBillingFrequency() []string {
	return []string{
		string(BillingFrequencyMonth),
		string(BillingFrequencyQuarter),
		string(BillingFrequencyYear),
	}
}
