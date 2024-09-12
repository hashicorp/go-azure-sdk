package sitepage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageClient struct {
	Client *msgraph.Client
}

func NewSitePageClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageClient: %+v", err)
	}

	return &SitePageClient{
		Client: client,
	}, nil
}
