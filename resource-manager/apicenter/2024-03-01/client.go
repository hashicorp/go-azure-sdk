package v2024_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/apidefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/apis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/apiversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/environments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/metadataschemas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/services"
	"github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-01/workspaces"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ApiDefinitions  *apidefinitions.ApiDefinitionsClient
	ApiVersions     *apiversions.ApiVersionsClient
	Apis            *apis.ApisClient
	Deployments     *deployments.DeploymentsClient
	Environments    *environments.EnvironmentsClient
	MetadataSchemas *metadataschemas.MetadataSchemasClient
	Services        *services.ServicesClient
	Workspaces      *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	apiDefinitionsClient, err := apidefinitions.NewApiDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApiDefinitions client: %+v", err)
	}
	configureFunc(apiDefinitionsClient.Client)

	apiVersionsClient, err := apiversions.NewApiVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApiVersions client: %+v", err)
	}
	configureFunc(apiVersionsClient.Client)

	apisClient, err := apis.NewApisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Apis client: %+v", err)
	}
	configureFunc(apisClient.Client)

	deploymentsClient, err := deployments.NewDeploymentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Deployments client: %+v", err)
	}
	configureFunc(deploymentsClient.Client)

	environmentsClient, err := environments.NewEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Environments client: %+v", err)
	}
	configureFunc(environmentsClient.Client)

	metadataSchemasClient, err := metadataschemas.NewMetadataSchemasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MetadataSchemas client: %+v", err)
	}
	configureFunc(metadataSchemasClient.Client)

	servicesClient, err := services.NewServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Services client: %+v", err)
	}
	configureFunc(servicesClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		ApiDefinitions:  apiDefinitionsClient,
		ApiVersions:     apiVersionsClient,
		Apis:            apisClient,
		Deployments:     deploymentsClient,
		Environments:    environmentsClient,
		MetadataSchemas: metadataSchemasClient,
		Services:        servicesClient,
		Workspaces:      workspacesClient,
	}, nil
}
