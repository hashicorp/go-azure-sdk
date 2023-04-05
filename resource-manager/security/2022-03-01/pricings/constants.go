package pricings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PricingTier string

const (
	PricingTierFree     PricingTier = "Free"
	PricingTierStandard PricingTier = "Standard"
)

func PossibleValuesForPricingTier() []string {
	return []string{
		string(PricingTierFree),
		string(PricingTierStandard),
	}
}
