package v2023_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/cognitiveservicesaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/cognitiveservicescommitmentplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/commitmentplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/commitmenttiers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/models"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/usages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CognitiveServicesAccounts        *cognitiveservicesaccounts.CognitiveServicesAccountsClient
	CognitiveServicesCommitmentPlans *cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient
	CommitmentPlans                  *commitmentplans.CommitmentPlansClient
	CommitmentTiers                  *commitmenttiers.CommitmentTiersClient
	Deployments                      *deployments.DeploymentsClient
	Models                           *models.ModelsClient
	PrivateEndpointConnections       *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources             *privatelinkresources.PrivateLinkResourcesClient
	Skus                             *skus.SkusClient
	Usages                           *usages.UsagesClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	cognitiveServicesAccountsClient, err := cognitiveservicesaccounts.NewCognitiveServicesAccountsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CognitiveServicesAccounts client: %+v", err)
	}
	configureFunc(cognitiveServicesAccountsClient.Client)

	cognitiveServicesCommitmentPlansClient, err := cognitiveservicescommitmentplans.NewCognitiveServicesCommitmentPlansClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CognitiveServicesCommitmentPlans client: %+v", err)
	}
	configureFunc(cognitiveServicesCommitmentPlansClient.Client)

	commitmentPlansClient, err := commitmentplans.NewCommitmentPlansClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CommitmentPlans client: %+v", err)
	}
	configureFunc(commitmentPlansClient.Client)

	commitmentTiersClient, err := commitmenttiers.NewCommitmentTiersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CommitmentTiers client: %+v", err)
	}
	configureFunc(commitmentTiersClient.Client)

	deploymentsClient, err := deployments.NewDeploymentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Deployments client: %+v", err)
	}
	configureFunc(deploymentsClient.Client)

	modelsClient, err := models.NewModelsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Models client: %+v", err)
	}
	configureFunc(modelsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	skusClient, err := skus.NewSkusClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Skus client: %+v", err)
	}
	configureFunc(skusClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	return &Client{
		CognitiveServicesAccounts:        cognitiveServicesAccountsClient,
		CognitiveServicesCommitmentPlans: cognitiveServicesCommitmentPlansClient,
		CommitmentPlans:                  commitmentPlansClient,
		CommitmentTiers:                  commitmentTiersClient,
		Deployments:                      deploymentsClient,
		Models:                           modelsClient,
		PrivateEndpointConnections:       privateEndpointConnectionsClient,
		PrivateLinkResources:             privateLinkResourcesClient,
		Skus:                             skusClient,
		Usages:                           usagesClient,
	}, nil
}
