package replicationprotectionintents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type A2ARecoveryAvailabilityType string

const (
	A2ARecoveryAvailabilityTypeAvailabilitySet  A2ARecoveryAvailabilityType = "AvailabilitySet"
	A2ARecoveryAvailabilityTypeAvailabilityZone A2ARecoveryAvailabilityType = "AvailabilityZone"
	A2ARecoveryAvailabilityTypeSingle           A2ARecoveryAvailabilityType = "Single"
)

func PossibleValuesForA2ARecoveryAvailabilityType() []string {
	return []string{
		string(A2ARecoveryAvailabilityTypeAvailabilitySet),
		string(A2ARecoveryAvailabilityTypeAvailabilityZone),
		string(A2ARecoveryAvailabilityTypeSingle),
	}
}

type AgentAutoUpdateStatus string

const (
	AgentAutoUpdateStatusDisabled AgentAutoUpdateStatus = "Disabled"
	AgentAutoUpdateStatusEnabled  AgentAutoUpdateStatus = "Enabled"
)

func PossibleValuesForAgentAutoUpdateStatus() []string {
	return []string{
		string(AgentAutoUpdateStatusDisabled),
		string(AgentAutoUpdateStatusEnabled),
	}
}

type AutoProtectionOfDataDisk string

const (
	AutoProtectionOfDataDiskDisabled AutoProtectionOfDataDisk = "Disabled"
	AutoProtectionOfDataDiskEnabled  AutoProtectionOfDataDisk = "Enabled"
)

func PossibleValuesForAutoProtectionOfDataDisk() []string {
	return []string{
		string(AutoProtectionOfDataDiskDisabled),
		string(AutoProtectionOfDataDiskEnabled),
	}
}

type AutomationAccountAuthenticationType string

const (
	AutomationAccountAuthenticationTypeRunAsAccount           AutomationAccountAuthenticationType = "RunAsAccount"
	AutomationAccountAuthenticationTypeSystemAssignedIdentity AutomationAccountAuthenticationType = "SystemAssignedIdentity"
)

func PossibleValuesForAutomationAccountAuthenticationType() []string {
	return []string{
		string(AutomationAccountAuthenticationTypeRunAsAccount),
		string(AutomationAccountAuthenticationTypeSystemAssignedIdentity),
	}
}

type SetMultiVMSyncStatus string

const (
	SetMultiVMSyncStatusDisable SetMultiVMSyncStatus = "Disable"
	SetMultiVMSyncStatusEnable  SetMultiVMSyncStatus = "Enable"
)

func PossibleValuesForSetMultiVMSyncStatus() []string {
	return []string{
		string(SetMultiVMSyncStatusDisable),
		string(SetMultiVMSyncStatusEnable),
	}
}
