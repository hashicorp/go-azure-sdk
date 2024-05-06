package datastores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatastoreProperties struct {
	CapacityGB         *int64             `json:"capacityGB,omitempty"`
	CustomResourceName *string            `json:"customResourceName,omitempty"`
	FreeSpaceGB        *int64             `json:"freeSpaceGB,omitempty"`
	InventoryItemId    *string            `json:"inventoryItemId,omitempty"`
	MoName             *string            `json:"moName,omitempty"`
	MoRefId            *string            `json:"moRefId,omitempty"`
	ProvisioningState  *ProvisioningState `json:"provisioningState,omitempty"`
	Statuses           *[]ResourceStatus  `json:"statuses,omitempty"`
	Uuid               *string            `json:"uuid,omitempty"`
	VCenterId          *string            `json:"vCenterId,omitempty"`
}
