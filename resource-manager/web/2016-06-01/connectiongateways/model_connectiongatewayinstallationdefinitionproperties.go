package connectiongateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionGatewayInstallationDefinitionProperties struct {
	BackendUri         *string                     `json:"backendUri,omitempty"`
	ConnectionGateway  *ConnectionGatewayReference `json:"connectionGateway,omitempty"`
	ContactInformation *[]string                   `json:"contactInformation,omitempty"`
	Description        *string                     `json:"description,omitempty"`
	DisplayName        *string                     `json:"displayName,omitempty"`
	MachineName        *string                     `json:"machineName,omitempty"`
	Status             *interface{}                `json:"status,omitempty"`
}
