package v2025_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/accountcapabilityhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/accountconnectionresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/cognitiveservicesaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/cognitiveservicescommitmentplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/cognitiveservicesprojects"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/commitmentplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/commitmenttiers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/defenderforaisettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/encryptionscopes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/modelcapacities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/models"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/networksecurityperimeterconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/projectcapabilityhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/projectconnectionresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/raiblocklists"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/raicontentfilters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/raipolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/usages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AccountCapabilityHost                  *accountcapabilityhost.AccountCapabilityHostClient
	AccountConnectionResource              *accountconnectionresource.AccountConnectionResourceClient
	CognitiveServicesAccounts              *cognitiveservicesaccounts.CognitiveServicesAccountsClient
	CognitiveServicesCommitmentPlans       *cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient
	CognitiveServicesProjects              *cognitiveservicesprojects.CognitiveServicesProjectsClient
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
	ProjectCapabilityHost                  *projectcapabilityhost.ProjectCapabilityHostClient
	ProjectConnectionResource              *projectconnectionresource.ProjectConnectionResourceClient
	RaiBlocklists                          *raiblocklists.RaiBlocklistsClient
	RaiContentFilters                      *raicontentfilters.RaiContentFiltersClient
	RaiPolicies                            *raipolicies.RaiPoliciesClient
	Skus                                   *skus.SkusClient
	Usages                                 *usages.UsagesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountCapabilityHostClient, err := accountcapabilityhost.NewAccountCapabilityHostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccountCapabilityHost client: %+v", err)
	}
	configureFunc(accountCapabilityHostClient.Client)

	accountConnectionResourceClient, err := accountconnectionresource.NewAccountConnectionResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccountConnectionResource client: %+v", err)
	}
	configureFunc(accountConnectionResourceClient.Client)

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

	cognitiveServicesProjectsClient, err := cognitiveservicesprojects.NewCognitiveServicesProjectsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CognitiveServicesProjects client: %+v", err)
	}
	configureFunc(cognitiveServicesProjectsClient.Client)

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

	projectCapabilityHostClient, err := projectcapabilityhost.NewProjectCapabilityHostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProjectCapabilityHost client: %+v", err)
	}
	configureFunc(projectCapabilityHostClient.Client)

	projectConnectionResourceClient, err := projectconnectionresource.NewProjectConnectionResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProjectConnectionResource client: %+v", err)
	}
	configureFunc(projectConnectionResourceClient.Client)

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
		AccountCapabilityHost:                  accountCapabilityHostClient,
		AccountConnectionResource:              accountConnectionResourceClient,
		CognitiveServicesAccounts:              cognitiveServicesAccountsClient,
		CognitiveServicesCommitmentPlans:       cognitiveServicesCommitmentPlansClient,
		CognitiveServicesProjects:              cognitiveServicesProjectsClient,
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
		ProjectCapabilityHost:                  projectCapabilityHostClient,
		ProjectConnectionResource:              projectConnectionResourceClient,
		RaiBlocklists:                          raiBlocklistsClient,
		RaiContentFilters:                      raiContentFiltersClient,
		RaiPolicies:                            raiPoliciesClient,
		Skus:                                   skusClient,
		Usages:                                 usagesClient,
	}, nil
}
