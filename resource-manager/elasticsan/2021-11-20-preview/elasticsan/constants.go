package elasticsan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStates string

const (
	ProvisioningStatesCanceled  ProvisioningStates = "Canceled"
	ProvisioningStatesCreating  ProvisioningStates = "Creating"
	ProvisioningStatesDeleting  ProvisioningStates = "Deleting"
	ProvisioningStatesFailed    ProvisioningStates = "Failed"
	ProvisioningStatesInvalid   ProvisioningStates = "Invalid"
	ProvisioningStatesPending   ProvisioningStates = "Pending"
	ProvisioningStatesSucceeded ProvisioningStates = "Succeeded"
	ProvisioningStatesUpdating  ProvisioningStates = "Updating"
)

func PossibleValuesForProvisioningStates() []string {
	return []string{
		string(ProvisioningStatesCanceled),
		string(ProvisioningStatesCreating),
		string(ProvisioningStatesDeleting),
		string(ProvisioningStatesFailed),
		string(ProvisioningStatesInvalid),
		string(ProvisioningStatesPending),
		string(ProvisioningStatesSucceeded),
		string(ProvisioningStatesUpdating),
	}
}

type SkuName string

const (
	SkuNamePremiumLRS SkuName = "Premium_LRS"
	SkuNamePremiumZRS SkuName = "Premium_ZRS"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNamePremiumLRS),
		string(SkuNamePremiumZRS),
	}
}

type SkuTier string

const (
	SkuTierPremium SkuTier = "Premium"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierPremium),
	}
}
