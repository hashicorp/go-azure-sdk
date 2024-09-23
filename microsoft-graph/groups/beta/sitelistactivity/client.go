package sitelistactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListActivityClient struct {
	Client *msgraph.Client
}

func NewSiteListActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListActivityClient: %+v", err)
	}

	return &SiteListActivityClient{
		Client: client,
	}, nil
}
