package hosts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostProperties struct {
	CustomResourceName *string           `json:"customResourceName,omitempty"`
	InventoryItemId    *string           `json:"inventoryItemId,omitempty"`
	MoName             *string           `json:"moName,omitempty"`
	MoRefId            *string           `json:"moRefId,omitempty"`
	ProvisioningState  *string           `json:"provisioningState,omitempty"`
	Statuses           *[]ResourceStatus `json:"statuses,omitempty"`
	Uuid               *string           `json:"uuid,omitempty"`
	VCenterId          *string           `json:"vCenterId,omitempty"`
}
