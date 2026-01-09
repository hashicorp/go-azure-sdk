package sitelistitemdocumentsetversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemDocumentSetVersionClient struct {
	Client *msgraph.Client
}

func NewSiteListItemDocumentSetVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemDocumentSetVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemdocumentsetversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemDocumentSetVersionClient: %+v", err)
	}

	return &SiteListItemDocumentSetVersionClient{
		Client: client,
	}, nil
}
