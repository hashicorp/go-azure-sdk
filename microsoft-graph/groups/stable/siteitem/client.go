package siteitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteItemClient struct {
	Client *msgraph.Client
}

func NewSiteItemClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteItemClient: %+v", err)
	}

	return &SiteItemClient{
		Client: client,
	}, nil
}
