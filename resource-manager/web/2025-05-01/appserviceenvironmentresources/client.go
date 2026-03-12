package appserviceenvironmentresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppServiceEnvironmentResourcesClient struct {
	Client *resourcemanager.Client
}

func NewAppServiceEnvironmentResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*AppServiceEnvironmentResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "appserviceenvironmentresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppServiceEnvironmentResourcesClient: %+v", err)
	}

	return &AppServiceEnvironmentResourcesClient{
		Client: client,
	}, nil
}
