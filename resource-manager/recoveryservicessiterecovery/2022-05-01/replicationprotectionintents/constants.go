package replicationprotectionintents

import "strings"

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

func parseA2ARecoveryAvailabilityType(input string) (*A2ARecoveryAvailabilityType, error) {
	vals := map[string]A2ARecoveryAvailabilityType{
		"availabilityset":  A2ARecoveryAvailabilityTypeAvailabilitySet,
		"availabilityzone": A2ARecoveryAvailabilityTypeAvailabilityZone,
		"single":           A2ARecoveryAvailabilityTypeSingle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := A2ARecoveryAvailabilityType(input)
	return &out, nil
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

func parseAgentAutoUpdateStatus(input string) (*AgentAutoUpdateStatus, error) {
	vals := map[string]AgentAutoUpdateStatus{
		"disabled": AgentAutoUpdateStatusDisabled,
		"enabled":  AgentAutoUpdateStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentAutoUpdateStatus(input)
	return &out, nil
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

func parseAutoProtectionOfDataDisk(input string) (*AutoProtectionOfDataDisk, error) {
	vals := map[string]AutoProtectionOfDataDisk{
		"disabled": AutoProtectionOfDataDiskDisabled,
		"enabled":  AutoProtectionOfDataDiskEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoProtectionOfDataDisk(input)
	return &out, nil
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

func parseAutomationAccountAuthenticationType(input string) (*AutomationAccountAuthenticationType, error) {
	vals := map[string]AutomationAccountAuthenticationType{
		"runasaccount":           AutomationAccountAuthenticationTypeRunAsAccount,
		"systemassignedidentity": AutomationAccountAuthenticationTypeSystemAssignedIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomationAccountAuthenticationType(input)
	return &out, nil
}

type SetMultiVmSyncStatus string

const (
	SetMultiVmSyncStatusDisable SetMultiVmSyncStatus = "Disable"
	SetMultiVmSyncStatusEnable  SetMultiVmSyncStatus = "Enable"
)

func PossibleValuesForSetMultiVmSyncStatus() []string {
	return []string{
		string(SetMultiVmSyncStatusDisable),
		string(SetMultiVmSyncStatusEnable),
	}
}

func parseSetMultiVmSyncStatus(input string) (*SetMultiVmSyncStatus, error) {
	vals := map[string]SetMultiVmSyncStatus{
		"disable": SetMultiVmSyncStatusDisable,
		"enable":  SetMultiVmSyncStatusEnable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SetMultiVmSyncStatus(input)
	return &out, nil
}
