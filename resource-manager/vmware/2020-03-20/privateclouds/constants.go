package privateclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
