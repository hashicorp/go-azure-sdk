package securityconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderForServersGcpOfferingArcAutoProvisioningConfiguration struct {
	PrivateLinkScope *string `json:"privateLinkScope,omitempty"`
	Proxy            *string `json:"proxy,omitempty"`
}
