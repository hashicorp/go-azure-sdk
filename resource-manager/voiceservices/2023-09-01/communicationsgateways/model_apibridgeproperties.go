package communicationsgateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiBridgeProperties struct {
	AllowedAddressPrefixes *[]string                 `json:"allowedAddressPrefixes,omitempty"`
	ConfigureApiBridge     *ApiBridgeActivationState `json:"configureApiBridge,omitempty"`
	EndpointFqdns          *[]string                 `json:"endpointFqdns,omitempty"`
}
