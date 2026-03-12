package staticsitebasicauthpropertiesarmresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSiteBasicAuthPropertiesARMResourcesClient struct {
	Client *resourcemanager.Client
}

func NewStaticSiteBasicAuthPropertiesARMResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*StaticSiteBasicAuthPropertiesARMResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "staticsitebasicauthpropertiesarmresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StaticSiteBasicAuthPropertiesARMResourcesClient: %+v", err)
	}

	return &StaticSiteBasicAuthPropertiesARMResourcesClient{
		Client: client,
	}, nil
}
