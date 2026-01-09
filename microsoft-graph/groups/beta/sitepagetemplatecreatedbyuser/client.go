package sitepagetemplatecreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatecreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateCreatedByUserClient: %+v", err)
	}

	return &SitePageTemplateCreatedByUserClient{
		Client: client,
	}, nil
}
