package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AADObjectType string

const (
	AADObjectTypeGroup            AADObjectType = "Group"
	AADObjectTypeServicePrincipal AADObjectType = "ServicePrincipal"
	AADObjectTypeUser             AADObjectType = "User"
)

func PossibleValuesForAADObjectType() []string {
	return []string{
		string(AADObjectTypeGroup),
		string(AADObjectTypeServicePrincipal),
		string(AADObjectTypeUser),
	}
}

type DataLakeAnalyticsAccountState string

const (
	DataLakeAnalyticsAccountStateActive    DataLakeAnalyticsAccountState = "Active"
	DataLakeAnalyticsAccountStateSuspended DataLakeAnalyticsAccountState = "Suspended"
)

func PossibleValuesForDataLakeAnalyticsAccountState() []string {
	return []string{
		string(DataLakeAnalyticsAccountStateActive),
		string(DataLakeAnalyticsAccountStateSuspended),
	}
}

type DataLakeAnalyticsAccountStatus string

const (
	DataLakeAnalyticsAccountStatusCanceled   DataLakeAnalyticsAccountStatus = "Canceled"
	DataLakeAnalyticsAccountStatusCreating   DataLakeAnalyticsAccountStatus = "Creating"
	DataLakeAnalyticsAccountStatusDeleted    DataLakeAnalyticsAccountStatus = "Deleted"
	DataLakeAnalyticsAccountStatusDeleting   DataLakeAnalyticsAccountStatus = "Deleting"
	DataLakeAnalyticsAccountStatusFailed     DataLakeAnalyticsAccountStatus = "Failed"
	DataLakeAnalyticsAccountStatusPatching   DataLakeAnalyticsAccountStatus = "Patching"
	DataLakeAnalyticsAccountStatusResuming   DataLakeAnalyticsAccountStatus = "Resuming"
	DataLakeAnalyticsAccountStatusRunning    DataLakeAnalyticsAccountStatus = "Running"
	DataLakeAnalyticsAccountStatusSucceeded  DataLakeAnalyticsAccountStatus = "Succeeded"
	DataLakeAnalyticsAccountStatusSuspending DataLakeAnalyticsAccountStatus = "Suspending"
	DataLakeAnalyticsAccountStatusUndeleting DataLakeAnalyticsAccountStatus = "Undeleting"
)

func PossibleValuesForDataLakeAnalyticsAccountStatus() []string {
	return []string{
		string(DataLakeAnalyticsAccountStatusCanceled),
		string(DataLakeAnalyticsAccountStatusCreating),
		string(DataLakeAnalyticsAccountStatusDeleted),
		string(DataLakeAnalyticsAccountStatusDeleting),
		string(DataLakeAnalyticsAccountStatusFailed),
		string(DataLakeAnalyticsAccountStatusPatching),
		string(DataLakeAnalyticsAccountStatusResuming),
		string(DataLakeAnalyticsAccountStatusRunning),
		string(DataLakeAnalyticsAccountStatusSucceeded),
		string(DataLakeAnalyticsAccountStatusSuspending),
		string(DataLakeAnalyticsAccountStatusUndeleting),
	}
}

type DebugDataAccessLevel string

const (
	DebugDataAccessLevelAll      DebugDataAccessLevel = "All"
	DebugDataAccessLevelCustomer DebugDataAccessLevel = "Customer"
	DebugDataAccessLevelNone     DebugDataAccessLevel = "None"
)

func PossibleValuesForDebugDataAccessLevel() []string {
	return []string{
		string(DebugDataAccessLevelAll),
		string(DebugDataAccessLevelCustomer),
		string(DebugDataAccessLevelNone),
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

type NestedResourceProvisioningState string

const (
	NestedResourceProvisioningStateCanceled  NestedResourceProvisioningState = "Canceled"
	NestedResourceProvisioningStateFailed    NestedResourceProvisioningState = "Failed"
	NestedResourceProvisioningStateSucceeded NestedResourceProvisioningState = "Succeeded"
)

func PossibleValuesForNestedResourceProvisioningState() []string {
	return []string{
		string(NestedResourceProvisioningStateCanceled),
		string(NestedResourceProvisioningStateFailed),
		string(NestedResourceProvisioningStateSucceeded),
	}
}

type TierType string

const (
	TierTypeCommitmentFiveHundredAUHours         TierType = "Commitment_500AUHours"
	TierTypeCommitmentFiveHundredThousandAUHours TierType = "Commitment_500000AUHours"
	TierTypeCommitmentFiveThousandAUHours        TierType = "Commitment_5000AUHours"
	TierTypeCommitmentFiveZeroThousandAUHours    TierType = "Commitment_50000AUHours"
	TierTypeCommitmentOneHundredAUHours          TierType = "Commitment_100AUHours"
	TierTypeCommitmentOneHundredThousandAUHours  TierType = "Commitment_100000AUHours"
	TierTypeCommitmentOneThousandAUHours         TierType = "Commitment_1000AUHours"
	TierTypeCommitmentOneZeroThousandAUHours     TierType = "Commitment_10000AUHours"
	TierTypeConsumption                          TierType = "Consumption"
)

func PossibleValuesForTierType() []string {
	return []string{
		string(TierTypeCommitmentFiveHundredAUHours),
		string(TierTypeCommitmentFiveHundredThousandAUHours),
		string(TierTypeCommitmentFiveThousandAUHours),
		string(TierTypeCommitmentFiveZeroThousandAUHours),
		string(TierTypeCommitmentOneHundredAUHours),
		string(TierTypeCommitmentOneHundredThousandAUHours),
		string(TierTypeCommitmentOneThousandAUHours),
		string(TierTypeCommitmentOneZeroThousandAUHours),
		string(TierTypeConsumption),
	}
}

type Type string

const (
	TypeMicrosoftPointDataLakeAnalyticsAccounts Type = "Microsoft.DataLakeAnalytics/accounts"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeMicrosoftPointDataLakeAnalyticsAccounts),
	}
}

type VirtualNetworkRuleState string

const (
	VirtualNetworkRuleStateActive               VirtualNetworkRuleState = "Active"
	VirtualNetworkRuleStateFailed               VirtualNetworkRuleState = "Failed"
	VirtualNetworkRuleStateNetworkSourceDeleted VirtualNetworkRuleState = "NetworkSourceDeleted"
)

func PossibleValuesForVirtualNetworkRuleState() []string {
	return []string{
		string(VirtualNetworkRuleStateActive),
		string(VirtualNetworkRuleStateFailed),
		string(VirtualNetworkRuleStateNetworkSourceDeleted),
	}
}
