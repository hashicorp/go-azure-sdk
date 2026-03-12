package staticsitelinkedbackendarmresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSiteLinkedBackendARMResourcesClient struct {
	Client *resourcemanager.Client
}

func NewStaticSiteLinkedBackendARMResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*StaticSiteLinkedBackendARMResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "staticsitelinkedbackendarmresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StaticSiteLinkedBackendARMResourcesClient: %+v", err)
	}

	return &StaticSiteLinkedBackendARMResourcesClient{
		Client: client,
	}, nil
}
