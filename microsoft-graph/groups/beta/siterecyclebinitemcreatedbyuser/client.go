package siterecyclebinitemcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinItemCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinItemCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinItemCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebinitemcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinItemCreatedByUserClient: %+v", err)
	}

	return &SiteRecycleBinItemCreatedByUserClient{
		Client: client,
	}, nil
}
