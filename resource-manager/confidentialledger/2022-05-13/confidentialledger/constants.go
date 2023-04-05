package confidentialledger

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LedgerRoleName string

const (
	LedgerRoleNameAdministrator LedgerRoleName = "Administrator"
	LedgerRoleNameContributor   LedgerRoleName = "Contributor"
	LedgerRoleNameReader        LedgerRoleName = "Reader"
)

func PossibleValuesForLedgerRoleName() []string {
	return []string{
		string(LedgerRoleNameAdministrator),
		string(LedgerRoleNameContributor),
		string(LedgerRoleNameReader),
	}
}

type LedgerType string

const (
	LedgerTypePrivate LedgerType = "Private"
	LedgerTypePublic  LedgerType = "Public"
	LedgerTypeUnknown LedgerType = "Unknown"
)

func PossibleValuesForLedgerType() []string {
	return []string{
		string(LedgerTypePrivate),
		string(LedgerTypePublic),
		string(LedgerTypeUnknown),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnknown),
		string(ProvisioningStateUpdating),
	}
}
