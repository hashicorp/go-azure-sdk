package websiteinstancestatusoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebSiteInstanceStatusOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewWebSiteInstanceStatusOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*WebSiteInstanceStatusOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "websiteinstancestatusoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WebSiteInstanceStatusOperationGroupClient: %+v", err)
	}

	return &WebSiteInstanceStatusOperationGroupClient{
		Client: client,
	}, nil
}
