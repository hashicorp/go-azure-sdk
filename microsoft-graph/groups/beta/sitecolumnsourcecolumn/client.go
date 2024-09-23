package sitecolumnsourcecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteColumnSourceColumnClient struct {
	Client *msgraph.Client
}

func NewSiteColumnSourceColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteColumnSourceColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitecolumnsourcecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteColumnSourceColumnClient: %+v", err)
	}

	return &SiteColumnSourceColumnClient{
		Client: client,
	}, nil
}
