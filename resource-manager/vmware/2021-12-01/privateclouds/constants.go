package privateclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailabilityStrategy string

const (
	AvailabilityStrategyDualZone   AvailabilityStrategy = "DualZone"
	AvailabilityStrategySingleZone AvailabilityStrategy = "SingleZone"
)

func PossibleValuesForAvailabilityStrategy() []string {
	return []string{
		string(AvailabilityStrategyDualZone),
		string(AvailabilityStrategySingleZone),
	}
}

type ClusterProvisioningState string

const (
	ClusterProvisioningStateCancelled ClusterProvisioningState = "Cancelled"
	ClusterProvisioningStateDeleting  ClusterProvisioningState = "Deleting"
	ClusterProvisioningStateFailed    ClusterProvisioningState = "Failed"
	ClusterProvisioningStateSucceeded ClusterProvisioningState = "Succeeded"
	ClusterProvisioningStateUpdating  ClusterProvisioningState = "Updating"
)

func PossibleValuesForClusterProvisioningState() []string {
	return []string{
		string(ClusterProvisioningStateCancelled),
		string(ClusterProvisioningStateDeleting),
		string(ClusterProvisioningStateFailed),
		string(ClusterProvisioningStateSucceeded),
		string(ClusterProvisioningStateUpdating),
	}
}

type EncryptionKeyStatus string

const (
	EncryptionKeyStatusAccessDenied EncryptionKeyStatus = "AccessDenied"
	EncryptionKeyStatusConnected    EncryptionKeyStatus = "Connected"
)

func PossibleValuesForEncryptionKeyStatus() []string {
	return []string{
		string(EncryptionKeyStatusAccessDenied),
		string(EncryptionKeyStatusConnected),
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

type EncryptionVersionType string

const (
	EncryptionVersionTypeAutoDetected EncryptionVersionType = "AutoDetected"
	EncryptionVersionTypeFixed        EncryptionVersionType = "Fixed"
)

func PossibleValuesForEncryptionVersionType() []string {
	return []string{
		string(EncryptionVersionTypeAutoDetected),
		string(EncryptionVersionTypeFixed),
	}
}

type InternetEnum string

const (
	InternetEnumDisabled InternetEnum = "Disabled"
	InternetEnumEnabled  InternetEnum = "Enabled"
)

func PossibleValuesForInternetEnum() []string {
	return []string{
		string(InternetEnumDisabled),
		string(InternetEnumEnabled),
	}
}

type PrivateCloudProvisioningState string

const (
	PrivateCloudProvisioningStateBuilding  PrivateCloudProvisioningState = "Building"
	PrivateCloudProvisioningStateCancelled PrivateCloudProvisioningState = "Cancelled"
	PrivateCloudProvisioningStateDeleting  PrivateCloudProvisioningState = "Deleting"
	PrivateCloudProvisioningStateFailed    PrivateCloudProvisioningState = "Failed"
	PrivateCloudProvisioningStatePending   PrivateCloudProvisioningState = "Pending"
	PrivateCloudProvisioningStateSucceeded PrivateCloudProvisioningState = "Succeeded"
	PrivateCloudProvisioningStateUpdating  PrivateCloudProvisioningState = "Updating"
)

func PossibleValuesForPrivateCloudProvisioningState() []string {
	return []string{
		string(PrivateCloudProvisioningStateBuilding),
		string(PrivateCloudProvisioningStateCancelled),
		string(PrivateCloudProvisioningStateDeleting),
		string(PrivateCloudProvisioningStateFailed),
		string(PrivateCloudProvisioningStatePending),
		string(PrivateCloudProvisioningStateSucceeded),
		string(PrivateCloudProvisioningStateUpdating),
	}
}

type SslEnum string

const (
	SslEnumDisabled SslEnum = "Disabled"
	SslEnumEnabled  SslEnum = "Enabled"
)

func PossibleValuesForSslEnum() []string {
	return []string{
		string(SslEnumDisabled),
		string(SslEnumEnabled),
	}
}
