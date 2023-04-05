package nodetype

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskType string

const (
	DiskTypePremiumLRS     DiskType = "Premium_LRS"
	DiskTypeStandardLRS    DiskType = "Standard_LRS"
	DiskTypeStandardSSDLRS DiskType = "StandardSSD_LRS"
)

func PossibleValuesForDiskType() []string {
	return []string{
		string(DiskTypePremiumLRS),
		string(DiskTypeStandardLRS),
		string(DiskTypeStandardSSDLRS),
	}
}

type ManagedResourceProvisioningState string

const (
	ManagedResourceProvisioningStateCanceled  ManagedResourceProvisioningState = "Canceled"
	ManagedResourceProvisioningStateCreated   ManagedResourceProvisioningState = "Created"
	ManagedResourceProvisioningStateCreating  ManagedResourceProvisioningState = "Creating"
	ManagedResourceProvisioningStateDeleted   ManagedResourceProvisioningState = "Deleted"
	ManagedResourceProvisioningStateDeleting  ManagedResourceProvisioningState = "Deleting"
	ManagedResourceProvisioningStateFailed    ManagedResourceProvisioningState = "Failed"
	ManagedResourceProvisioningStateNone      ManagedResourceProvisioningState = "None"
	ManagedResourceProvisioningStateOther     ManagedResourceProvisioningState = "Other"
	ManagedResourceProvisioningStateSucceeded ManagedResourceProvisioningState = "Succeeded"
	ManagedResourceProvisioningStateUpdating  ManagedResourceProvisioningState = "Updating"
)

func PossibleValuesForManagedResourceProvisioningState() []string {
	return []string{
		string(ManagedResourceProvisioningStateCanceled),
		string(ManagedResourceProvisioningStateCreated),
		string(ManagedResourceProvisioningStateCreating),
		string(ManagedResourceProvisioningStateDeleted),
		string(ManagedResourceProvisioningStateDeleting),
		string(ManagedResourceProvisioningStateFailed),
		string(ManagedResourceProvisioningStateNone),
		string(ManagedResourceProvisioningStateOther),
		string(ManagedResourceProvisioningStateSucceeded),
		string(ManagedResourceProvisioningStateUpdating),
	}
}
