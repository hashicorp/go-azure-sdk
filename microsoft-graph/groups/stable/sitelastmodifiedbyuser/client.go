package sitelastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewSiteLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteLastModifiedByUserClient: %+v", err)
	}

	return &SiteLastModifiedByUserClient{
		Client: client,
	}, nil
}
