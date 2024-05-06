package vcenters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VCenterProperties struct {
	ConnectionStatus   *string            `json:"connectionStatus,omitempty"`
	Credentials        *VICredential      `json:"credentials,omitempty"`
	CustomResourceName *string            `json:"customResourceName,omitempty"`
	Fqdn               string             `json:"fqdn"`
	InstanceUuid       *string            `json:"instanceUuid,omitempty"`
	Port               *int64             `json:"port,omitempty"`
	ProvisioningState  *ProvisioningState `json:"provisioningState,omitempty"`
	Statuses           *[]ResourceStatus  `json:"statuses,omitempty"`
	Uuid               *string            `json:"uuid,omitempty"`
	Version            *string            `json:"version,omitempty"`
}
