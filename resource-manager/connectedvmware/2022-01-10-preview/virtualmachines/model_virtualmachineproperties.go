package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineProperties struct {
	CustomResourceName *string            `json:"customResourceName,omitempty"`
	FirmwareType       *FirmwareType      `json:"firmwareType,omitempty"`
	FolderPath         *string            `json:"folderPath,omitempty"`
	GuestAgentProfile  *GuestAgentProfile `json:"guestAgentProfile,omitempty"`
	HardwareProfile    *HardwareProfile   `json:"hardwareProfile,omitempty"`
	InstanceUuid       *string            `json:"instanceUuid,omitempty"`
	InventoryItemId    *string            `json:"inventoryItemId,omitempty"`
	MoName             *string            `json:"moName,omitempty"`
	MoRefId            *string            `json:"moRefId,omitempty"`
	NetworkProfile     *NetworkProfile    `json:"networkProfile,omitempty"`
	OsProfile          *OsProfile         `json:"osProfile,omitempty"`
	PlacementProfile   *PlacementProfile  `json:"placementProfile,omitempty"`
	PowerState         *string            `json:"powerState,omitempty"`
	ProvisioningState  *string            `json:"provisioningState,omitempty"`
	ResourcePoolId     *string            `json:"resourcePoolId,omitempty"`
	SecurityProfile    *SecurityProfile   `json:"securityProfile,omitempty"`
	SmbiosUuid         *string            `json:"smbiosUuid,omitempty"`
	Statuses           *[]ResourceStatus  `json:"statuses,omitempty"`
	StorageProfile     *StorageProfile    `json:"storageProfile,omitempty"`
	TemplateId         *string            `json:"templateId,omitempty"`
	Uuid               *string            `json:"uuid,omitempty"`
	VCenterId          *string            `json:"vCenterId,omitempty"`
	VmId               *string            `json:"vmId,omitempty"`
}
