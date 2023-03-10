package v2021_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2021-10-01/remediations"
)

type Client struct {
	Remediations *remediations.RemediationsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	remediationsClient := remediations.NewRemediationsClientWithBaseURI(endpoint)
	configureAuthFunc(&remediationsClient.Client)

	return Client{
		Remediations: &remediationsClient,
	}
}
