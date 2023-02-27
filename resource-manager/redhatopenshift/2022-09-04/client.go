// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_09_04

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
