package virtualmachine

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineConnectionProfile struct {
	AdminUsername    *string `json:"adminUsername,omitempty"`
	NonAdminUsername *string `json:"nonAdminUsername,omitempty"`
	PrivateIPAddress *string `json:"privateIpAddress,omitempty"`
	RdpAuthority     *string `json:"rdpAuthority,omitempty"`
	RdpInBrowserURL  *string `json:"rdpInBrowserUrl,omitempty"`
	SshAuthority     *string `json:"sshAuthority,omitempty"`
	SshInBrowserURL  *string `json:"sshInBrowserUrl,omitempty"`
}
