package apimconfig

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type APIMConfigClient struct {
	Client *resourcemanager.Client
}

func NewAPIMConfigClientWithBaseURI(sdkApi sdkEnv.Api) (*APIMConfigClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "apimconfig", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating APIMConfigClient: %+v", err)
	}

	return &APIMConfigClient{
		Client: client,
	}, nil
}
