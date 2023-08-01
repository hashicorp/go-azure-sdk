package v2018_06_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/databasecolumns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/databaseschemas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/databasesecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/databasetables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/failoverdatabases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/failoverelasticpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/instancepools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/locationcapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/longtermretentionmanagedinstancebackups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/manageddatabasecolumns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/manageddatabaserestoredetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/manageddatabases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/manageddatabaseschemas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/manageddatabasesensitivitylabels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/manageddatabasetables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/managedinstancelongtermretentionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/managedinstanceoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/managedinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/managedinstancevulnerabilityassessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/serverazureadadministrators"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/servervulnerabilityassessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/usages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DatabaseColumns                          *databasecolumns.DatabaseColumnsClient
	DatabaseSchemas                          *databaseschemas.DatabaseSchemasClient
	DatabaseSecurityAlertPolicies            *databasesecurityalertpolicies.DatabaseSecurityAlertPoliciesClient
	DatabaseTables                           *databasetables.DatabaseTablesClient
	FailoverDatabases                        *failoverdatabases.FailoverDatabasesClient
	FailoverElasticPools                     *failoverelasticpools.FailoverElasticPoolsClient
	InstancePools                            *instancepools.InstancePoolsClient
	LocationCapabilities                     *locationcapabilities.LocationCapabilitiesClient
	LongTermRetentionManagedInstanceBackups  *longtermretentionmanagedinstancebackups.LongTermRetentionManagedInstanceBackupsClient
	ManagedDatabaseColumns                   *manageddatabasecolumns.ManagedDatabaseColumnsClient
	ManagedDatabaseRestoreDetails            *manageddatabaserestoredetails.ManagedDatabaseRestoreDetailsClient
	ManagedDatabaseSchemas                   *manageddatabaseschemas.ManagedDatabaseSchemasClient
	ManagedDatabaseSensitivityLabels         *manageddatabasesensitivitylabels.ManagedDatabaseSensitivityLabelsClient
	ManagedDatabaseTables                    *manageddatabasetables.ManagedDatabaseTablesClient
	ManagedDatabases                         *manageddatabases.ManagedDatabasesClient
	ManagedInstanceLongTermRetentionPolicies *managedinstancelongtermretentionpolicies.ManagedInstanceLongTermRetentionPoliciesClient
	ManagedInstanceOperations                *managedinstanceoperations.ManagedInstanceOperationsClient
	ManagedInstanceVulnerabilityAssessments  *managedinstancevulnerabilityassessments.ManagedInstanceVulnerabilityAssessmentsClient
	ManagedInstances                         *managedinstances.ManagedInstancesClient
	PrivateEndpointConnections               *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                     *privatelinkresources.PrivateLinkResourcesClient
	ServerAzureADAdministrators              *serverazureadadministrators.ServerAzureADAdministratorsClient
	ServerVulnerabilityAssessments           *servervulnerabilityassessments.ServerVulnerabilityAssessmentsClient
	Usages                                   *usages.UsagesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	databaseColumnsClient, err := databasecolumns.NewDatabaseColumnsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseColumns client: %+v", err)
	}
	configureFunc(databaseColumnsClient.Client)

	databaseSchemasClient, err := databaseschemas.NewDatabaseSchemasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseSchemas client: %+v", err)
	}
	configureFunc(databaseSchemasClient.Client)

	databaseSecurityAlertPoliciesClient, err := databasesecurityalertpolicies.NewDatabaseSecurityAlertPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(databaseSecurityAlertPoliciesClient.Client)

	databaseTablesClient, err := databasetables.NewDatabaseTablesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseTables client: %+v", err)
	}
	configureFunc(databaseTablesClient.Client)

	failoverDatabasesClient, err := failoverdatabases.NewFailoverDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FailoverDatabases client: %+v", err)
	}
	configureFunc(failoverDatabasesClient.Client)

	failoverElasticPoolsClient, err := failoverelasticpools.NewFailoverElasticPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FailoverElasticPools client: %+v", err)
	}
	configureFunc(failoverElasticPoolsClient.Client)

	instancePoolsClient, err := instancepools.NewInstancePoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InstancePools client: %+v", err)
	}
	configureFunc(instancePoolsClient.Client)

	locationCapabilitiesClient, err := locationcapabilities.NewLocationCapabilitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LocationCapabilities client: %+v", err)
	}
	configureFunc(locationCapabilitiesClient.Client)

	longTermRetentionManagedInstanceBackupsClient, err := longtermretentionmanagedinstancebackups.NewLongTermRetentionManagedInstanceBackupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LongTermRetentionManagedInstanceBackups client: %+v", err)
	}
	configureFunc(longTermRetentionManagedInstanceBackupsClient.Client)

	managedDatabaseColumnsClient, err := manageddatabasecolumns.NewManagedDatabaseColumnsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseColumns client: %+v", err)
	}
	configureFunc(managedDatabaseColumnsClient.Client)

	managedDatabaseRestoreDetailsClient, err := manageddatabaserestoredetails.NewManagedDatabaseRestoreDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseRestoreDetails client: %+v", err)
	}
	configureFunc(managedDatabaseRestoreDetailsClient.Client)

	managedDatabaseSchemasClient, err := manageddatabaseschemas.NewManagedDatabaseSchemasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseSchemas client: %+v", err)
	}
	configureFunc(managedDatabaseSchemasClient.Client)

	managedDatabaseSensitivityLabelsClient, err := manageddatabasesensitivitylabels.NewManagedDatabaseSensitivityLabelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseSensitivityLabels client: %+v", err)
	}
	configureFunc(managedDatabaseSensitivityLabelsClient.Client)

	managedDatabaseTablesClient, err := manageddatabasetables.NewManagedDatabaseTablesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseTables client: %+v", err)
	}
	configureFunc(managedDatabaseTablesClient.Client)

	managedDatabasesClient, err := manageddatabases.NewManagedDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabases client: %+v", err)
	}
	configureFunc(managedDatabasesClient.Client)

	managedInstanceLongTermRetentionPoliciesClient, err := managedinstancelongtermretentionpolicies.NewManagedInstanceLongTermRetentionPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceLongTermRetentionPolicies client: %+v", err)
	}
	configureFunc(managedInstanceLongTermRetentionPoliciesClient.Client)

	managedInstanceOperationsClient, err := managedinstanceoperations.NewManagedInstanceOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceOperations client: %+v", err)
	}
	configureFunc(managedInstanceOperationsClient.Client)

	managedInstanceVulnerabilityAssessmentsClient, err := managedinstancevulnerabilityassessments.NewManagedInstanceVulnerabilityAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceVulnerabilityAssessments client: %+v", err)
	}
	configureFunc(managedInstanceVulnerabilityAssessmentsClient.Client)

	managedInstancesClient, err := managedinstances.NewManagedInstancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstances client: %+v", err)
	}
	configureFunc(managedInstancesClient.Client)

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

	serverAzureADAdministratorsClient, err := serverazureadadministrators.NewServerAzureADAdministratorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerAzureADAdministrators client: %+v", err)
	}
	configureFunc(serverAzureADAdministratorsClient.Client)

	serverVulnerabilityAssessmentsClient, err := servervulnerabilityassessments.NewServerVulnerabilityAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerVulnerabilityAssessments client: %+v", err)
	}
	configureFunc(serverVulnerabilityAssessmentsClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	return &Client{
		DatabaseColumns:                          databaseColumnsClient,
		DatabaseSchemas:                          databaseSchemasClient,
		DatabaseSecurityAlertPolicies:            databaseSecurityAlertPoliciesClient,
		DatabaseTables:                           databaseTablesClient,
		FailoverDatabases:                        failoverDatabasesClient,
		FailoverElasticPools:                     failoverElasticPoolsClient,
		InstancePools:                            instancePoolsClient,
		LocationCapabilities:                     locationCapabilitiesClient,
		LongTermRetentionManagedInstanceBackups:  longTermRetentionManagedInstanceBackupsClient,
		ManagedDatabaseColumns:                   managedDatabaseColumnsClient,
		ManagedDatabaseRestoreDetails:            managedDatabaseRestoreDetailsClient,
		ManagedDatabaseSchemas:                   managedDatabaseSchemasClient,
		ManagedDatabaseSensitivityLabels:         managedDatabaseSensitivityLabelsClient,
		ManagedDatabaseTables:                    managedDatabaseTablesClient,
		ManagedDatabases:                         managedDatabasesClient,
		ManagedInstanceLongTermRetentionPolicies: managedInstanceLongTermRetentionPoliciesClient,
		ManagedInstanceOperations:                managedInstanceOperationsClient,
		ManagedInstanceVulnerabilityAssessments:  managedInstanceVulnerabilityAssessmentsClient,
		ManagedInstances:                         managedInstancesClient,
		PrivateEndpointConnections:               privateEndpointConnectionsClient,
		PrivateLinkResources:                     privateLinkResourcesClient,
		ServerAzureADAdministrators:              serverAzureADAdministratorsClient,
		ServerVulnerabilityAssessments:           serverVulnerabilityAssessmentsClient,
		Usages:                                   usagesClient,
	}, nil
}
