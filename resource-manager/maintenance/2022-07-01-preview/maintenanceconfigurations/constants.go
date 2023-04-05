package maintenanceconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceScope string

const (
	MaintenanceScopeExtension          MaintenanceScope = "Extension"
	MaintenanceScopeHost               MaintenanceScope = "Host"
	MaintenanceScopeInGuestPatch       MaintenanceScope = "InGuestPatch"
	MaintenanceScopeOSImage            MaintenanceScope = "OSImage"
	MaintenanceScopeResource           MaintenanceScope = "Resource"
	MaintenanceScopeSQLDB              MaintenanceScope = "SQLDB"
	MaintenanceScopeSQLManagedInstance MaintenanceScope = "SQLManagedInstance"
)

func PossibleValuesForMaintenanceScope() []string {
	return []string{
		string(MaintenanceScopeExtension),
		string(MaintenanceScopeHost),
		string(MaintenanceScopeInGuestPatch),
		string(MaintenanceScopeOSImage),
		string(MaintenanceScopeResource),
		string(MaintenanceScopeSQLDB),
		string(MaintenanceScopeSQLManagedInstance),
	}
}

type RebootOptions string

const (
	RebootOptionsAlways     RebootOptions = "Always"
	RebootOptionsIfRequired RebootOptions = "IfRequired"
	RebootOptionsNever      RebootOptions = "Never"
)

func PossibleValuesForRebootOptions() []string {
	return []string{
		string(RebootOptionsAlways),
		string(RebootOptionsIfRequired),
		string(RebootOptionsNever),
	}
}

type TaskScope string

const (
	TaskScopeGlobal   TaskScope = "Global"
	TaskScopeResource TaskScope = "Resource"
)

func PossibleValuesForTaskScope() []string {
	return []string{
		string(TaskScopeGlobal),
		string(TaskScopeResource),
	}
}

type Visibility string

const (
	VisibilityCustom Visibility = "Custom"
	VisibilityPublic Visibility = "Public"
)

func PossibleValuesForVisibility() []string {
	return []string{
		string(VisibilityCustom),
		string(VisibilityPublic),
	}
}
