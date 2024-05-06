package virtualmachinetemplates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineTemplateProperties struct {
	CustomResourceName *string             `json:"customResourceName,omitempty"`
	Disks              *[]VirtualDisk      `json:"disks,omitempty"`
	FirmwareType       *FirmwareType       `json:"firmwareType,omitempty"`
	FolderPath         *string             `json:"folderPath,omitempty"`
	InventoryItemId    *string             `json:"inventoryItemId,omitempty"`
	MemorySizeMB       *int64              `json:"memorySizeMB,omitempty"`
	MoName             *string             `json:"moName,omitempty"`
	MoRefId            *string             `json:"moRefId,omitempty"`
	NetworkInterfaces  *[]NetworkInterface `json:"networkInterfaces,omitempty"`
	NumCPUs            *int64              `json:"numCPUs,omitempty"`
	NumCoresPerSocket  *int64              `json:"numCoresPerSocket,omitempty"`
	OsName             *string             `json:"osName,omitempty"`
	OsType             *OsType             `json:"osType,omitempty"`
	ProvisioningState  *ProvisioningState  `json:"provisioningState,omitempty"`
	Statuses           *[]ResourceStatus   `json:"statuses,omitempty"`
	ToolsVersion       *string             `json:"toolsVersion,omitempty"`
	ToolsVersionStatus *string             `json:"toolsVersionStatus,omitempty"`
	Uuid               *string             `json:"uuid,omitempty"`
	VCenterId          *string             `json:"vCenterId,omitempty"`
}
