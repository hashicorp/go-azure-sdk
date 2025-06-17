package v2025_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/availableservicetiers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/dataexport"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/datasources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/deletedworkspaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/intelligencepacks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/linkedservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/linkedstorageaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/networksecurityperimeterconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/querypackqueries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/querypacks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/savedsearches"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/storageinsights"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/tables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/workspaces"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AvailableServiceTiers                  *availableservicetiers.AvailableServiceTiersClient
	Clusters                               *clusters.ClustersClient
	DataExport                             *dataexport.DataExportClient
	DataSources                            *datasources.DataSourcesClient
	DeletedWorkspaces                      *deletedworkspaces.DeletedWorkspacesClient
	IntelligencePacks                      *intelligencepacks.IntelligencePacksClient
	LinkedServices                         *linkedservices.LinkedServicesClient
	LinkedStorageAccounts                  *linkedstorageaccounts.LinkedStorageAccountsClient
	NetworkSecurityPerimeterConfigurations *networksecurityperimeterconfigurations.NetworkSecurityPerimeterConfigurationsClient
	QueryPackQueries                       *querypackqueries.QueryPackQueriesClient
	QueryPacks                             *querypacks.QueryPacksClient
	SavedSearches                          *savedsearches.SavedSearchesClient
	StorageInsights                        *storageinsights.StorageInsightsClient
	Tables                                 *tables.TablesClient
	Workspaces                             *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	availableServiceTiersClient, err := availableservicetiers.NewAvailableServiceTiersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailableServiceTiers client: %+v", err)
	}
	configureFunc(availableServiceTiersClient.Client)

	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	dataExportClient, err := dataexport.NewDataExportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataExport client: %+v", err)
	}
	configureFunc(dataExportClient.Client)

	dataSourcesClient, err := datasources.NewDataSourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataSources client: %+v", err)
	}
	configureFunc(dataSourcesClient.Client)

	deletedWorkspacesClient, err := deletedworkspaces.NewDeletedWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedWorkspaces client: %+v", err)
	}
	configureFunc(deletedWorkspacesClient.Client)

	intelligencePacksClient, err := intelligencepacks.NewIntelligencePacksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntelligencePacks client: %+v", err)
	}
	configureFunc(intelligencePacksClient.Client)

	linkedServicesClient, err := linkedservices.NewLinkedServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LinkedServices client: %+v", err)
	}
	configureFunc(linkedServicesClient.Client)

	linkedStorageAccountsClient, err := linkedstorageaccounts.NewLinkedStorageAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LinkedStorageAccounts client: %+v", err)
	}
	configureFunc(linkedStorageAccountsClient.Client)

	networkSecurityPerimeterConfigurationsClient, err := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterConfigurations client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterConfigurationsClient.Client)

	queryPackQueriesClient, err := querypackqueries.NewQueryPackQueriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QueryPackQueries client: %+v", err)
	}
	configureFunc(queryPackQueriesClient.Client)

	queryPacksClient, err := querypacks.NewQueryPacksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QueryPacks client: %+v", err)
	}
	configureFunc(queryPacksClient.Client)

	savedSearchesClient, err := savedsearches.NewSavedSearchesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SavedSearches client: %+v", err)
	}
	configureFunc(savedSearchesClient.Client)

	storageInsightsClient, err := storageinsights.NewStorageInsightsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StorageInsights client: %+v", err)
	}
	configureFunc(storageInsightsClient.Client)

	tablesClient, err := tables.NewTablesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Tables client: %+v", err)
	}
	configureFunc(tablesClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		AvailableServiceTiers:                  availableServiceTiersClient,
		Clusters:                               clustersClient,
		DataExport:                             dataExportClient,
		DataSources:                            dataSourcesClient,
		DeletedWorkspaces:                      deletedWorkspacesClient,
		IntelligencePacks:                      intelligencePacksClient,
		LinkedServices:                         linkedServicesClient,
		LinkedStorageAccounts:                  linkedStorageAccountsClient,
		NetworkSecurityPerimeterConfigurations: networkSecurityPerimeterConfigurationsClient,
		QueryPackQueries:                       queryPackQueriesClient,
		QueryPacks:                             queryPacksClient,
		SavedSearches:                          savedSearchesClient,
		StorageInsights:                        storageInsightsClient,
		Tables:                                 tablesClient,
		Workspaces:                             workspacesClient,
	}, nil
}
