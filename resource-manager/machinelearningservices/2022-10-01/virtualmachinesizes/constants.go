package virtualmachinesizes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingCurrency string

const (
	BillingCurrencyUSD BillingCurrency = "USD"
)

func PossibleValuesForBillingCurrency() []string {
	return []string{
		string(BillingCurrencyUSD),
	}
}

type UnitOfMeasure string

const (
	UnitOfMeasureOneHour UnitOfMeasure = "OneHour"
)

func PossibleValuesForUnitOfMeasure() []string {
	return []string{
		string(UnitOfMeasureOneHour),
	}
}

type VMPriceOSType string

const (
	VMPriceOSTypeLinux   VMPriceOSType = "Linux"
	VMPriceOSTypeWindows VMPriceOSType = "Windows"
)

func PossibleValuesForVMPriceOSType() []string {
	return []string{
		string(VMPriceOSTypeLinux),
		string(VMPriceOSTypeWindows),
	}
}

type VMTier string

const (
	VMTierLowPriority VMTier = "LowPriority"
	VMTierSpot        VMTier = "Spot"
	VMTierStandard    VMTier = "Standard"
)

func PossibleValuesForVMTier() []string {
	return []string{
		string(VMTierLowPriority),
		string(VMTierSpot),
		string(VMTierStandard),
	}
}
