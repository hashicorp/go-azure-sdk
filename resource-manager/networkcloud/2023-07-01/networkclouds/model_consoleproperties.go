package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsoleProperties struct {
	DetailedStatus         *ConsoleDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage  *string                   `json:"detailedStatusMessage,omitempty"`
	Enabled                ConsoleEnabled            `json:"enabled"`
	Expiration             *string                   `json:"expiration,omitempty"`
	PrivateLinkServiceId   *string                   `json:"privateLinkServiceId,omitempty"`
	ProvisioningState      *ConsoleProvisioningState `json:"provisioningState,omitempty"`
	SshPublicKey           SshPublicKey              `json:"sshPublicKey"`
	VirtualMachineAccessId *string                   `json:"virtualMachineAccessId,omitempty"`
}
