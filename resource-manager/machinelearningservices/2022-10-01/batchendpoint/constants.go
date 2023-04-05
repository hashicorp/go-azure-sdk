package batchendpoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointAuthMode string

const (
	EndpointAuthModeAADToken EndpointAuthMode = "AADToken"
	EndpointAuthModeAMLToken EndpointAuthMode = "AMLToken"
	EndpointAuthModeKey      EndpointAuthMode = "Key"
)

func PossibleValuesForEndpointAuthMode() []string {
	return []string{
		string(EndpointAuthModeAADToken),
		string(EndpointAuthModeAMLToken),
		string(EndpointAuthModeKey),
	}
}

type EndpointProvisioningState string

const (
	EndpointProvisioningStateCanceled  EndpointProvisioningState = "Canceled"
	EndpointProvisioningStateCreating  EndpointProvisioningState = "Creating"
	EndpointProvisioningStateDeleting  EndpointProvisioningState = "Deleting"
	EndpointProvisioningStateFailed    EndpointProvisioningState = "Failed"
	EndpointProvisioningStateSucceeded EndpointProvisioningState = "Succeeded"
	EndpointProvisioningStateUpdating  EndpointProvisioningState = "Updating"
)

func PossibleValuesForEndpointProvisioningState() []string {
	return []string{
		string(EndpointProvisioningStateCanceled),
		string(EndpointProvisioningStateCreating),
		string(EndpointProvisioningStateDeleting),
		string(EndpointProvisioningStateFailed),
		string(EndpointProvisioningStateSucceeded),
		string(EndpointProvisioningStateUpdating),
	}
}

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

func PossibleValuesForManagedServiceIdentityType() []string {
	return []string{
		string(ManagedServiceIdentityTypeNone),
		string(ManagedServiceIdentityTypeSystemAssigned),
		string(ManagedServiceIdentityTypeSystemAssignedUserAssigned),
		string(ManagedServiceIdentityTypeUserAssigned),
	}
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierFree     SkuTier = "Free"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}
