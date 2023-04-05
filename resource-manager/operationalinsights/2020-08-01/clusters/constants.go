package clusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterEntityStatus string

const (
	ClusterEntityStatusCanceled            ClusterEntityStatus = "Canceled"
	ClusterEntityStatusCreating            ClusterEntityStatus = "Creating"
	ClusterEntityStatusDeleting            ClusterEntityStatus = "Deleting"
	ClusterEntityStatusFailed              ClusterEntityStatus = "Failed"
	ClusterEntityStatusProvisioningAccount ClusterEntityStatus = "ProvisioningAccount"
	ClusterEntityStatusSucceeded           ClusterEntityStatus = "Succeeded"
	ClusterEntityStatusUpdating            ClusterEntityStatus = "Updating"
)

func PossibleValuesForClusterEntityStatus() []string {
	return []string{
		string(ClusterEntityStatusCanceled),
		string(ClusterEntityStatusCreating),
		string(ClusterEntityStatusDeleting),
		string(ClusterEntityStatusFailed),
		string(ClusterEntityStatusProvisioningAccount),
		string(ClusterEntityStatusSucceeded),
		string(ClusterEntityStatusUpdating),
	}
}

type ClusterSkuNameEnum string

const (
	ClusterSkuNameEnumCapacityReservation ClusterSkuNameEnum = "CapacityReservation"
)

func PossibleValuesForClusterSkuNameEnum() []string {
	return []string{
		string(ClusterSkuNameEnumCapacityReservation),
	}
}
