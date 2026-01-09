package sitelistoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListOperationClient struct {
	Client *msgraph.Client
}

func NewSiteListOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListOperationClient: %+v", err)
	}

	return &SiteListOperationClient{
		Client: client,
	}, nil
}
