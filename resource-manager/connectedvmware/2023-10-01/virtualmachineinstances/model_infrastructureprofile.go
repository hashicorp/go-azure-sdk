package virtualmachineinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InfrastructureProfile struct {
	CustomResourceName *string       `json:"customResourceName,omitempty"`
	FirmwareType       *FirmwareType `json:"firmwareType,omitempty"`
	FolderPath         *string       `json:"folderPath,omitempty"`
	InstanceUuid       *string       `json:"instanceUuid,omitempty"`
	InventoryItemId    *string       `json:"inventoryItemId,omitempty"`
	MoName             *string       `json:"moName,omitempty"`
	MoRefId            *string       `json:"moRefId,omitempty"`
	SmbiosUuid         *string       `json:"smbiosUuid,omitempty"`
	TemplateId         *string       `json:"templateId,omitempty"`
	VCenterId          *string       `json:"vCenterId,omitempty"`
}
