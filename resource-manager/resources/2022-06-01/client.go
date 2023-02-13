package v2022_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-06-01/policyassignments"
)

type Client struct {
	PolicyAssignments *policyassignments.PolicyAssignmentsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	policyAssignmentsClient := policyassignments.NewPolicyAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&policyAssignmentsClient.Client)

	return Client{
		PolicyAssignments: &policyAssignmentsClient,
	}
}
