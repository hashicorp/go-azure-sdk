package applicationgateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayListenerPropertiesFormat struct {
	FrontendIPConfiguration *CommonSubResource          `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *CommonSubResource          `json:"frontendPort,omitempty"`
	HostNames               *[]string                   `json:"hostNames,omitempty"`
	Protocol                *ApplicationGatewayProtocol `json:"protocol,omitempty"`
	ProvisioningState       *ProvisioningState          `json:"provisioningState,omitempty"`
	SslCertificate          *CommonSubResource          `json:"sslCertificate,omitempty"`
	SslProfile              *CommonSubResource          `json:"sslProfile,omitempty"`
}
