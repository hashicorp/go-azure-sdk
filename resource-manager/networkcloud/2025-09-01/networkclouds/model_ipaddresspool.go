package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPAddressPool struct {
	Addresses      []string    `json:"addresses"`
	AutoAssign     *BfdEnabled `json:"autoAssign,omitempty"`
	Name           string      `json:"name"`
	OnlyUseHostIPs *BfdEnabled `json:"onlyUseHostIps,omitempty"`
}
