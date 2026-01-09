package sitelistitempermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemPermissionClient struct {
	Client *msgraph.Client
}

func NewSiteListItemPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitempermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemPermissionClient: %+v", err)
	}

	return &SiteListItemPermissionClient{
		Client: client,
	}, nil
}
