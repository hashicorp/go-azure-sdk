package subaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type ManagedIdentityTypes string

const (
	ManagedIdentityTypesSystemAssigned ManagedIdentityTypes = "SystemAssigned"
	ManagedIdentityTypesUserAssigned   ManagedIdentityTypes = "UserAssigned"
)

func PossibleValuesForManagedIdentityTypes() []string {
	return []string{
		string(ManagedIdentityTypesSystemAssigned),
		string(ManagedIdentityTypesUserAssigned),
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
