package v2022_07_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-07-01-preview/policyexemptions"
)

type Client struct {
	PolicyExemptions *policyexemptions.PolicyExemptionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	policyExemptionsClient := policyexemptions.NewPolicyExemptionsClientWithBaseURI(endpoint)
	configureAuthFunc(&policyExemptionsClient.Client)

	return Client{
		PolicyExemptions: &policyExemptionsClient,
	}
}
