package libraries

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LibrariesClient struct {
	Client *resourcemanager.Client
}

func NewLibrariesClientWithBaseURI(sdkApi sdkEnv.Api) (*LibrariesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "libraries", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LibrariesClient: %+v", err)
	}

	return &LibrariesClient{
		Client: client,
	}, nil
}
