package powerbidedicated

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacitySkuTier string

const (
	CapacitySkuTierAutoPremiumHost CapacitySkuTier = "AutoPremiumHost"
	CapacitySkuTierPBIEAzure       CapacitySkuTier = "PBIE_Azure"
	CapacitySkuTierPremium         CapacitySkuTier = "Premium"
)

func PossibleValuesForCapacitySkuTier() []string {
	return []string{
		string(CapacitySkuTierAutoPremiumHost),
		string(CapacitySkuTierPBIEAzure),
		string(CapacitySkuTierPremium),
	}
}
