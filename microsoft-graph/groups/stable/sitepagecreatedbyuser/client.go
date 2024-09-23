package sitepagecreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewSitePageCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagecreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageCreatedByUserClient: %+v", err)
	}

	return &SitePageCreatedByUserClient{
		Client: client,
	}, nil
}
