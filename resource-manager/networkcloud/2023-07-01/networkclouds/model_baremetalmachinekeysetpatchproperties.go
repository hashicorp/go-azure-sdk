package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineKeySetPatchProperties struct {
	Expiration       *string       `json:"expiration,omitempty"`
	JumpHostsAllowed *[]string     `json:"jumpHostsAllowed,omitempty"`
	UserList         *[]KeySetUser `json:"userList,omitempty"`
}
