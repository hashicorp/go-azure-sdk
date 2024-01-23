package v2023_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/keys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/managedhsmkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/managedhsms"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/mhsmlistprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/mhsmlistregions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/mhsmprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/mhsmprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/secrets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-07-01/vaults"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Keys                               *keys.KeysClient
	MHSMListPrivateEndpointConnections *mhsmlistprivateendpointconnections.MHSMListPrivateEndpointConnectionsClient
	MHSMListRegions                    *mhsmlistregions.MHSMListRegionsClient
	MHSMPrivateEndpointConnections     *mhsmprivateendpointconnections.MHSMPrivateEndpointConnectionsClient
	MHSMPrivateLinkResources           *mhsmprivatelinkresources.MHSMPrivateLinkResourcesClient
	ManagedHsmKeys                     *managedhsmkeys.ManagedHsmKeysClient
	ManagedHsms                        *managedhsms.ManagedHsmsClient
	PrivateEndpointConnections         *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources               *privatelinkresources.PrivateLinkResourcesClient
	Secrets                            *secrets.SecretsClient
	Vaults                             *vaults.VaultsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	keysClient, err := keys.NewKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Keys client: %+v", err)
	}
	configureFunc(keysClient.Client)

	mHSMListPrivateEndpointConnectionsClient, err := mhsmlistprivateendpointconnections.NewMHSMListPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MHSMListPrivateEndpointConnections client: %+v", err)
	}
	configureFunc(mHSMListPrivateEndpointConnectionsClient.Client)

	mHSMListRegionsClient, err := mhsmlistregions.NewMHSMListRegionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MHSMListRegions client: %+v", err)
	}
	configureFunc(mHSMListRegionsClient.Client)

	mHSMPrivateEndpointConnectionsClient, err := mhsmprivateendpointconnections.NewMHSMPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MHSMPrivateEndpointConnections client: %+v", err)
	}
	configureFunc(mHSMPrivateEndpointConnectionsClient.Client)

	mHSMPrivateLinkResourcesClient, err := mhsmprivatelinkresources.NewMHSMPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MHSMPrivateLinkResources client: %+v", err)
	}
	configureFunc(mHSMPrivateLinkResourcesClient.Client)

	managedHsmKeysClient, err := managedhsmkeys.NewManagedHsmKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedHsmKeys client: %+v", err)
	}
	configureFunc(managedHsmKeysClient.Client)

	managedHsmsClient, err := managedhsms.NewManagedHsmsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedHsms client: %+v", err)
	}
	configureFunc(managedHsmsClient.Client)

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

	secretsClient, err := secrets.NewSecretsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Secrets client: %+v", err)
	}
	configureFunc(secretsClient.Client)

	vaultsClient, err := vaults.NewVaultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Vaults client: %+v", err)
	}
	configureFunc(vaultsClient.Client)

	return &Client{
		Keys:                               keysClient,
		MHSMListPrivateEndpointConnections: mHSMListPrivateEndpointConnectionsClient,
		MHSMListRegions:                    mHSMListRegionsClient,
		MHSMPrivateEndpointConnections:     mHSMPrivateEndpointConnectionsClient,
		MHSMPrivateLinkResources:           mHSMPrivateLinkResourcesClient,
		ManagedHsmKeys:                     managedHsmKeysClient,
		ManagedHsms:                        managedHsmsClient,
		PrivateEndpointConnections:         privateEndpointConnectionsClient,
		PrivateLinkResources:               privateLinkResourcesClient,
		Secrets:                            secretsClient,
		Vaults:                             vaultsClient,
	}, nil
}
