package v2025_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/deploymentoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/policyassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/policydefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/policydefinitionversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/policysetdefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/policysetdefinitionversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/policytokens"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/providers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/resourcegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/resources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2025-03-01/tags"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DeploymentOperations        *deploymentoperations.DeploymentOperationsClient
	Deployments                 *deployments.DeploymentsClient
	PolicyAssignments           *policyassignments.PolicyAssignmentsClient
	PolicyDefinitionVersions    *policydefinitionversions.PolicyDefinitionVersionsClient
	PolicyDefinitions           *policydefinitions.PolicyDefinitionsClient
	PolicySetDefinitionVersions *policysetdefinitionversions.PolicySetDefinitionVersionsClient
	PolicySetDefinitions        *policysetdefinitions.PolicySetDefinitionsClient
	PolicyTokens                *policytokens.PolicyTokensClient
	Providers                   *providers.ProvidersClient
	ResourceGroups              *resourcegroups.ResourceGroupsClient
	Resources                   *resources.ResourcesClient
	Tags                        *tags.TagsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	deploymentOperationsClient, err := deploymentoperations.NewDeploymentOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentOperations client: %+v", err)
	}
	configureFunc(deploymentOperationsClient.Client)

	deploymentsClient, err := deployments.NewDeploymentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Deployments client: %+v", err)
	}
	configureFunc(deploymentsClient.Client)

	policyAssignmentsClient, err := policyassignments.NewPolicyAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAssignments client: %+v", err)
	}
	configureFunc(policyAssignmentsClient.Client)

	policyDefinitionVersionsClient, err := policydefinitionversions.NewPolicyDefinitionVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyDefinitionVersions client: %+v", err)
	}
	configureFunc(policyDefinitionVersionsClient.Client)

	policyDefinitionsClient, err := policydefinitions.NewPolicyDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyDefinitions client: %+v", err)
	}
	configureFunc(policyDefinitionsClient.Client)

	policySetDefinitionVersionsClient, err := policysetdefinitionversions.NewPolicySetDefinitionVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicySetDefinitionVersions client: %+v", err)
	}
	configureFunc(policySetDefinitionVersionsClient.Client)

	policySetDefinitionsClient, err := policysetdefinitions.NewPolicySetDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicySetDefinitions client: %+v", err)
	}
	configureFunc(policySetDefinitionsClient.Client)

	policyTokensClient, err := policytokens.NewPolicyTokensClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyTokens client: %+v", err)
	}
	configureFunc(policyTokensClient.Client)

	providersClient, err := providers.NewProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Providers client: %+v", err)
	}
	configureFunc(providersClient.Client)

	resourceGroupsClient, err := resourcegroups.NewResourceGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceGroups client: %+v", err)
	}
	configureFunc(resourceGroupsClient.Client)

	resourcesClient, err := resources.NewResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Resources client: %+v", err)
	}
	configureFunc(resourcesClient.Client)

	tagsClient, err := tags.NewTagsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Tags client: %+v", err)
	}
	configureFunc(tagsClient.Client)

	return &Client{
		DeploymentOperations:        deploymentOperationsClient,
		Deployments:                 deploymentsClient,
		PolicyAssignments:           policyAssignmentsClient,
		PolicyDefinitionVersions:    policyDefinitionVersionsClient,
		PolicyDefinitions:           policyDefinitionsClient,
		PolicySetDefinitionVersions: policySetDefinitionVersionsClient,
		PolicySetDefinitions:        policySetDefinitionsClient,
		PolicyTokens:                policyTokensClient,
		Providers:                   providersClient,
		ResourceGroups:              resourceGroupsClient,
		Resources:                   resourcesClient,
		Tags:                        tagsClient,
	}, nil
}
