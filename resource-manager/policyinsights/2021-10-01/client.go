package v2021_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2021-10-01/remediations"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Remediations *remediations.RemediationsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	remediationsClient, err := remediations.NewRemediationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Remediations client: %+v", err)
	}
	configureFunc(remediationsClient.Client)

	return &Client{
		Remediations: remediationsClient,
	}, nil
}
