package updates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpactType string

const (
	ImpactTypeFreeze   ImpactType = "Freeze"
	ImpactTypeNone     ImpactType = "None"
	ImpactTypeRedeploy ImpactType = "Redeploy"
	ImpactTypeRestart  ImpactType = "Restart"
)

func PossibleValuesForImpactType() []string {
	return []string{
		string(ImpactTypeFreeze),
		string(ImpactTypeNone),
		string(ImpactTypeRedeploy),
		string(ImpactTypeRestart),
	}
}

type MaintenanceScope string

const (
	MaintenanceScopeExtension          MaintenanceScope = "Extension"
	MaintenanceScopeHost               MaintenanceScope = "Host"
	MaintenanceScopeInGuestPatch       MaintenanceScope = "InGuestPatch"
	MaintenanceScopeOSImage            MaintenanceScope = "OSImage"
	MaintenanceScopeSQLDB              MaintenanceScope = "SQLDB"
	MaintenanceScopeSQLManagedInstance MaintenanceScope = "SQLManagedInstance"
)

func PossibleValuesForMaintenanceScope() []string {
	return []string{
		string(MaintenanceScopeExtension),
		string(MaintenanceScopeHost),
		string(MaintenanceScopeInGuestPatch),
		string(MaintenanceScopeOSImage),
		string(MaintenanceScopeSQLDB),
		string(MaintenanceScopeSQLManagedInstance),
	}
}

type UpdateStatus string

const (
	UpdateStatusCompleted  UpdateStatus = "Completed"
	UpdateStatusInProgress UpdateStatus = "InProgress"
	UpdateStatusPending    UpdateStatus = "Pending"
	UpdateStatusRetryLater UpdateStatus = "RetryLater"
	UpdateStatusRetryNow   UpdateStatus = "RetryNow"
)

func PossibleValuesForUpdateStatus() []string {
	return []string{
		string(UpdateStatusCompleted),
		string(UpdateStatusInProgress),
		string(UpdateStatusPending),
		string(UpdateStatusRetryLater),
		string(UpdateStatusRetryNow),
	}
}
