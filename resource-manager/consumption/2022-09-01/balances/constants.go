package balances

import "strings"

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

func parseBillingFrequency(input string) (*BillingFrequency, error) {
	vals := map[string]BillingFrequency{
		"month":   BillingFrequencyMonth,
		"quarter": BillingFrequencyQuarter,
		"year":    BillingFrequencyYear,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingFrequency(input)
	return &out, nil
}
