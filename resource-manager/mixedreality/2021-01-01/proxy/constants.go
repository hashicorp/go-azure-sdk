package proxy

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NameUnavailableReason string

const (
	NameUnavailableReasonAlreadyExists NameUnavailableReason = "AlreadyExists"
	NameUnavailableReasonInvalid       NameUnavailableReason = "Invalid"
)

func PossibleValuesForNameUnavailableReason() []string {
	return []string{
		string(NameUnavailableReasonAlreadyExists),
		string(NameUnavailableReasonInvalid),
	}
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierFree     SkuTier = "Free"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}
