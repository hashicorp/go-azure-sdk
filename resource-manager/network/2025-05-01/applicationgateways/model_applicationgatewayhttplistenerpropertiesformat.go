package applicationgateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayHTTPListenerPropertiesFormat struct {
	CustomErrorConfigurations   *[]ApplicationGatewayCustomError `json:"customErrorConfigurations,omitempty"`
	FirewallPolicy              *CommonSubResource               `json:"firewallPolicy,omitempty"`
	FrontendIPConfiguration     *CommonSubResource               `json:"frontendIPConfiguration,omitempty"`
	FrontendPort                *CommonSubResource               `json:"frontendPort,omitempty"`
	HostName                    *string                          `json:"hostName,omitempty"`
	HostNames                   *[]string                        `json:"hostNames,omitempty"`
	Protocol                    *ApplicationGatewayProtocol      `json:"protocol,omitempty"`
	ProvisioningState           *ProvisioningState               `json:"provisioningState,omitempty"`
	RequireServerNameIndication *bool                            `json:"requireServerNameIndication,omitempty"`
	SslCertificate              *CommonSubResource               `json:"sslCertificate,omitempty"`
	SslProfile                  *CommonSubResource               `json:"sslProfile,omitempty"`
}
