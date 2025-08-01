package sshpublickeyresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SshPublicKeyResourcesClient struct {
	Client *resourcemanager.Client
}

func NewSshPublicKeyResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*SshPublicKeyResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sshpublickeyresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SshPublicKeyResourcesClient: %+v", err)
	}

	return &SshPublicKeyResourcesClient{
		Client: client,
	}, nil
}
