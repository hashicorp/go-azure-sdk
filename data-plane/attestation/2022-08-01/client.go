package v2022_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/data-plane/attestation/2022-08-01/attestation"
	"github.com/hashicorp/go-azure-sdk/data-plane/attestation/2022-08-01/openidmetadatadiscovery"
	"github.com/hashicorp/go-azure-sdk/data-plane/attestation/2022-08-01/policy"
	"github.com/hashicorp/go-azure-sdk/data-plane/attestation/2022-08-01/policymanagementcertificates"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

type Client struct {
	Attestation                  *attestation.AttestationClient
	OpenIDMetadataDiscovery      *openidmetadatadiscovery.OpenIDMetadataDiscoveryClient
	Policy                       *policy.PolicyClient
	PolicyManagementCertificates *policymanagementcertificates.PolicyManagementCertificatesClient
}

func NewClient(configureFunc func(c *dataplane.Client)) (*Client, error) {
	attestationClient, err := attestation.NewAttestationClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Attestation client: %+v", err)
	}
	configureFunc(attestationClient.Client)

	openIDMetadataDiscoveryClient, err := openidmetadatadiscovery.NewOpenIDMetadataDiscoveryClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building OpenIDMetadataDiscovery client: %+v", err)
	}
	configureFunc(openIDMetadataDiscoveryClient.Client)

	policyClient, err := policy.NewPolicyClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Policy client: %+v", err)
	}
	configureFunc(policyClient.Client)

	policyManagementCertificatesClient, err := policymanagementcertificates.NewPolicyManagementCertificatesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building PolicyManagementCertificates client: %+v", err)
	}
	configureFunc(policyManagementCertificatesClient.Client)

	return &Client{
		Attestation:                  attestationClient,
		OpenIDMetadataDiscovery:      openIDMetadataDiscoveryClient,
		Policy:                       policyClient,
		PolicyManagementCertificates: policyManagementCertificatesClient,
	}, nil
}
