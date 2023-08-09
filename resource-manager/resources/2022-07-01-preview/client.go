package v2022_07_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-07-01-preview/policyexemptions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	PolicyExemptions *policyexemptions.PolicyExemptionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	policyExemptionsClient, err := policyexemptions.NewPolicyExemptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyExemptions client: %+v", err)
	}
	configureFunc(policyExemptionsClient.Client)

	return &Client{
		PolicyExemptions: policyExemptionsClient,
	}, nil
}
