package v2024_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/cognitiveservicesaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/cognitiveservicescommitmentplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/commitmentplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/commitmenttiers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/defenderforaisettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/encryptionscopes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/modelcapacities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/models"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/networksecurityperimeterconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/raiblocklists"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/raicontentfilters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/raipolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2024-10-01/usages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CognitiveServicesAccounts              *cognitiveservicesaccounts.CognitiveServicesAccountsClient
	CognitiveServicesCommitmentPlans       *cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient
	CommitmentPlans                        *commitmentplans.CommitmentPlansClient
	CommitmentTiers                        *commitmenttiers.CommitmentTiersClient
	DefenderForAISettings                  *defenderforaisettings.DefenderForAISettingsClient
	Deployments                            *deployments.DeploymentsClient
	EncryptionScopes                       *encryptionscopes.EncryptionScopesClient
	ModelCapacities                        *modelcapacities.ModelCapacitiesClient
	Models                                 *models.ModelsClient
	NetworkSecurityPerimeterConfigurations *networksecurityperimeterconfigurations.NetworkSecurityPerimeterConfigurationsClient
	PrivateEndpointConnections             *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                   *privatelinkresources.PrivateLinkResourcesClient
	RaiBlocklists                          *raiblocklists.RaiBlocklistsClient
	RaiContentFilters                      *raicontentfilters.RaiContentFiltersClient
	RaiPolicies                            *raipolicies.RaiPoliciesClient
	Skus                                   *skus.SkusClient
	Usages                                 *usages.UsagesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	cognitiveServicesAccountsClient, err := cognitiveservicesaccounts.NewCognitiveServicesAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CognitiveServicesAccounts client: %+v", err)
	}
	configureFunc(cognitiveServicesAccountsClient.Client)

	cognitiveServicesCommitmentPlansClient, err := cognitiveservicescommitmentplans.NewCognitiveServicesCommitmentPlansClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CognitiveServicesCommitmentPlans client: %+v", err)
	}
	configureFunc(cognitiveServicesCommitmentPlansClient.Client)

	commitmentPlansClient, err := commitmentplans.NewCommitmentPlansClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CommitmentPlans client: %+v", err)
	}
	configureFunc(commitmentPlansClient.Client)

	commitmentTiersClient, err := commitmenttiers.NewCommitmentTiersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CommitmentTiers client: %+v", err)
	}
	configureFunc(commitmentTiersClient.Client)

	defenderForAISettingsClient, err := defenderforaisettings.NewDefenderForAISettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderForAISettings client: %+v", err)
	}
	configureFunc(defenderForAISettingsClient.Client)

	deploymentsClient, err := deployments.NewDeploymentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Deployments client: %+v", err)
	}
	configureFunc(deploymentsClient.Client)

	encryptionScopesClient, err := encryptionscopes.NewEncryptionScopesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EncryptionScopes client: %+v", err)
	}
	configureFunc(encryptionScopesClient.Client)

	modelCapacitiesClient, err := modelcapacities.NewModelCapacitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ModelCapacities client: %+v", err)
	}
	configureFunc(modelCapacitiesClient.Client)

	modelsClient, err := models.NewModelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Models client: %+v", err)
	}
	configureFunc(modelsClient.Client)

	networkSecurityPerimeterConfigurationsClient, err := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterConfigurations client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterConfigurationsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	raiBlocklistsClient, err := raiblocklists.NewRaiBlocklistsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RaiBlocklists client: %+v", err)
	}
	configureFunc(raiBlocklistsClient.Client)

	raiContentFiltersClient, err := raicontentfilters.NewRaiContentFiltersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RaiContentFilters client: %+v", err)
	}
	configureFunc(raiContentFiltersClient.Client)

	raiPoliciesClient, err := raipolicies.NewRaiPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RaiPolicies client: %+v", err)
	}
	configureFunc(raiPoliciesClient.Client)

	skusClient, err := skus.NewSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Skus client: %+v", err)
	}
	configureFunc(skusClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	return &Client{
		CognitiveServicesAccounts:              cognitiveServicesAccountsClient,
		CognitiveServicesCommitmentPlans:       cognitiveServicesCommitmentPlansClient,
		CommitmentPlans:                        commitmentPlansClient,
		CommitmentTiers:                        commitmentTiersClient,
		DefenderForAISettings:                  defenderForAISettingsClient,
		Deployments:                            deploymentsClient,
		EncryptionScopes:                       encryptionScopesClient,
		ModelCapacities:                        modelCapacitiesClient,
		Models:                                 modelsClient,
		NetworkSecurityPerimeterConfigurations: networkSecurityPerimeterConfigurationsClient,
		PrivateEndpointConnections:             privateEndpointConnectionsClient,
		PrivateLinkResources:                   privateLinkResourcesClient,
		RaiBlocklists:                          raiBlocklistsClient,
		RaiContentFilters:                      raiContentFiltersClient,
		RaiPolicies:                            raiPoliciesClient,
		Skus:                                   skusClient,
		Usages:                                 usagesClient,
	}, nil
}
