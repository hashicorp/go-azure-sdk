package sitelistdrive

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListDriveClient struct {
	Client *msgraph.Client
}

func NewSiteListDriveClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListDriveClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistdrive", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListDriveClient: %+v", err)
	}

	return &SiteListDriveClient{
		Client: client,
	}, nil
}
