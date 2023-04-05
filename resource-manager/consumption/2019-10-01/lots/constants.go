package lots

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
