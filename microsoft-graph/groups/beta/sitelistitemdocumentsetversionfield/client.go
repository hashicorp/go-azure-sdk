package sitelistitemdocumentsetversionfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemDocumentSetVersionFieldClient struct {
	Client *msgraph.Client
}

func NewSiteListItemDocumentSetVersionFieldClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemDocumentSetVersionFieldClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemdocumentsetversionfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemDocumentSetVersionFieldClient: %+v", err)
	}

	return &SiteListItemDocumentSetVersionFieldClient{
		Client: client,
	}, nil
}
