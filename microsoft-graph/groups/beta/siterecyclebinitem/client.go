package siterecyclebinitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinItemClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinItemClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebinitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinItemClient: %+v", err)
	}

	return &SiteRecycleBinItemClient{
		Client: client,
	}, nil
}
