package sitepagelastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewSitePageLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagelastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageLastModifiedByUserClient: %+v", err)
	}

	return &SitePageLastModifiedByUserClient{
		Client: client,
	}, nil
}
