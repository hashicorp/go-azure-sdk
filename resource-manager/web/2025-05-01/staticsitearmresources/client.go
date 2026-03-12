package staticsitearmresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSiteARMResourcesClient struct {
	Client *resourcemanager.Client
}

func NewStaticSiteARMResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*StaticSiteARMResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "staticsitearmresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StaticSiteARMResourcesClient: %+v", err)
	}

	return &StaticSiteARMResourcesClient{
		Client: client,
	}, nil
}
