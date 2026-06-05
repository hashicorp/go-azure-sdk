package networkinterfaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EffectiveNetworkSecurityGroupAssociation struct {
	NetworkInterface *CommonSubResource `json:"networkInterface,omitempty"`
	NetworkManager   *CommonSubResource `json:"networkManager,omitempty"`
	Subnet           *CommonSubResource `json:"subnet,omitempty"`
}
