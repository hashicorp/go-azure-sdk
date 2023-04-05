package analysisservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkuTier string

const (
	SkuTierBasic       SkuTier = "Basic"
	SkuTierDevelopment SkuTier = "Development"
	SkuTierStandard    SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierDevelopment),
		string(SkuTierStandard),
	}
}
