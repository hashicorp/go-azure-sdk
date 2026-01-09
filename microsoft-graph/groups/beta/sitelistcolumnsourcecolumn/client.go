package sitelistcolumnsourcecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListColumnSourceColumnClient struct {
	Client *msgraph.Client
}

func NewSiteListColumnSourceColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListColumnSourceColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcolumnsourcecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListColumnSourceColumnClient: %+v", err)
	}

	return &SiteListColumnSourceColumnClient{
		Client: client,
	}, nil
}
