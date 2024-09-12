package sitelist

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListClient struct {
	Client *msgraph.Client
}

func NewSiteListClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelist", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListClient: %+v", err)
	}

	return &SiteListClient{
		Client: client,
	}, nil
}
