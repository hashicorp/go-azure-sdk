package sitelistcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewSiteListCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListCreatedByUserClient: %+v", err)
	}

	return &SiteListCreatedByUserClient{
		Client: client,
	}, nil
}
