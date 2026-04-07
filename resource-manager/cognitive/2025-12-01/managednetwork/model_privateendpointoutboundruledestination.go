package managednetwork

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointOutboundRuleDestination struct {
	ServiceResourceId *string `json:"serviceResourceId,omitempty"`
	SubresourceTarget *string `json:"subresourceTarget,omitempty"`
}
