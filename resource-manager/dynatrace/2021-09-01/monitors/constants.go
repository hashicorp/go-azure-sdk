package monitors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoUpdateSetting string

const (
	AutoUpdateSettingDISABLED AutoUpdateSetting = "DISABLED"
	AutoUpdateSettingENABLED  AutoUpdateSetting = "ENABLED"
)

func PossibleValuesForAutoUpdateSetting() []string {
	return []string{
		string(AutoUpdateSettingDISABLED),
		string(AutoUpdateSettingENABLED),
	}
}

type AvailabilityState string

const (
	AvailabilityStateCRASHED            AvailabilityState = "CRASHED"
	AvailabilityStateLOST               AvailabilityState = "LOST"
	AvailabilityStateMONITORED          AvailabilityState = "MONITORED"
	AvailabilityStatePREMONITORED       AvailabilityState = "PRE_MONITORED"
	AvailabilityStateSHUTDOWN           AvailabilityState = "SHUTDOWN"
	AvailabilityStateUNEXPECTEDSHUTDOWN AvailabilityState = "UNEXPECTED_SHUTDOWN"
	AvailabilityStateUNKNOWN            AvailabilityState = "UNKNOWN"
	AvailabilityStateUNMONITORED        AvailabilityState = "UNMONITORED"
)

func PossibleValuesForAvailabilityState() []string {
	return []string{
		string(AvailabilityStateCRASHED),
		string(AvailabilityStateLOST),
		string(AvailabilityStateMONITORED),
		string(AvailabilityStatePREMONITORED),
		string(AvailabilityStateSHUTDOWN),
		string(AvailabilityStateUNEXPECTEDSHUTDOWN),
		string(AvailabilityStateUNKNOWN),
		string(AvailabilityStateUNMONITORED),
	}
}

type LiftrResourceCategories string

const (
	LiftrResourceCategoriesMonitorLogs LiftrResourceCategories = "MonitorLogs"
	LiftrResourceCategoriesUnknown     LiftrResourceCategories = "Unknown"
)

func PossibleValuesForLiftrResourceCategories() []string {
	return []string{
		string(LiftrResourceCategoriesMonitorLogs),
		string(LiftrResourceCategoriesUnknown),
	}
}

type LogModule string

const (
	LogModuleDISABLED LogModule = "DISABLED"
	LogModuleENABLED  LogModule = "ENABLED"
)

func PossibleValuesForLogModule() []string {
	return []string{
		string(LogModuleDISABLED),
		string(LogModuleENABLED),
	}
}

type ManagedIdentityType string

const (
	ManagedIdentityTypeSystemAndUserAssigned ManagedIdentityType = "SystemAndUserAssigned"
	ManagedIdentityTypeSystemAssigned        ManagedIdentityType = "SystemAssigned"
	ManagedIdentityTypeUserAssigned          ManagedIdentityType = "UserAssigned"
)

func PossibleValuesForManagedIdentityType() []string {
	return []string{
		string(ManagedIdentityTypeSystemAndUserAssigned),
		string(ManagedIdentityTypeSystemAssigned),
		string(ManagedIdentityTypeUserAssigned),
	}
}

type MarketplaceSubscriptionStatus string

const (
	MarketplaceSubscriptionStatusActive    MarketplaceSubscriptionStatus = "Active"
	MarketplaceSubscriptionStatusSuspended MarketplaceSubscriptionStatus = "Suspended"
)

func PossibleValuesForMarketplaceSubscriptionStatus() []string {
	return []string{
		string(MarketplaceSubscriptionStatusActive),
		string(MarketplaceSubscriptionStatusSuspended),
	}
}

type MonitoringStatus string

const (
	MonitoringStatusDisabled MonitoringStatus = "Disabled"
	MonitoringStatusEnabled  MonitoringStatus = "Enabled"
)

func PossibleValuesForMonitoringStatus() []string {
	return []string{
		string(MonitoringStatusDisabled),
		string(MonitoringStatusEnabled),
	}
}

type MonitoringType string

const (
	MonitoringTypeCLOUDINFRASTRUCTURE MonitoringType = "CLOUD_INFRASTRUCTURE"
	MonitoringTypeFULLSTACK           MonitoringType = "FULL_STACK"
)

func PossibleValuesForMonitoringType() []string {
	return []string{
		string(MonitoringTypeCLOUDINFRASTRUCTURE),
		string(MonitoringTypeFULLSTACK),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateNotSpecified),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type SSOStatus string

const (
	SSOStatusDisabled SSOStatus = "Disabled"
	SSOStatusEnabled  SSOStatus = "Enabled"
)

func PossibleValuesForSSOStatus() []string {
	return []string{
		string(SSOStatusDisabled),
		string(SSOStatusEnabled),
	}
}

type SendingLogsStatus string

const (
	SendingLogsStatusDisabled SendingLogsStatus = "Disabled"
	SendingLogsStatusEnabled  SendingLogsStatus = "Enabled"
)

func PossibleValuesForSendingLogsStatus() []string {
	return []string{
		string(SendingLogsStatusDisabled),
		string(SendingLogsStatusEnabled),
	}
}

type SendingMetricsStatus string

const (
	SendingMetricsStatusDisabled SendingMetricsStatus = "Disabled"
	SendingMetricsStatusEnabled  SendingMetricsStatus = "Enabled"
)

func PossibleValuesForSendingMetricsStatus() []string {
	return []string{
		string(SendingMetricsStatusDisabled),
		string(SendingMetricsStatusEnabled),
	}
}

type SingleSignOnStates string

const (
	SingleSignOnStatesDisable  SingleSignOnStates = "Disable"
	SingleSignOnStatesEnable   SingleSignOnStates = "Enable"
	SingleSignOnStatesExisting SingleSignOnStates = "Existing"
	SingleSignOnStatesInitial  SingleSignOnStates = "Initial"
)

func PossibleValuesForSingleSignOnStates() []string {
	return []string{
		string(SingleSignOnStatesDisable),
		string(SingleSignOnStatesEnable),
		string(SingleSignOnStatesExisting),
		string(SingleSignOnStatesInitial),
	}
}

type UpdateStatus string

const (
	UpdateStatusINCOMPATIBLE     UpdateStatus = "INCOMPATIBLE"
	UpdateStatusOUTDATED         UpdateStatus = "OUTDATED"
	UpdateStatusSCHEDULED        UpdateStatus = "SCHEDULED"
	UpdateStatusSUPPRESSED       UpdateStatus = "SUPPRESSED"
	UpdateStatusUNKNOWN          UpdateStatus = "UNKNOWN"
	UpdateStatusUPDATEINPROGRESS UpdateStatus = "UPDATE_IN_PROGRESS"
	UpdateStatusUPDATEPENDING    UpdateStatus = "UPDATE_PENDING"
	UpdateStatusUPDATEPROBLEM    UpdateStatus = "UPDATE_PROBLEM"
	UpdateStatusUPTwoDATE        UpdateStatus = "UP2DATE"
)

func PossibleValuesForUpdateStatus() []string {
	return []string{
		string(UpdateStatusINCOMPATIBLE),
		string(UpdateStatusOUTDATED),
		string(UpdateStatusSCHEDULED),
		string(UpdateStatusSUPPRESSED),
		string(UpdateStatusUNKNOWN),
		string(UpdateStatusUPDATEINPROGRESS),
		string(UpdateStatusUPDATEPENDING),
		string(UpdateStatusUPDATEPROBLEM),
		string(UpdateStatusUPTwoDATE),
	}
}
