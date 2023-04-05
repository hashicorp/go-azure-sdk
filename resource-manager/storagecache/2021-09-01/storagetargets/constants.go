package storagetargets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationalStateType string

const (
	OperationalStateTypeBusy      OperationalStateType = "Busy"
	OperationalStateTypeFlushing  OperationalStateType = "Flushing"
	OperationalStateTypeReady     OperationalStateType = "Ready"
	OperationalStateTypeSuspended OperationalStateType = "Suspended"
)

func PossibleValuesForOperationalStateType() []string {
	return []string{
		string(OperationalStateTypeBusy),
		string(OperationalStateTypeFlushing),
		string(OperationalStateTypeReady),
		string(OperationalStateTypeSuspended),
	}
}

type ProvisioningStateType string

const (
	ProvisioningStateTypeCancelled ProvisioningStateType = "Cancelled"
	ProvisioningStateTypeCreating  ProvisioningStateType = "Creating"
	ProvisioningStateTypeDeleting  ProvisioningStateType = "Deleting"
	ProvisioningStateTypeFailed    ProvisioningStateType = "Failed"
	ProvisioningStateTypeSucceeded ProvisioningStateType = "Succeeded"
	ProvisioningStateTypeUpdating  ProvisioningStateType = "Updating"
)

func PossibleValuesForProvisioningStateType() []string {
	return []string{
		string(ProvisioningStateTypeCancelled),
		string(ProvisioningStateTypeCreating),
		string(ProvisioningStateTypeDeleting),
		string(ProvisioningStateTypeFailed),
		string(ProvisioningStateTypeSucceeded),
		string(ProvisioningStateTypeUpdating),
	}
}

type StorageTargetType string

const (
	StorageTargetTypeBlobNfs  StorageTargetType = "blobNfs"
	StorageTargetTypeClfs     StorageTargetType = "clfs"
	StorageTargetTypeNfsThree StorageTargetType = "nfs3"
	StorageTargetTypeUnknown  StorageTargetType = "unknown"
)

func PossibleValuesForStorageTargetType() []string {
	return []string{
		string(StorageTargetTypeBlobNfs),
		string(StorageTargetTypeClfs),
		string(StorageTargetTypeNfsThree),
		string(StorageTargetTypeUnknown),
	}
}
