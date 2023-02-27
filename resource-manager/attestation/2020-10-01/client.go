// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_10_01

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
