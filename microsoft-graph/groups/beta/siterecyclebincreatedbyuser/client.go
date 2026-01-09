package siterecyclebincreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebincreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinCreatedByUserClient: %+v", err)
	}

	return &SiteRecycleBinCreatedByUserClient{
		Client: client,
	}, nil
}
