package replicationjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FailoverReplicationProtectedItemDetails struct {
	FriendlyName            *string `json:"friendlyName,omitempty"`
	Name                    *string `json:"name,omitempty"`
	NetworkConnectionStatus *string `json:"networkConnectionStatus,omitempty"`
	NetworkFriendlyName     *string `json:"networkFriendlyName,omitempty"`
	RecoveryPointId         *string `json:"recoveryPointId,omitempty"`
	RecoveryPointTime       *string `json:"recoveryPointTime,omitempty"`
	Subnet                  *string `json:"subnet,omitempty"`
	TestVMFriendlyName      *string `json:"testVmFriendlyName,omitempty"`
	TestVMName              *string `json:"testVmName,omitempty"`
}
