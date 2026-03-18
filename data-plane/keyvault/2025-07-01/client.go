package v2025_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/data-plane/keyvault/2025-07-01/administrations"
	"github.com/hashicorp/go-azure-sdk/data-plane/keyvault/2025-07-01/certificates"
	"github.com/hashicorp/go-azure-sdk/data-plane/keyvault/2025-07-01/deletedstorage"
	"github.com/hashicorp/go-azure-sdk/data-plane/keyvault/2025-07-01/keys"
	"github.com/hashicorp/go-azure-sdk/data-plane/keyvault/2025-07-01/secrets"
	"github.com/hashicorp/go-azure-sdk/data-plane/keyvault/2025-07-01/securitydomains"
	"github.com/hashicorp/go-azure-sdk/data-plane/keyvault/2025-07-01/storage"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

type Client struct {
	Administrations *administrations.AdministrationsClient
	Certificates    *certificates.CertificatesClient
	DeletedStorage  *deletedstorage.DeletedStorageClient
	Keys            *keys.KeysClient
	Secrets         *secrets.SecretsClient
	Securitydomains *securitydomains.SecuritydomainsClient
	Storage         *storage.StorageClient
}

func NewClient(configureFunc func(c *dataplane.Client)) (*Client, error) {
	administrationsClient, err := administrations.NewAdministrationsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Administrations client: %+v", err)
	}
	configureFunc(administrationsClient.Client)

	certificatesClient, err := certificates.NewCertificatesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Certificates client: %+v", err)
	}
	configureFunc(certificatesClient.Client)

	deletedStorageClient, err := deletedstorage.NewDeletedStorageClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building DeletedStorage client: %+v", err)
	}
	configureFunc(deletedStorageClient.Client)

	keysClient, err := keys.NewKeysClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Keys client: %+v", err)
	}
	configureFunc(keysClient.Client)

	secretsClient, err := secrets.NewSecretsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Secrets client: %+v", err)
	}
	configureFunc(secretsClient.Client)

	securitydomainsClient, err := securitydomains.NewSecuritydomainsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Securitydomains client: %+v", err)
	}
	configureFunc(securitydomainsClient.Client)

	storageClient, err := storage.NewStorageClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Storage client: %+v", err)
	}
	configureFunc(storageClient.Client)

	return &Client{
		Administrations: administrationsClient,
		Certificates:    certificatesClient,
		DeletedStorage:  deletedStorageClient,
		Keys:            keysClient,
		Secrets:         secretsClient,
		Securitydomains: securitydomainsClient,
		Storage:         storageClient,
	}, nil
}
