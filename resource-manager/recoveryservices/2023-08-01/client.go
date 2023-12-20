package v2023_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-08-01/getprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-08-01/listprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-08-01/recoveryservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-08-01/registeredidentities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-08-01/replicationusages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-08-01/vaultcertificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-08-01/vaultextendedinfo"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-08-01/vaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2023-08-01/vaultusages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	GetPrivateLinkResources  *getprivatelinkresources.GetPrivateLinkResourcesClient
	ListPrivateLinkResources *listprivatelinkresources.ListPrivateLinkResourcesClient
	RecoveryServices         *recoveryservices.RecoveryServicesClient
	RegisteredIdentities     *registeredidentities.RegisteredIdentitiesClient
	ReplicationUsages        *replicationusages.ReplicationUsagesClient
	VaultCertificates        *vaultcertificates.VaultCertificatesClient
	VaultExtendedInfo        *vaultextendedinfo.VaultExtendedInfoClient
	VaultUsages              *vaultusages.VaultUsagesClient
	Vaults                   *vaults.VaultsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	getPrivateLinkResourcesClient, err := getprivatelinkresources.NewGetPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GetPrivateLinkResources client: %+v", err)
	}
	configureFunc(getPrivateLinkResourcesClient.Client)

	listPrivateLinkResourcesClient, err := listprivatelinkresources.NewListPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ListPrivateLinkResources client: %+v", err)
	}
	configureFunc(listPrivateLinkResourcesClient.Client)

	recoveryServicesClient, err := recoveryservices.NewRecoveryServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecoveryServices client: %+v", err)
	}
	configureFunc(recoveryServicesClient.Client)

	registeredIdentitiesClient, err := registeredidentities.NewRegisteredIdentitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RegisteredIdentities client: %+v", err)
	}
	configureFunc(registeredIdentitiesClient.Client)

	replicationUsagesClient, err := replicationusages.NewReplicationUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationUsages client: %+v", err)
	}
	configureFunc(replicationUsagesClient.Client)

	vaultCertificatesClient, err := vaultcertificates.NewVaultCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VaultCertificates client: %+v", err)
	}
	configureFunc(vaultCertificatesClient.Client)

	vaultExtendedInfoClient, err := vaultextendedinfo.NewVaultExtendedInfoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VaultExtendedInfo client: %+v", err)
	}
	configureFunc(vaultExtendedInfoClient.Client)

	vaultUsagesClient, err := vaultusages.NewVaultUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VaultUsages client: %+v", err)
	}
	configureFunc(vaultUsagesClient.Client)

	vaultsClient, err := vaults.NewVaultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Vaults client: %+v", err)
	}
	configureFunc(vaultsClient.Client)

	return &Client{
		GetPrivateLinkResources:  getPrivateLinkResourcesClient,
		ListPrivateLinkResources: listPrivateLinkResourcesClient,
		RecoveryServices:         recoveryServicesClient,
		RegisteredIdentities:     registeredIdentitiesClient,
		ReplicationUsages:        replicationUsagesClient,
		VaultCertificates:        vaultCertificatesClient,
		VaultExtendedInfo:        vaultExtendedInfoClient,
		VaultUsages:              vaultUsagesClient,
		Vaults:                   vaultsClient,
	}, nil
}
