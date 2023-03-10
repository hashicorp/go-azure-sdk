package v2022_09_04

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-09-04/machinepools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-09-04/openshiftclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-09-04/openshiftversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-09-04/secrets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-09-04/syncidentityproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-09-04/syncsets"
)

type Client struct {
	MachinePools          *machinepools.MachinePoolsClient
	OpenShiftClusters     *openshiftclusters.OpenShiftClustersClient
	OpenShiftVersions     *openshiftversions.OpenShiftVersionsClient
	Secrets               *secrets.SecretsClient
	SyncIdentityProviders *syncidentityproviders.SyncIdentityProvidersClient
	SyncSets              *syncsets.SyncSetsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	machinePoolsClient := machinepools.NewMachinePoolsClientWithBaseURI(endpoint)
	configureAuthFunc(&machinePoolsClient.Client)

	openShiftClustersClient := openshiftclusters.NewOpenShiftClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&openShiftClustersClient.Client)

	openShiftVersionsClient := openshiftversions.NewOpenShiftVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&openShiftVersionsClient.Client)

	secretsClient := secrets.NewSecretsClientWithBaseURI(endpoint)
	configureAuthFunc(&secretsClient.Client)

	syncIdentityProvidersClient := syncidentityproviders.NewSyncIdentityProvidersClientWithBaseURI(endpoint)
	configureAuthFunc(&syncIdentityProvidersClient.Client)

	syncSetsClient := syncsets.NewSyncSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&syncSetsClient.Client)

	return Client{
		MachinePools:          &machinePoolsClient,
		OpenShiftClusters:     &openShiftClustersClient,
		OpenShiftVersions:     &openShiftVersionsClient,
		Secrets:               &secretsClient,
		SyncIdentityProviders: &syncIdentityProvidersClient,
		SyncSets:              &syncSetsClient,
	}
}
