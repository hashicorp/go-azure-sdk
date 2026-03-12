package sitesourcecontroloperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteSourceControlOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewSiteSourceControlOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteSourceControlOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sitesourcecontroloperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteSourceControlOperationGroupClient: %+v", err)
	}

	return &SiteSourceControlOperationGroupClient{
		Client: client,
	}, nil
}
