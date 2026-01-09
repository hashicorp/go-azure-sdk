package site

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteClient struct {
	Client *msgraph.Client
}

func NewSiteClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteClient, error) {
	client, err := msgraph.NewClient(sdkApi, "site", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteClient: %+v", err)
	}

	return &SiteClient{
		Client: client,
	}, nil
}
