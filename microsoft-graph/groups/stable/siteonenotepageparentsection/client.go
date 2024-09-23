package siteonenotepageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenotePageParentSectionClient struct {
	Client *msgraph.Client
}

func NewSiteOnenotePageParentSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenotePageParentSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotepageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenotePageParentSectionClient: %+v", err)
	}

	return &SiteOnenotePageParentSectionClient{
		Client: client,
	}, nil
}
