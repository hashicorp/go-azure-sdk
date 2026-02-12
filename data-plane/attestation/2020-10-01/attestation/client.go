package attestation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationClient struct {
	Client *dataplane.Client
}

func NewAttestationClientUnconfigured() (*AttestationClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "attestation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AttestationClient: %+v", err)
	}

	return &AttestationClient{
		Client: client,
	}, nil
}

func (c *AttestationClient) AttestationClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewAttestationClientWithBaseURI(endpoint string) (*AttestationClient, error) {
	client, err := dataplane.NewClient(endpoint, "attestation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AttestationClient: %+v", err)
	}

	return &AttestationClient{
		Client: client,
	}, nil
}

func (c *AttestationClient) Clone(endpoint string) *AttestationClient {
	return &AttestationClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
