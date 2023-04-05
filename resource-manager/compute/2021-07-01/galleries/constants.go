package galleries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GallerySharingPermissionTypes string

const (
	GallerySharingPermissionTypesGroups  GallerySharingPermissionTypes = "Groups"
	GallerySharingPermissionTypesPrivate GallerySharingPermissionTypes = "Private"
)

func PossibleValuesForGallerySharingPermissionTypes() []string {
	return []string{
		string(GallerySharingPermissionTypesGroups),
		string(GallerySharingPermissionTypesPrivate),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMigrating ProvisioningState = "Migrating"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMigrating),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type SelectPermissions string

const (
	SelectPermissionsPermissions SelectPermissions = "Permissions"
)

func PossibleValuesForSelectPermissions() []string {
	return []string{
		string(SelectPermissionsPermissions),
	}
}

type SharingProfileGroupTypes string

const (
	SharingProfileGroupTypesAADTenants    SharingProfileGroupTypes = "AADTenants"
	SharingProfileGroupTypesSubscriptions SharingProfileGroupTypes = "Subscriptions"
)

func PossibleValuesForSharingProfileGroupTypes() []string {
	return []string{
		string(SharingProfileGroupTypesAADTenants),
		string(SharingProfileGroupTypesSubscriptions),
	}
}
