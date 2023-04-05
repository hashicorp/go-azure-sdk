package availableservicetiers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkuNameEnum string

const (
	SkuNameEnumCapacityReservation  SkuNameEnum = "CapacityReservation"
	SkuNameEnumFree                 SkuNameEnum = "Free"
	SkuNameEnumPerGBTwoZeroOneEight SkuNameEnum = "PerGB2018"
	SkuNameEnumPerNode              SkuNameEnum = "PerNode"
	SkuNameEnumPremium              SkuNameEnum = "Premium"
	SkuNameEnumStandalone           SkuNameEnum = "Standalone"
	SkuNameEnumStandard             SkuNameEnum = "Standard"
)

func PossibleValuesForSkuNameEnum() []string {
	return []string{
		string(SkuNameEnumCapacityReservation),
		string(SkuNameEnumFree),
		string(SkuNameEnumPerGBTwoZeroOneEight),
		string(SkuNameEnumPerNode),
		string(SkuNameEnumPremium),
		string(SkuNameEnumStandalone),
		string(SkuNameEnumStandard),
	}
}
