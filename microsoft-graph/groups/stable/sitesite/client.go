package sitesite

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteSiteClient struct {
	Client *msgraph.Client
}

func NewSiteSiteClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteSiteClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitesite", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteSiteClient: %+v", err)
	}

	return &SiteSiteClient{
		Client: client,
	}, nil
}
