package sitepagesitepagelastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageSitePageLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewSitePageSitePageLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageSitePageLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagesitepagelastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageSitePageLastModifiedByUserClient: %+v", err)
	}

	return &SitePageSitePageLastModifiedByUserClient{
		Client: client,
	}, nil
}
