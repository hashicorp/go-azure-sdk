package elasticsans

import "strings"

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

func parseProvisioningStates(input string) (*ProvisioningStates, error) {
	vals := map[string]ProvisioningStates{
		"canceled":  ProvisioningStatesCanceled,
		"creating":  ProvisioningStatesCreating,
		"deleting":  ProvisioningStatesDeleting,
		"failed":    ProvisioningStatesFailed,
		"invalid":   ProvisioningStatesInvalid,
		"pending":   ProvisioningStatesPending,
		"succeeded": ProvisioningStatesSucceeded,
		"updating":  ProvisioningStatesUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStates(input)
	return &out, nil
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

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"premium_lrs": SkuNamePremiumLRS,
		"premium_zrs": SkuNamePremiumZRS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
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

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"premium": SkuTierPremium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}
