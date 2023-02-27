// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_09_10

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/getprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/listprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/recoveryservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/registeredidentities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/replicationusages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/vaultcertificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/vaultextendedinfo"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/vaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/vaultusages"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	getPrivateLinkResourcesClient := getprivatelinkresources.NewGetPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&getPrivateLinkResourcesClient.Client)

	listPrivateLinkResourcesClient := listprivatelinkresources.NewListPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&listPrivateLinkResourcesClient.Client)

	recoveryServicesClient := recoveryservices.NewRecoveryServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&recoveryServicesClient.Client)

	registeredIdentitiesClient := registeredidentities.NewRegisteredIdentitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&registeredIdentitiesClient.Client)

	replicationUsagesClient := replicationusages.NewReplicationUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationUsagesClient.Client)

	vaultCertificatesClient := vaultcertificates.NewVaultCertificatesClientWithBaseURI(endpoint)
	configureAuthFunc(&vaultCertificatesClient.Client)

	vaultExtendedInfoClient := vaultextendedinfo.NewVaultExtendedInfoClientWithBaseURI(endpoint)
	configureAuthFunc(&vaultExtendedInfoClient.Client)

	vaultUsagesClient := vaultusages.NewVaultUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&vaultUsagesClient.Client)

	vaultsClient := vaults.NewVaultsClientWithBaseURI(endpoint)
	configureAuthFunc(&vaultsClient.Client)

	return Client{
		GetPrivateLinkResources:  &getPrivateLinkResourcesClient,
		ListPrivateLinkResources: &listPrivateLinkResourcesClient,
		RecoveryServices:         &recoveryServicesClient,
		RegisteredIdentities:     &registeredIdentitiesClient,
		ReplicationUsages:        &replicationUsagesClient,
		VaultCertificates:        &vaultCertificatesClient,
		VaultExtendedInfo:        &vaultExtendedInfoClient,
		VaultUsages:              &vaultUsagesClient,
		Vaults:                   &vaultsClient,
	}
}
