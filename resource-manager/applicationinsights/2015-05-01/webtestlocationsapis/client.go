package webtestlocationsapis

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebTestLocationsAPIsClient struct {
	Client *resourcemanager.Client
}

func NewWebTestLocationsAPIsClientWithBaseURI(sdkApi sdkEnv.Api) (*WebTestLocationsAPIsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "webtestlocationsapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WebTestLocationsAPIsClient: %+v", err)
	}

	return &WebTestLocationsAPIsClient{
		Client: client,
	}, nil
}
