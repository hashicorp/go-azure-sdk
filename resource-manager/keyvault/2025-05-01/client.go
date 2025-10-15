package v2025_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/deletedmanagedhsms"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/deletedvaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/keyoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/keys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/managedhsmkeyoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/managedhsmkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/managedhsms"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/mhsmprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/openapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/secrets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/vaults"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DeletedManagedHsms             *deletedmanagedhsms.DeletedManagedHsmsClient
	DeletedVaults                  *deletedvaults.DeletedVaultsClient
	KeyOperationGroup              *keyoperationgroup.KeyOperationGroupClient
	Keys                           *keys.KeysClient
	ManagedHsmKeyOperationGroup    *managedhsmkeyoperationgroup.ManagedHsmKeyOperationGroupClient
	ManagedHsmKeys                 *managedhsmkeys.ManagedHsmKeysClient
	ManagedHsms                    *managedhsms.ManagedHsmsClient
	MhsmPrivateEndpointConnections *mhsmprivateendpointconnections.MhsmPrivateEndpointConnectionsClient
	Openapis                       *openapis.OpenapisClient
	PrivateEndpointConnections     *privateendpointconnections.PrivateEndpointConnectionsClient
	Secrets                        *secrets.SecretsClient
	Vaults                         *vaults.VaultsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	deletedManagedHsmsClient, err := deletedmanagedhsms.NewDeletedManagedHsmsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedManagedHsms client: %+v", err)
	}
	configureFunc(deletedManagedHsmsClient.Client)

	deletedVaultsClient, err := deletedvaults.NewDeletedVaultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedVaults client: %+v", err)
	}
	configureFunc(deletedVaultsClient.Client)

	keyOperationGroupClient, err := keyoperationgroup.NewKeyOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building KeyOperationGroup client: %+v", err)
	}
	configureFunc(keyOperationGroupClient.Client)

	keysClient, err := keys.NewKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Keys client: %+v", err)
	}
	configureFunc(keysClient.Client)

	managedHsmKeyOperationGroupClient, err := managedhsmkeyoperationgroup.NewManagedHsmKeyOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedHsmKeyOperationGroup client: %+v", err)
	}
	configureFunc(managedHsmKeyOperationGroupClient.Client)

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

	mhsmPrivateEndpointConnectionsClient, err := mhsmprivateendpointconnections.NewMhsmPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MhsmPrivateEndpointConnections client: %+v", err)
	}
	configureFunc(mhsmPrivateEndpointConnectionsClient.Client)

	openapisClient, err := openapis.NewOpenapisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Openapis client: %+v", err)
	}
	configureFunc(openapisClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

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
		DeletedManagedHsms:             deletedManagedHsmsClient,
		DeletedVaults:                  deletedVaultsClient,
		KeyOperationGroup:              keyOperationGroupClient,
		Keys:                           keysClient,
		ManagedHsmKeyOperationGroup:    managedHsmKeyOperationGroupClient,
		ManagedHsmKeys:                 managedHsmKeysClient,
		ManagedHsms:                    managedHsmsClient,
		MhsmPrivateEndpointConnections: mhsmPrivateEndpointConnectionsClient,
		Openapis:                       openapisClient,
		PrivateEndpointConnections:     privateEndpointConnectionsClient,
		Secrets:                        secretsClient,
		Vaults:                         vaultsClient,
	}, nil
}
