package v2025_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-08-01/deletedvaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-08-01/getoperationresult"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-08-01/openapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-08-01/privatelinkresourceoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-08-01/vaultextendedinforesources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-08-01/vaults"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DeletedVaults                     *deletedvaults.DeletedVaultsClient
	GetOperationResult                *getoperationresult.GetOperationResultClient
	Openapis                          *openapis.OpenapisClient
	PrivateLinkResourceOperationGroup *privatelinkresourceoperationgroup.PrivateLinkResourceOperationGroupClient
	VaultExtendedInfoResources        *vaultextendedinforesources.VaultExtendedInfoResourcesClient
	Vaults                            *vaults.VaultsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	deletedVaultsClient, err := deletedvaults.NewDeletedVaultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedVaults client: %+v", err)
	}
	configureFunc(deletedVaultsClient.Client)

	getOperationResultClient, err := getoperationresult.NewGetOperationResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GetOperationResult client: %+v", err)
	}
	configureFunc(getOperationResultClient.Client)

	openapisClient, err := openapis.NewOpenapisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Openapis client: %+v", err)
	}
	configureFunc(openapisClient.Client)

	privateLinkResourceOperationGroupClient, err := privatelinkresourceoperationgroup.NewPrivateLinkResourceOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResourceOperationGroup client: %+v", err)
	}
	configureFunc(privateLinkResourceOperationGroupClient.Client)

	vaultExtendedInfoResourcesClient, err := vaultextendedinforesources.NewVaultExtendedInfoResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VaultExtendedInfoResources client: %+v", err)
	}
	configureFunc(vaultExtendedInfoResourcesClient.Client)

	vaultsClient, err := vaults.NewVaultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Vaults client: %+v", err)
	}
	configureFunc(vaultsClient.Client)

	return &Client{
		DeletedVaults:                     deletedVaultsClient,
		GetOperationResult:                getOperationResultClient,
		Openapis:                          openapisClient,
		PrivateLinkResourceOperationGroup: privateLinkResourceOperationGroupClient,
		VaultExtendedInfoResources:        vaultExtendedInfoResourcesClient,
		Vaults:                            vaultsClient,
	}, nil
}
