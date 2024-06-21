package customimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomImageProperties struct {
	Author              *string                        `json:"author,omitempty"`
	CreationDate        *string                        `json:"creationDate,omitempty"`
	CustomImagePlan     *CustomImagePropertiesFromPlan `json:"customImagePlan,omitempty"`
	DataDiskStorageInfo *[]DataDiskStorageTypeInfo     `json:"dataDiskStorageInfo,omitempty"`
	Description         *string                        `json:"description,omitempty"`
	IsPlanAuthorized    *bool                          `json:"isPlanAuthorized,omitempty"`
	ManagedImageId      *string                        `json:"managedImageId,omitempty"`
	ManagedSnapshotId   *string                        `json:"managedSnapshotId,omitempty"`
	ProvisioningState   *string                        `json:"provisioningState,omitempty"`
	UniqueIdentifier    *string                        `json:"uniqueIdentifier,omitempty"`
	VM                  *CustomImagePropertiesFromVM   `json:"vm,omitempty"`
	Vhd                 *CustomImagePropertiesCustom   `json:"vhd,omitempty"`
}
