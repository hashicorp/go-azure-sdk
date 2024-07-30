package sitecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteColumnClient struct {
	Client *msgraph.Client
}

func NewSiteColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteColumnClient: %+v", err)
	}

	return &SiteColumnClient{
		Client: client,
	}, nil
}
