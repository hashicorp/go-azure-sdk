package v2022_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2022-09-01/attestations"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Attestations *attestations.AttestationsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	attestationsClient, err := attestations.NewAttestationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Attestations client: %+v", err)
	}
	configureFunc(attestationsClient.Client)

	return &Client{
		Attestations: attestationsClient,
	}, nil
}
