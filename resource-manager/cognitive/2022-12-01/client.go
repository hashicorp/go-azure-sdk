// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_12_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2022-12-01/cognitiveservicesaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2022-12-01/cognitiveservicescommitmentplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2022-12-01/commitmentplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2022-12-01/commitmenttiers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2022-12-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2022-12-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2022-12-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2022-12-01/skus"
)

type Client struct {
	CognitiveServicesAccounts        *cognitiveservicesaccounts.CognitiveServicesAccountsClient
	CognitiveServicesCommitmentPlans *cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient
	CommitmentPlans                  *commitmentplans.CommitmentPlansClient
	CommitmentTiers                  *commitmenttiers.CommitmentTiersClient
	Deployments                      *deployments.DeploymentsClient
	PrivateEndpointConnections       *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources             *privatelinkresources.PrivateLinkResourcesClient
	Skus                             *skus.SkusClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	cognitiveServicesAccountsClient := cognitiveservicesaccounts.NewCognitiveServicesAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&cognitiveServicesAccountsClient.Client)

	cognitiveServicesCommitmentPlansClient := cognitiveservicescommitmentplans.NewCognitiveServicesCommitmentPlansClientWithBaseURI(endpoint)
	configureAuthFunc(&cognitiveServicesCommitmentPlansClient.Client)

	commitmentPlansClient := commitmentplans.NewCommitmentPlansClientWithBaseURI(endpoint)
	configureAuthFunc(&commitmentPlansClient.Client)

	commitmentTiersClient := commitmenttiers.NewCommitmentTiersClientWithBaseURI(endpoint)
	configureAuthFunc(&commitmentTiersClient.Client)

	deploymentsClient := deployments.NewDeploymentsClientWithBaseURI(endpoint)
	configureAuthFunc(&deploymentsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	skusClient := skus.NewSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&skusClient.Client)

	return Client{
		CognitiveServicesAccounts:        &cognitiveServicesAccountsClient,
		CognitiveServicesCommitmentPlans: &cognitiveServicesCommitmentPlansClient,
		CommitmentPlans:                  &commitmentPlansClient,
		CommitmentTiers:                  &commitmentTiersClient,
		Deployments:                      &deploymentsClient,
		PrivateEndpointConnections:       &privateEndpointConnectionsClient,
		PrivateLinkResources:             &privateLinkResourcesClient,
		Skus:                             &skusClient,
	}
}
