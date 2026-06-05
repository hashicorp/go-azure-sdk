package publicipaddresses

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonApplicationGatewayIPConfigurationPropertiesFormat struct {
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	Subnet            *CommonSubResource `json:"subnet,omitempty"`
}
