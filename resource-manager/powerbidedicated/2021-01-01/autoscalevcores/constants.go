package autoscalevcores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VCoreProvisioningState string

const (
	VCoreProvisioningStateSucceeded VCoreProvisioningState = "Succeeded"
)

func PossibleValuesForVCoreProvisioningState() []string {
	return []string{
		string(VCoreProvisioningStateSucceeded),
	}
}

type VCoreSkuTier string

const (
	VCoreSkuTierAutoScale VCoreSkuTier = "AutoScale"
)

func PossibleValuesForVCoreSkuTier() []string {
	return []string{
		string(VCoreSkuTierAutoScale),
	}
}
