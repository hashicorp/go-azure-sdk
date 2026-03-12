package siteextensioninfooperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteExtensionInfoOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewSiteExtensionInfoOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteExtensionInfoOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "siteextensioninfooperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteExtensionInfoOperationGroupClient: %+v", err)
	}

	return &SiteExtensionInfoOperationGroupClient{
		Client: client,
	}, nil
}
