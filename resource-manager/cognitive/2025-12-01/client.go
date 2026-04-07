package v2025_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/accountcapabilityhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/accountconnectionresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/agentapplication"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/agentdeployment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/cognitiveservicesaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/cognitiveservicescommitmentplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/cognitiveservicesprojects"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/commitmentplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/commitmenttiers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/defenderforaisettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/encryptionscopes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/managednetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/modelcapacities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/models"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/networksecurityperimeterconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/projectcapabilityhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/projectconnectionresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/quotatier"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/raiblocklists"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/raicontentfilters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/raiexternalsafetyprovider"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/raipolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/raitoollabels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/raitopics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/subscriptionraipolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/testraiexternalsafetyprovider"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/usages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AccountCapabilityHost                  *accountcapabilityhost.AccountCapabilityHostClient
	AccountConnectionResource              *accountconnectionresource.AccountConnectionResourceClient
	AgentApplication                       *agentapplication.AgentApplicationClient
	AgentDeployment                        *agentdeployment.AgentDeploymentClient
	CognitiveServicesAccounts              *cognitiveservicesaccounts.CognitiveServicesAccountsClient
	CognitiveServicesCommitmentPlans       *cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient
	CognitiveServicesProjects              *cognitiveservicesprojects.CognitiveServicesProjectsClient
	CommitmentPlans                        *commitmentplans.CommitmentPlansClient
	CommitmentTiers                        *commitmenttiers.CommitmentTiersClient
	DefenderForAISettings                  *defenderforaisettings.DefenderForAISettingsClient
	Deployments                            *deployments.DeploymentsClient
	EncryptionScopes                       *encryptionscopes.EncryptionScopesClient
	ManagedNetwork                         *managednetwork.ManagedNetworkClient
	ModelCapacities                        *modelcapacities.ModelCapacitiesClient
	Models                                 *models.ModelsClient
	NetworkSecurityPerimeterConfigurations *networksecurityperimeterconfigurations.NetworkSecurityPerimeterConfigurationsClient
	PrivateEndpointConnections             *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                   *privatelinkresources.PrivateLinkResourcesClient
	ProjectCapabilityHost                  *projectcapabilityhost.ProjectCapabilityHostClient
	ProjectConnectionResource              *projectconnectionresource.ProjectConnectionResourceClient
	QuotaTier                              *quotatier.QuotaTierClient
	RaiBlocklists                          *raiblocklists.RaiBlocklistsClient
	RaiContentFilters                      *raicontentfilters.RaiContentFiltersClient
	RaiExternalSafetyProvider              *raiexternalsafetyprovider.RaiExternalSafetyProviderClient
	RaiPolicies                            *raipolicies.RaiPoliciesClient
	RaiToolLabels                          *raitoollabels.RaiToolLabelsClient
	RaiTopics                              *raitopics.RaiTopicsClient
	Skus                                   *skus.SkusClient
	SubscriptionRaiPolicy                  *subscriptionraipolicy.SubscriptionRaiPolicyClient
	TestRaiExternalSafetyProvider          *testraiexternalsafetyprovider.TestRaiExternalSafetyProviderClient
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

	agentApplicationClient, err := agentapplication.NewAgentApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AgentApplication client: %+v", err)
	}
	configureFunc(agentApplicationClient.Client)

	agentDeploymentClient, err := agentdeployment.NewAgentDeploymentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AgentDeployment client: %+v", err)
	}
	configureFunc(agentDeploymentClient.Client)

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

	managedNetworkClient, err := managednetwork.NewManagedNetworkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedNetwork client: %+v", err)
	}
	configureFunc(managedNetworkClient.Client)

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

	quotaTierClient, err := quotatier.NewQuotaTierClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QuotaTier client: %+v", err)
	}
	configureFunc(quotaTierClient.Client)

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

	raiExternalSafetyProviderClient, err := raiexternalsafetyprovider.NewRaiExternalSafetyProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RaiExternalSafetyProvider client: %+v", err)
	}
	configureFunc(raiExternalSafetyProviderClient.Client)

	raiPoliciesClient, err := raipolicies.NewRaiPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RaiPolicies client: %+v", err)
	}
	configureFunc(raiPoliciesClient.Client)

	raiToolLabelsClient, err := raitoollabels.NewRaiToolLabelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RaiToolLabels client: %+v", err)
	}
	configureFunc(raiToolLabelsClient.Client)

	raiTopicsClient, err := raitopics.NewRaiTopicsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RaiTopics client: %+v", err)
	}
	configureFunc(raiTopicsClient.Client)

	skusClient, err := skus.NewSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Skus client: %+v", err)
	}
	configureFunc(skusClient.Client)

	subscriptionRaiPolicyClient, err := subscriptionraipolicy.NewSubscriptionRaiPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubscriptionRaiPolicy client: %+v", err)
	}
	configureFunc(subscriptionRaiPolicyClient.Client)

	testRaiExternalSafetyProviderClient, err := testraiexternalsafetyprovider.NewTestRaiExternalSafetyProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TestRaiExternalSafetyProvider client: %+v", err)
	}
	configureFunc(testRaiExternalSafetyProviderClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	return &Client{
		AccountCapabilityHost:                  accountCapabilityHostClient,
		AccountConnectionResource:              accountConnectionResourceClient,
		AgentApplication:                       agentApplicationClient,
		AgentDeployment:                        agentDeploymentClient,
		CognitiveServicesAccounts:              cognitiveServicesAccountsClient,
		CognitiveServicesCommitmentPlans:       cognitiveServicesCommitmentPlansClient,
		CognitiveServicesProjects:              cognitiveServicesProjectsClient,
		CommitmentPlans:                        commitmentPlansClient,
		CommitmentTiers:                        commitmentTiersClient,
		DefenderForAISettings:                  defenderForAISettingsClient,
		Deployments:                            deploymentsClient,
		EncryptionScopes:                       encryptionScopesClient,
		ManagedNetwork:                         managedNetworkClient,
		ModelCapacities:                        modelCapacitiesClient,
		Models:                                 modelsClient,
		NetworkSecurityPerimeterConfigurations: networkSecurityPerimeterConfigurationsClient,
		PrivateEndpointConnections:             privateEndpointConnectionsClient,
		PrivateLinkResources:                   privateLinkResourcesClient,
		ProjectCapabilityHost:                  projectCapabilityHostClient,
		ProjectConnectionResource:              projectConnectionResourceClient,
		QuotaTier:                              quotaTierClient,
		RaiBlocklists:                          raiBlocklistsClient,
		RaiContentFilters:                      raiContentFiltersClient,
		RaiExternalSafetyProvider:              raiExternalSafetyProviderClient,
		RaiPolicies:                            raiPoliciesClient,
		RaiToolLabels:                          raiToolLabelsClient,
		RaiTopics:                              raiTopicsClient,
		Skus:                                   skusClient,
		SubscriptionRaiPolicy:                  subscriptionRaiPolicyClient,
		TestRaiExternalSafetyProvider:          testRaiExternalSafetyProviderClient,
		Usages:                                 usagesClient,
	}, nil
}
