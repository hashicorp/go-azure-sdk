package sitelistpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListPermissionClient struct {
	Client *msgraph.Client
}

func NewSiteListPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListPermissionClient: %+v", err)
	}

	return &SiteListPermissionClient{
		Client: client,
	}, nil
}
