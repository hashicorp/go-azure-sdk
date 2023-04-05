package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataLakeStoreAccountState string

const (
	DataLakeStoreAccountStateActive    DataLakeStoreAccountState = "Active"
	DataLakeStoreAccountStateSuspended DataLakeStoreAccountState = "Suspended"
)

func PossibleValuesForDataLakeStoreAccountState() []string {
	return []string{
		string(DataLakeStoreAccountStateActive),
		string(DataLakeStoreAccountStateSuspended),
	}
}

type DataLakeStoreAccountStatus string

const (
	DataLakeStoreAccountStatusCanceled   DataLakeStoreAccountStatus = "Canceled"
	DataLakeStoreAccountStatusCreating   DataLakeStoreAccountStatus = "Creating"
	DataLakeStoreAccountStatusDeleted    DataLakeStoreAccountStatus = "Deleted"
	DataLakeStoreAccountStatusDeleting   DataLakeStoreAccountStatus = "Deleting"
	DataLakeStoreAccountStatusFailed     DataLakeStoreAccountStatus = "Failed"
	DataLakeStoreAccountStatusPatching   DataLakeStoreAccountStatus = "Patching"
	DataLakeStoreAccountStatusResuming   DataLakeStoreAccountStatus = "Resuming"
	DataLakeStoreAccountStatusRunning    DataLakeStoreAccountStatus = "Running"
	DataLakeStoreAccountStatusSucceeded  DataLakeStoreAccountStatus = "Succeeded"
	DataLakeStoreAccountStatusSuspending DataLakeStoreAccountStatus = "Suspending"
	DataLakeStoreAccountStatusUndeleting DataLakeStoreAccountStatus = "Undeleting"
)

func PossibleValuesForDataLakeStoreAccountStatus() []string {
	return []string{
		string(DataLakeStoreAccountStatusCanceled),
		string(DataLakeStoreAccountStatusCreating),
		string(DataLakeStoreAccountStatusDeleted),
		string(DataLakeStoreAccountStatusDeleting),
		string(DataLakeStoreAccountStatusFailed),
		string(DataLakeStoreAccountStatusPatching),
		string(DataLakeStoreAccountStatusResuming),
		string(DataLakeStoreAccountStatusRunning),
		string(DataLakeStoreAccountStatusSucceeded),
		string(DataLakeStoreAccountStatusSuspending),
		string(DataLakeStoreAccountStatusUndeleting),
	}
}

type EncryptionConfigType string

const (
	EncryptionConfigTypeServiceManaged EncryptionConfigType = "ServiceManaged"
	EncryptionConfigTypeUserManaged    EncryptionConfigType = "UserManaged"
)

func PossibleValuesForEncryptionConfigType() []string {
	return []string{
		string(EncryptionConfigTypeServiceManaged),
		string(EncryptionConfigTypeUserManaged),
	}
}

type EncryptionProvisioningState string

const (
	EncryptionProvisioningStateCreating  EncryptionProvisioningState = "Creating"
	EncryptionProvisioningStateSucceeded EncryptionProvisioningState = "Succeeded"
)

func PossibleValuesForEncryptionProvisioningState() []string {
	return []string{
		string(EncryptionProvisioningStateCreating),
		string(EncryptionProvisioningStateSucceeded),
	}
}

type EncryptionState string

const (
	EncryptionStateDisabled EncryptionState = "Disabled"
	EncryptionStateEnabled  EncryptionState = "Enabled"
)

func PossibleValuesForEncryptionState() []string {
	return []string{
		string(EncryptionStateDisabled),
		string(EncryptionStateEnabled),
	}
}

type FirewallAllowAzureIPsState string

const (
	FirewallAllowAzureIPsStateDisabled FirewallAllowAzureIPsState = "Disabled"
	FirewallAllowAzureIPsStateEnabled  FirewallAllowAzureIPsState = "Enabled"
)

func PossibleValuesForFirewallAllowAzureIPsState() []string {
	return []string{
		string(FirewallAllowAzureIPsStateDisabled),
		string(FirewallAllowAzureIPsStateEnabled),
	}
}

type FirewallState string

const (
	FirewallStateDisabled FirewallState = "Disabled"
	FirewallStateEnabled  FirewallState = "Enabled"
)

func PossibleValuesForFirewallState() []string {
	return []string{
		string(FirewallStateDisabled),
		string(FirewallStateEnabled),
	}
}

type TierType string

const (
	TierTypeCommitmentFiveHundredTB TierType = "Commitment_500TB"
	TierTypeCommitmentFivePB        TierType = "Commitment_5PB"
	TierTypeCommitmentOneHundredTB  TierType = "Commitment_100TB"
	TierTypeCommitmentOnePB         TierType = "Commitment_1PB"
	TierTypeCommitmentOneTB         TierType = "Commitment_1TB"
	TierTypeCommitmentOneZeroTB     TierType = "Commitment_10TB"
	TierTypeConsumption             TierType = "Consumption"
)

func PossibleValuesForTierType() []string {
	return []string{
		string(TierTypeCommitmentFiveHundredTB),
		string(TierTypeCommitmentFivePB),
		string(TierTypeCommitmentOneHundredTB),
		string(TierTypeCommitmentOnePB),
		string(TierTypeCommitmentOneTB),
		string(TierTypeCommitmentOneZeroTB),
		string(TierTypeConsumption),
	}
}

type TrustedIdProviderState string

const (
	TrustedIdProviderStateDisabled TrustedIdProviderState = "Disabled"
	TrustedIdProviderStateEnabled  TrustedIdProviderState = "Enabled"
)

func PossibleValuesForTrustedIdProviderState() []string {
	return []string{
		string(TrustedIdProviderStateDisabled),
		string(TrustedIdProviderStateEnabled),
	}
}

type Type string

const (
	TypeMicrosoftPointDataLakeStoreAccounts Type = "Microsoft.DataLakeStore/accounts"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeMicrosoftPointDataLakeStoreAccounts),
	}
}
