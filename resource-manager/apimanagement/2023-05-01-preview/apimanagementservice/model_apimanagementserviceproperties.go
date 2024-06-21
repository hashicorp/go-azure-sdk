package apimanagementservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiManagementServiceProperties struct {
	AdditionalLocations         *[]AdditionalLocation                     `json:"additionalLocations,omitempty"`
	ApiVersionConstraint        *ApiVersionConstraint                     `json:"apiVersionConstraint,omitempty"`
	Certificates                *[]CertificateConfiguration               `json:"certificates,omitempty"`
	ConfigurationApi            *ConfigurationApi                         `json:"configurationApi,omitempty"`
	CreatedAtUtc                *string                                   `json:"createdAtUtc,omitempty"`
	CustomProperties            *map[string]string                        `json:"customProperties,omitempty"`
	DeveloperPortalStatus       *DeveloperPortalStatus                    `json:"developerPortalStatus,omitempty"`
	DeveloperPortalUrl          *string                                   `json:"developerPortalUrl,omitempty"`
	DisableGateway              *bool                                     `json:"disableGateway,omitempty"`
	EnableClientCertificate     *bool                                     `json:"enableClientCertificate,omitempty"`
	GatewayRegionalUrl          *string                                   `json:"gatewayRegionalUrl,omitempty"`
	GatewayUrl                  *string                                   `json:"gatewayUrl,omitempty"`
	HostnameConfigurations      *[]HostnameConfiguration                  `json:"hostnameConfigurations,omitempty"`
	LegacyPortalStatus          *LegacyPortalStatus                       `json:"legacyPortalStatus,omitempty"`
	ManagementApiUrl            *string                                   `json:"managementApiUrl,omitempty"`
	NatGatewayState             *NatGatewayState                          `json:"natGatewayState,omitempty"`
	NotificationSenderEmail     *string                                   `json:"notificationSenderEmail,omitempty"`
	OutboundPublicIPAddresses   *[]string                                 `json:"outboundPublicIPAddresses,omitempty"`
	PlatformVersion             *PlatformVersion                          `json:"platformVersion,omitempty"`
	PortalUrl                   *string                                   `json:"portalUrl,omitempty"`
	PrivateEndpointConnections  *[]RemotePrivateEndpointConnectionWrapper `json:"privateEndpointConnections,omitempty"`
	PrivateIPAddresses          *[]string                                 `json:"privateIPAddresses,omitempty"`
	ProvisioningState           *string                                   `json:"provisioningState,omitempty"`
	PublicIPAddressId           *string                                   `json:"publicIpAddressId,omitempty"`
	PublicIPAddresses           *[]string                                 `json:"publicIPAddresses,omitempty"`
	PublicNetworkAccess         *PublicNetworkAccess                      `json:"publicNetworkAccess,omitempty"`
	PublisherEmail              string                                    `json:"publisherEmail"`
	PublisherName               string                                    `json:"publisherName"`
	Restore                     *bool                                     `json:"restore,omitempty"`
	ScmUrl                      *string                                   `json:"scmUrl,omitempty"`
	TargetProvisioningState     *string                                   `json:"targetProvisioningState,omitempty"`
	VirtualNetworkConfiguration *VirtualNetworkConfiguration              `json:"virtualNetworkConfiguration,omitempty"`
	VirtualNetworkType          *VirtualNetworkType                       `json:"virtualNetworkType,omitempty"`
}
