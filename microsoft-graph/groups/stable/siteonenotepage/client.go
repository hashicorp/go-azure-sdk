package siteonenotepage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenotePageClient struct {
	Client *msgraph.Client
}

func NewSiteOnenotePageClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenotePageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotepage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenotePageClient: %+v", err)
	}

	return &SiteOnenotePageClient{
		Client: client,
	}, nil
}
