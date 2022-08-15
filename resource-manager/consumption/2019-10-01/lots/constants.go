package lots

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LotSource string

const (
	LotSourcePromotionalCredit LotSource = "PromotionalCredit"
	LotSourcePurchasedCredit   LotSource = "PurchasedCredit"
)

func PossibleValuesForLotSource() []string {
	return []string{
		string(LotSourcePromotionalCredit),
		string(LotSourcePurchasedCredit),
	}
}

func parseLotSource(input string) (*LotSource, error) {
	vals := map[string]LotSource{
		"promotionalcredit": LotSourcePromotionalCredit,
		"purchasedcredit":   LotSourcePurchasedCredit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LotSource(input)
	return &out, nil
}
