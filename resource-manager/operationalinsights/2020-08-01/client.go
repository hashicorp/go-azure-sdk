package v2020_08_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/availableservicetiers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/dataexport"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/datasources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/deletedworkspaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/intelligencepacks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/linkedservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/linkedstorageaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/savedsearches"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/storageinsights"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/tables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2020-08-01/workspaces"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	AvailableServiceTiers *availableservicetiers.AvailableServiceTiersClient
	Clusters              *clusters.ClustersClient
	DataExport            *dataexport.DataExportClient
	DataSources           *datasources.DataSourcesClient
	DeletedWorkspaces     *deletedworkspaces.DeletedWorkspacesClient
	IntelligencePacks     *intelligencepacks.IntelligencePacksClient
	LinkedServices        *linkedservices.LinkedServicesClient
	LinkedStorageAccounts *linkedstorageaccounts.LinkedStorageAccountsClient
	SavedSearches         *savedsearches.SavedSearchesClient
	StorageInsights       *storageinsights.StorageInsightsClient
	Tables                *tables.TablesClient
	Workspaces            *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	availableServiceTiersClient := availableservicetiers.NewAvailableServiceTiersClientWithBaseURI(endpoint)
	configureAuthFunc(&availableServiceTiersClient.Client)

	clustersClient := clusters.NewClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&clustersClient.Client)

	dataExportClient := dataexport.NewDataExportClientWithBaseURI(endpoint)
	configureAuthFunc(&dataExportClient.Client)

	dataSourcesClient := datasources.NewDataSourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&dataSourcesClient.Client)

	deletedWorkspacesClient := deletedworkspaces.NewDeletedWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&deletedWorkspacesClient.Client)

	intelligencePacksClient := intelligencepacks.NewIntelligencePacksClientWithBaseURI(endpoint)
	configureAuthFunc(&intelligencePacksClient.Client)

	linkedServicesClient := linkedservices.NewLinkedServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&linkedServicesClient.Client)

	linkedStorageAccountsClient := linkedstorageaccounts.NewLinkedStorageAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&linkedStorageAccountsClient.Client)

	savedSearchesClient := savedsearches.NewSavedSearchesClientWithBaseURI(endpoint)
	configureAuthFunc(&savedSearchesClient.Client)

	storageInsightsClient := storageinsights.NewStorageInsightsClientWithBaseURI(endpoint)
	configureAuthFunc(&storageInsightsClient.Client)

	tablesClient := tables.NewTablesClientWithBaseURI(endpoint)
	configureAuthFunc(&tablesClient.Client)

	workspacesClient := workspaces.NewWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacesClient.Client)

	return Client{
		AvailableServiceTiers: &availableServiceTiersClient,
		Clusters:              &clustersClient,
		DataExport:            &dataExportClient,
		DataSources:           &dataSourcesClient,
		DeletedWorkspaces:     &deletedWorkspacesClient,
		IntelligencePacks:     &intelligencePacksClient,
		LinkedServices:        &linkedServicesClient,
		LinkedStorageAccounts: &linkedStorageAccountsClient,
		SavedSearches:         &savedSearchesClient,
		StorageInsights:       &storageInsightsClient,
		Tables:                &tablesClient,
		Workspaces:            &workspacesClient,
	}
}
