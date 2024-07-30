package sitelistitemlastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewSiteListItemLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemLastModifiedByUserClient: %+v", err)
	}

	return &SiteListItemLastModifiedByUserClient{
		Client: client,
	}, nil
}
