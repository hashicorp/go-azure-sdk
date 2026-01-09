package sitepagesitepagecreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageSitePageCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewSitePageSitePageCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageSitePageCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagesitepagecreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageSitePageCreatedByUserClient: %+v", err)
	}

	return &SitePageSitePageCreatedByUserClient{
		Client: client,
	}, nil
}
