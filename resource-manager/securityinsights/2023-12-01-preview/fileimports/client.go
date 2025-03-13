package fileimports

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileImportsClient struct {
	Client *resourcemanager.Client
}

func NewFileImportsClientWithBaseURI(sdkApi sdkEnv.Api) (*FileImportsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "fileimports", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FileImportsClient: %+v", err)
	}

	return &FileImportsClient{
		Client: client,
	}, nil
}
