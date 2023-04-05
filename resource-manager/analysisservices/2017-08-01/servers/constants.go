package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionMode string

const (
	ConnectionModeAll      ConnectionMode = "All"
	ConnectionModeReadOnly ConnectionMode = "ReadOnly"
)

func PossibleValuesForConnectionMode() []string {
	return []string{
		string(ConnectionModeAll),
		string(ConnectionModeReadOnly),
	}
}

type ManagedMode int64

const (
	ManagedModeOne  ManagedMode = 1
	ManagedModeZero ManagedMode = 0
)

func PossibleValuesForManagedMode() []int64 {
	return []int64{
		int64(ManagedModeOne),
		int64(ManagedModeZero),
	}
}

type ProvisioningState string

const (
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStatePaused       ProvisioningState = "Paused"
	ProvisioningStatePausing      ProvisioningState = "Pausing"
	ProvisioningStatePreparing    ProvisioningState = "Preparing"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateResuming     ProvisioningState = "Resuming"
	ProvisioningStateScaling      ProvisioningState = "Scaling"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateSuspended    ProvisioningState = "Suspended"
	ProvisioningStateSuspending   ProvisioningState = "Suspending"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStatePaused),
		string(ProvisioningStatePausing),
		string(ProvisioningStatePreparing),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateResuming),
		string(ProvisioningStateScaling),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateSuspended),
		string(ProvisioningStateSuspending),
		string(ProvisioningStateUpdating),
	}
}

type ServerMonitorMode int64

const (
	ServerMonitorModeOne  ServerMonitorMode = 1
	ServerMonitorModeZero ServerMonitorMode = 0
)

func PossibleValuesForServerMonitorMode() []int64 {
	return []int64{
		int64(ServerMonitorModeOne),
		int64(ServerMonitorModeZero),
	}
}

type SkuTier string

const (
	SkuTierBasic       SkuTier = "Basic"
	SkuTierDevelopment SkuTier = "Development"
	SkuTierStandard    SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierDevelopment),
		string(SkuTierStandard),
	}
}

type State string

const (
	StateDeleting     State = "Deleting"
	StateFailed       State = "Failed"
	StatePaused       State = "Paused"
	StatePausing      State = "Pausing"
	StatePreparing    State = "Preparing"
	StateProvisioning State = "Provisioning"
	StateResuming     State = "Resuming"
	StateScaling      State = "Scaling"
	StateSucceeded    State = "Succeeded"
	StateSuspended    State = "Suspended"
	StateSuspending   State = "Suspending"
	StateUpdating     State = "Updating"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateDeleting),
		string(StateFailed),
		string(StatePaused),
		string(StatePausing),
		string(StatePreparing),
		string(StateProvisioning),
		string(StateResuming),
		string(StateScaling),
		string(StateSucceeded),
		string(StateSuspended),
		string(StateSuspending),
		string(StateUpdating),
	}
}

type Status int64

const (
	StatusZero Status = 0
)

func PossibleValuesForStatus() []int64 {
	return []int64{
		int64(StatusZero),
	}
}
