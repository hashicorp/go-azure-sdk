package v2021_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-06-01/policyassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-06-01/policydefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-06-01/policysetdefinitions"
)

type Client struct {
	PolicyAssignments    *policyassignments.PolicyAssignmentsClient
	PolicyDefinitions    *policydefinitions.PolicyDefinitionsClient
	PolicySetDefinitions *policysetdefinitions.PolicySetDefinitionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	policyAssignmentsClient := policyassignments.NewPolicyAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&policyAssignmentsClient.Client)

	policyDefinitionsClient := policydefinitions.NewPolicyDefinitionsClientWithBaseURI(endpoint)
	configureAuthFunc(&policyDefinitionsClient.Client)

	policySetDefinitionsClient := policysetdefinitions.NewPolicySetDefinitionsClientWithBaseURI(endpoint)
	configureAuthFunc(&policySetDefinitionsClient.Client)

	return Client{
		PolicyAssignments:    &policyAssignmentsClient,
		PolicyDefinitions:    &policyDefinitionsClient,
		PolicySetDefinitions: &policySetDefinitionsClient,
	}
}
