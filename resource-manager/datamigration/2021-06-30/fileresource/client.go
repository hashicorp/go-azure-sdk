package fileresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileResourceClient struct {
	Client *resourcemanager.Client
}

func NewFileResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*FileResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "fileresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FileResourceClient: %+v", err)
	}

	return &FileResourceClient{
		Client: client,
	}, nil
}
