package sshpublickeys

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/client"
	"github.com/hashicorp/go-azure-sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SshPublicKeysClient struct {
	Client  autorest.Client
	Client2 *client.ResourceManagerClient
	baseUri string
}

func NewSshPublicKeysClientWithBaseURI(endpoint string) SshPublicKeysClient {
	return SshPublicKeysClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		Client2: client.NewResourceManagerClient(environments.ApiEndpoint(endpoint), ""),
		baseUri: endpoint,
	}
}
