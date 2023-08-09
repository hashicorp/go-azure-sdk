package v2022_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-06-01/policyassignments"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	PolicyAssignments *policyassignments.PolicyAssignmentsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	policyAssignmentsClient, err := policyassignments.NewPolicyAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAssignments client: %+v", err)
	}
	configureFunc(policyAssignmentsClient.Client)

	return &Client{
		PolicyAssignments: policyAssignmentsClient,
	}, nil
}
