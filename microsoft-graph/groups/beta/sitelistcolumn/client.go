package sitelistcolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListColumnClient struct {
	Client *msgraph.Client
}

func NewSiteListColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListColumnClient: %+v", err)
	}

	return &SiteListColumnClient{
		Client: client,
	}, nil
}
