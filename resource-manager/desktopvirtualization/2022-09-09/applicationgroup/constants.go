package applicationgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGroupType string

const (
	ApplicationGroupTypeDesktop   ApplicationGroupType = "Desktop"
	ApplicationGroupTypeRemoteApp ApplicationGroupType = "RemoteApp"
)

func PossibleValuesForApplicationGroupType() []string {
	return []string{
		string(ApplicationGroupTypeDesktop),
		string(ApplicationGroupTypeRemoteApp),
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
