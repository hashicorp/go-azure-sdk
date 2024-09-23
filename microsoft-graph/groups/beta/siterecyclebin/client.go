package siterecyclebin

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebin", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinClient: %+v", err)
	}

	return &SiteRecycleBinClient{
		Client: client,
	}, nil
}
