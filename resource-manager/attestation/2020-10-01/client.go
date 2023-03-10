package v2020_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/attestation/2020-10-01/attestationproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/attestation/2020-10-01/privateendpointconnections"
)

type Client struct {
	AttestationProviders       *attestationproviders.AttestationProvidersClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	attestationProvidersClient := attestationproviders.NewAttestationProvidersClientWithBaseURI(endpoint)
	configureAuthFunc(&attestationProvidersClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	return Client{
		AttestationProviders:       &attestationProvidersClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
	}
}
