package v2021_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-06-01/policyassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-06-01/policydefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-06-01/policysetdefinitions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	PolicyAssignments    *policyassignments.PolicyAssignmentsClient
	PolicyDefinitions    *policydefinitions.PolicyDefinitionsClient
	PolicySetDefinitions *policysetdefinitions.PolicySetDefinitionsClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	policyAssignmentsClient, err := policyassignments.NewPolicyAssignmentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAssignments client: %+v", err)
	}
	configureFunc(policyAssignmentsClient.Client)

	policyDefinitionsClient, err := policydefinitions.NewPolicyDefinitionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PolicyDefinitions client: %+v", err)
	}
	configureFunc(policyDefinitionsClient.Client)

	policySetDefinitionsClient, err := policysetdefinitions.NewPolicySetDefinitionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PolicySetDefinitions client: %+v", err)
	}
	configureFunc(policySetDefinitionsClient.Client)

	return &Client{
		PolicyAssignments:    policyAssignmentsClient,
		PolicyDefinitions:    policyDefinitionsClient,
		PolicySetDefinitions: policySetDefinitionsClient,
	}, nil
}
