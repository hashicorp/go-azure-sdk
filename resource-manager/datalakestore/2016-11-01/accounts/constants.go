package accounts

import "strings"

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

func parseDataLakeStoreAccountState(input string) (*DataLakeStoreAccountState, error) {
	vals := map[string]DataLakeStoreAccountState{
		"active":    DataLakeStoreAccountStateActive,
		"suspended": DataLakeStoreAccountStateSuspended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataLakeStoreAccountState(input)
	return &out, nil
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

func parseDataLakeStoreAccountStatus(input string) (*DataLakeStoreAccountStatus, error) {
	vals := map[string]DataLakeStoreAccountStatus{
		"canceled":   DataLakeStoreAccountStatusCanceled,
		"creating":   DataLakeStoreAccountStatusCreating,
		"deleted":    DataLakeStoreAccountStatusDeleted,
		"deleting":   DataLakeStoreAccountStatusDeleting,
		"failed":     DataLakeStoreAccountStatusFailed,
		"patching":   DataLakeStoreAccountStatusPatching,
		"resuming":   DataLakeStoreAccountStatusResuming,
		"running":    DataLakeStoreAccountStatusRunning,
		"succeeded":  DataLakeStoreAccountStatusSucceeded,
		"suspending": DataLakeStoreAccountStatusSuspending,
		"undeleting": DataLakeStoreAccountStatusUndeleting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataLakeStoreAccountStatus(input)
	return &out, nil
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

func parseEncryptionConfigType(input string) (*EncryptionConfigType, error) {
	vals := map[string]EncryptionConfigType{
		"servicemanaged": EncryptionConfigTypeServiceManaged,
		"usermanaged":    EncryptionConfigTypeUserManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionConfigType(input)
	return &out, nil
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

func parseEncryptionProvisioningState(input string) (*EncryptionProvisioningState, error) {
	vals := map[string]EncryptionProvisioningState{
		"creating":  EncryptionProvisioningStateCreating,
		"succeeded": EncryptionProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionProvisioningState(input)
	return &out, nil
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

func parseEncryptionState(input string) (*EncryptionState, error) {
	vals := map[string]EncryptionState{
		"disabled": EncryptionStateDisabled,
		"enabled":  EncryptionStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionState(input)
	return &out, nil
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

func parseFirewallAllowAzureIPsState(input string) (*FirewallAllowAzureIPsState, error) {
	vals := map[string]FirewallAllowAzureIPsState{
		"disabled": FirewallAllowAzureIPsStateDisabled,
		"enabled":  FirewallAllowAzureIPsStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FirewallAllowAzureIPsState(input)
	return &out, nil
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

func parseFirewallState(input string) (*FirewallState, error) {
	vals := map[string]FirewallState{
		"disabled": FirewallStateDisabled,
		"enabled":  FirewallStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FirewallState(input)
	return &out, nil
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

func parseTierType(input string) (*TierType, error) {
	vals := map[string]TierType{
		"commitment_500tb": TierTypeCommitmentFiveHundredTB,
		"commitment_5pb":   TierTypeCommitmentFivePB,
		"commitment_100tb": TierTypeCommitmentOneHundredTB,
		"commitment_1pb":   TierTypeCommitmentOnePB,
		"commitment_1tb":   TierTypeCommitmentOneTB,
		"commitment_10tb":  TierTypeCommitmentOneZeroTB,
		"consumption":      TierTypeConsumption,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TierType(input)
	return &out, nil
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

func parseTrustedIdProviderState(input string) (*TrustedIdProviderState, error) {
	vals := map[string]TrustedIdProviderState{
		"disabled": TrustedIdProviderStateDisabled,
		"enabled":  TrustedIdProviderStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrustedIdProviderState(input)
	return &out, nil
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

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"microsoft.datalakestore/accounts": TypeMicrosoftPointDataLakeStoreAccounts,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
