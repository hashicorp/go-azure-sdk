package staticsitesasyncoperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesAsyncOperationsClient struct {
	Client *resourcemanager.Client
}

func NewStaticSitesAsyncOperationsClientWithBaseURI(sdkApi sdkEnv.Api) (*StaticSitesAsyncOperationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "staticsitesasyncoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StaticSitesAsyncOperationsClient: %+v", err)
	}

	return &StaticSitesAsyncOperationsClient{
		Client: client,
	}, nil
}
