package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachedNetworkConfiguration struct {
	L2Networks      *[]L2NetworkAttachmentConfiguration      `json:"l2Networks,omitempty"`
	L3Networks      *[]L3NetworkAttachmentConfiguration      `json:"l3Networks,omitempty"`
	TrunkedNetworks *[]TrunkedNetworkAttachmentConfiguration `json:"trunkedNetworks,omitempty"`
}
