package privateendpointconnection

import (
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionStatus string

const (
	PrivateEndpointConnectionStatusApproved     PrivateEndpointConnectionStatus = "Approved"
	PrivateEndpointConnectionStatusDisconnected PrivateEndpointConnectionStatus = "Disconnected"
	PrivateEndpointConnectionStatusPending      PrivateEndpointConnectionStatus = "Pending"
	PrivateEndpointConnectionStatusRejected     PrivateEndpointConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointConnectionStatus() []string {
	return []string{
		string(PrivateEndpointConnectionStatusApproved),
		string(PrivateEndpointConnectionStatusDisconnected),
		string(PrivateEndpointConnectionStatusPending),
		string(PrivateEndpointConnectionStatusRejected),
	}
}

func parsePrivateEndpointConnectionStatus(input string) (*PrivateEndpointConnectionStatus, error) {
	vals := map[string]PrivateEndpointConnectionStatus{
		"approved":     PrivateEndpointConnectionStatusApproved,
		"disconnected": PrivateEndpointConnectionStatusDisconnected,
		"pending":      PrivateEndpointConnectionStatusPending,
		"rejected":     PrivateEndpointConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointConnectionStatus(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStatePending   ProvisioningState = "Pending"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStatePending),
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"pending":   ProvisioningStatePending,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type VaultSubResourceType string

const (
	VaultSubResourceTypeAzureBackup          VaultSubResourceType = "AzureBackup"
	VaultSubResourceTypeAzureBackupSecondary VaultSubResourceType = "AzureBackup_secondary"
	VaultSubResourceTypeAzureSiteRecovery    VaultSubResourceType = "AzureSiteRecovery"
)

func PossibleValuesForVaultSubResourceType() []string {
	return []string{
		string(VaultSubResourceTypeAzureBackup),
		string(VaultSubResourceTypeAzureBackupSecondary),
		string(VaultSubResourceTypeAzureSiteRecovery),
	}
}

func parseVaultSubResourceType(input string) (*VaultSubResourceType, error) {
	vals := map[string]VaultSubResourceType{
		"azurebackup":           VaultSubResourceTypeAzureBackup,
		"azurebackup_secondary": VaultSubResourceTypeAzureBackupSecondary,
		"azuresiterecovery":     VaultSubResourceTypeAzureSiteRecovery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VaultSubResourceType(input)
	return &out, nil
}
