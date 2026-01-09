package siteonenotesectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupClient{
		Client: client,
	}, nil
}
