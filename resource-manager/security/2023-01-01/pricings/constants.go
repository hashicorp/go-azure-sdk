package pricings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Code string

const (
	CodeFailed    Code = "Failed"
	CodeSucceeded Code = "Succeeded"
)

func PossibleValuesForCode() []string {
	return []string{
		string(CodeFailed),
		string(CodeSucceeded),
	}
}

type IsEnabled string

const (
	IsEnabledFalse IsEnabled = "False"
	IsEnabledTrue  IsEnabled = "True"
)

func PossibleValuesForIsEnabled() []string {
	return []string{
		string(IsEnabledFalse),
		string(IsEnabledTrue),
	}
}

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
