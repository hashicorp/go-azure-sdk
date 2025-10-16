package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumeProperties struct {
	AllocatedSizeMiB      *int64                   `json:"allocatedSizeMiB,omitempty"`
	AttachedTo            *[]string                `json:"attachedTo,omitempty"`
	DetailedStatus        *VolumeDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage *string                  `json:"detailedStatusMessage,omitempty"`
	ProvisioningState     *VolumeProvisioningState `json:"provisioningState,omitempty"`
	SerialNumber          *string                  `json:"serialNumber,omitempty"`
	SizeMiB               int64                    `json:"sizeMiB"`
	StorageApplianceId    *string                  `json:"storageApplianceId,omitempty"`
}
