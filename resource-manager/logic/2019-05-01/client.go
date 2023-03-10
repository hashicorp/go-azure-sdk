package v2019_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountagreements"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountassemblies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountbatchconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountcertificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountmaps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountpartners"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountschemas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationaccountsessions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentmanagedapi"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentmanagedapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentnetworkhealth"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentskus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/workflowrunactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/workflowruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/workflows"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/workflowtriggerhistories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/workflowtriggers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/workflowversions"
)

type Client struct {
	IntegrationAccountAgreements               *integrationaccountagreements.IntegrationAccountAgreementsClient
	IntegrationAccountAssemblies               *integrationaccountassemblies.IntegrationAccountAssembliesClient
	IntegrationAccountBatchConfigurations      *integrationaccountbatchconfigurations.IntegrationAccountBatchConfigurationsClient
	IntegrationAccountCertificates             *integrationaccountcertificates.IntegrationAccountCertificatesClient
	IntegrationAccountMaps                     *integrationaccountmaps.IntegrationAccountMapsClient
	IntegrationAccountPartners                 *integrationaccountpartners.IntegrationAccountPartnersClient
	IntegrationAccountSchemas                  *integrationaccountschemas.IntegrationAccountSchemasClient
	IntegrationAccountSessions                 *integrationaccountsessions.IntegrationAccountSessionsClient
	IntegrationAccounts                        *integrationaccounts.IntegrationAccountsClient
	IntegrationServiceEnvironmentManagedApi    *integrationserviceenvironmentmanagedapi.IntegrationServiceEnvironmentManagedApiClient
	IntegrationServiceEnvironmentManagedApis   *integrationserviceenvironmentmanagedapis.IntegrationServiceEnvironmentManagedApisClient
	IntegrationServiceEnvironmentNetworkHealth *integrationserviceenvironmentnetworkhealth.IntegrationServiceEnvironmentNetworkHealthClient
	IntegrationServiceEnvironmentRestart       *integrationserviceenvironmentrestart.IntegrationServiceEnvironmentRestartClient
	IntegrationServiceEnvironmentSkus          *integrationserviceenvironmentskus.IntegrationServiceEnvironmentSkusClient
	IntegrationServiceEnvironments             *integrationserviceenvironments.IntegrationServiceEnvironmentsClient
	WorkflowRunActions                         *workflowrunactions.WorkflowRunActionsClient
	WorkflowRuns                               *workflowruns.WorkflowRunsClient
	WorkflowTriggerHistories                   *workflowtriggerhistories.WorkflowTriggerHistoriesClient
	WorkflowTriggers                           *workflowtriggers.WorkflowTriggersClient
	WorkflowVersions                           *workflowversions.WorkflowVersionsClient
	Workflows                                  *workflows.WorkflowsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	integrationAccountAgreementsClient := integrationaccountagreements.NewIntegrationAccountAgreementsClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationAccountAgreementsClient.Client)

	integrationAccountAssembliesClient := integrationaccountassemblies.NewIntegrationAccountAssembliesClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationAccountAssembliesClient.Client)

	integrationAccountBatchConfigurationsClient := integrationaccountbatchconfigurations.NewIntegrationAccountBatchConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationAccountBatchConfigurationsClient.Client)

	integrationAccountCertificatesClient := integrationaccountcertificates.NewIntegrationAccountCertificatesClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationAccountCertificatesClient.Client)

	integrationAccountMapsClient := integrationaccountmaps.NewIntegrationAccountMapsClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationAccountMapsClient.Client)

	integrationAccountPartnersClient := integrationaccountpartners.NewIntegrationAccountPartnersClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationAccountPartnersClient.Client)

	integrationAccountSchemasClient := integrationaccountschemas.NewIntegrationAccountSchemasClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationAccountSchemasClient.Client)

	integrationAccountSessionsClient := integrationaccountsessions.NewIntegrationAccountSessionsClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationAccountSessionsClient.Client)

	integrationAccountsClient := integrationaccounts.NewIntegrationAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationAccountsClient.Client)

	integrationServiceEnvironmentManagedApiClient := integrationserviceenvironmentmanagedapi.NewIntegrationServiceEnvironmentManagedApiClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationServiceEnvironmentManagedApiClient.Client)

	integrationServiceEnvironmentManagedApisClient := integrationserviceenvironmentmanagedapis.NewIntegrationServiceEnvironmentManagedApisClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationServiceEnvironmentManagedApisClient.Client)

	integrationServiceEnvironmentNetworkHealthClient := integrationserviceenvironmentnetworkhealth.NewIntegrationServiceEnvironmentNetworkHealthClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationServiceEnvironmentNetworkHealthClient.Client)

	integrationServiceEnvironmentRestartClient := integrationserviceenvironmentrestart.NewIntegrationServiceEnvironmentRestartClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationServiceEnvironmentRestartClient.Client)

	integrationServiceEnvironmentSkusClient := integrationserviceenvironmentskus.NewIntegrationServiceEnvironmentSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationServiceEnvironmentSkusClient.Client)

	integrationServiceEnvironmentsClient := integrationserviceenvironments.NewIntegrationServiceEnvironmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationServiceEnvironmentsClient.Client)

	workflowRunActionsClient := workflowrunactions.NewWorkflowRunActionsClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowRunActionsClient.Client)

	workflowRunsClient := workflowruns.NewWorkflowRunsClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowRunsClient.Client)

	workflowTriggerHistoriesClient := workflowtriggerhistories.NewWorkflowTriggerHistoriesClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowTriggerHistoriesClient.Client)

	workflowTriggersClient := workflowtriggers.NewWorkflowTriggersClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowTriggersClient.Client)

	workflowVersionsClient := workflowversions.NewWorkflowVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowVersionsClient.Client)

	workflowsClient := workflows.NewWorkflowsClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowsClient.Client)

	return Client{
		IntegrationAccountAgreements:               &integrationAccountAgreementsClient,
		IntegrationAccountAssemblies:               &integrationAccountAssembliesClient,
		IntegrationAccountBatchConfigurations:      &integrationAccountBatchConfigurationsClient,
		IntegrationAccountCertificates:             &integrationAccountCertificatesClient,
		IntegrationAccountMaps:                     &integrationAccountMapsClient,
		IntegrationAccountPartners:                 &integrationAccountPartnersClient,
		IntegrationAccountSchemas:                  &integrationAccountSchemasClient,
		IntegrationAccountSessions:                 &integrationAccountSessionsClient,
		IntegrationAccounts:                        &integrationAccountsClient,
		IntegrationServiceEnvironmentManagedApi:    &integrationServiceEnvironmentManagedApiClient,
		IntegrationServiceEnvironmentManagedApis:   &integrationServiceEnvironmentManagedApisClient,
		IntegrationServiceEnvironmentNetworkHealth: &integrationServiceEnvironmentNetworkHealthClient,
		IntegrationServiceEnvironmentRestart:       &integrationServiceEnvironmentRestartClient,
		IntegrationServiceEnvironmentSkus:          &integrationServiceEnvironmentSkusClient,
		IntegrationServiceEnvironments:             &integrationServiceEnvironmentsClient,
		WorkflowRunActions:                         &workflowRunActionsClient,
		WorkflowRuns:                               &workflowRunsClient,
		WorkflowTriggerHistories:                   &workflowTriggerHistoriesClient,
		WorkflowTriggers:                           &workflowTriggersClient,
		WorkflowVersions:                           &workflowVersionsClient,
		Workflows:                                  &workflowsClient,
	}
}
