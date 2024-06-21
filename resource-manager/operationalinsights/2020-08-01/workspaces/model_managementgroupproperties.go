package workspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementGroupProperties struct {
	Created      *string `json:"created,omitempty"`
	DataReceived *string `json:"dataReceived,omitempty"`
	Id           *string `json:"id,omitempty"`
	IsGateway    *bool   `json:"isGateway,omitempty"`
	Name         *string `json:"name,omitempty"`
	ServerCount  *int64  `json:"serverCount,omitempty"`
	Sku          *string `json:"sku,omitempty"`
	Version      *string `json:"version,omitempty"`
}
