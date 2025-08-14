package serviceresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMigrationServiceProperties struct {
	AutoStopDelay         *string                   `json:"autoStopDelay,omitempty"`
	DeleteResourcesOnStop *bool                     `json:"deleteResourcesOnStop,omitempty"`
	ProvisioningState     *ServiceProvisioningState `json:"provisioningState,omitempty"`
	PublicKey             *string                   `json:"publicKey,omitempty"`
	VirtualNicId          *string                   `json:"virtualNicId,omitempty"`
	VirtualSubnetId       *string                   `json:"virtualSubnetId,omitempty"`
}
