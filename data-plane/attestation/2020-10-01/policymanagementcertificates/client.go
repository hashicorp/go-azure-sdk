package policymanagementcertificates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyManagementCertificatesClient struct {
	Client *dataplane.Client
}

func NewPolicyManagementCertificatesClientUnconfigured() (*PolicyManagementCertificatesClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "policymanagementcertificates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyManagementCertificatesClient: %+v", err)
	}

	return &PolicyManagementCertificatesClient{
		Client: client,
	}, nil
}

func (c *PolicyManagementCertificatesClient) PolicyManagementCertificatesClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewPolicyManagementCertificatesClientWithBaseURI(endpoint string) (*PolicyManagementCertificatesClient, error) {
	client, err := dataplane.NewClient(endpoint, "policymanagementcertificates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyManagementCertificatesClient: %+v", err)
	}

	return &PolicyManagementCertificatesClient{
		Client: client,
	}, nil
}

func (c *PolicyManagementCertificatesClient) Clone(endpoint string) *PolicyManagementCertificatesClient {
	return &PolicyManagementCertificatesClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
