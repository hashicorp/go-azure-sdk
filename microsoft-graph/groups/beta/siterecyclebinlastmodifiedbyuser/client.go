package siterecyclebinlastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebinlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinLastModifiedByUserClient: %+v", err)
	}

	return &SiteRecycleBinLastModifiedByUserClient{
		Client: client,
	}, nil
}
