package connectiongateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionGatewayReference struct {
	Id       *string `json:"id,omitempty"`
	Location *string `json:"location,omitempty"`
	Name     *string `json:"name,omitempty"`
	Type     *string `json:"type,omitempty"`
}
