package extensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AKSIdentityType string

const (
	AKSIdentityTypeSystemAssigned AKSIdentityType = "SystemAssigned"
	AKSIdentityTypeUserAssigned   AKSIdentityType = "UserAssigned"
)

func PossibleValuesForAKSIdentityType() []string {
	return []string{
		string(AKSIdentityTypeSystemAssigned),
		string(AKSIdentityTypeUserAssigned),
	}
}

type LevelType string

const (
	LevelTypeError       LevelType = "Error"
	LevelTypeInformation LevelType = "Information"
	LevelTypeWarning     LevelType = "Warning"
)

func PossibleValuesForLevelType() []string {
	return []string{
		string(LevelTypeError),
		string(LevelTypeInformation),
		string(LevelTypeWarning),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}
