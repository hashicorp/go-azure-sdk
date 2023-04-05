package globalreachconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalReachConnectionProvisioningState string

const (
	GlobalReachConnectionProvisioningStateCanceled  GlobalReachConnectionProvisioningState = "Canceled"
	GlobalReachConnectionProvisioningStateFailed    GlobalReachConnectionProvisioningState = "Failed"
	GlobalReachConnectionProvisioningStateSucceeded GlobalReachConnectionProvisioningState = "Succeeded"
	GlobalReachConnectionProvisioningStateUpdating  GlobalReachConnectionProvisioningState = "Updating"
)

func PossibleValuesForGlobalReachConnectionProvisioningState() []string {
	return []string{
		string(GlobalReachConnectionProvisioningStateCanceled),
		string(GlobalReachConnectionProvisioningStateFailed),
		string(GlobalReachConnectionProvisioningStateSucceeded),
		string(GlobalReachConnectionProvisioningStateUpdating),
	}
}

type GlobalReachConnectionStatus string

const (
	GlobalReachConnectionStatusConnected    GlobalReachConnectionStatus = "Connected"
	GlobalReachConnectionStatusConnecting   GlobalReachConnectionStatus = "Connecting"
	GlobalReachConnectionStatusDisconnected GlobalReachConnectionStatus = "Disconnected"
)

func PossibleValuesForGlobalReachConnectionStatus() []string {
	return []string{
		string(GlobalReachConnectionStatusConnected),
		string(GlobalReachConnectionStatusConnecting),
		string(GlobalReachConnectionStatusDisconnected),
	}
}
