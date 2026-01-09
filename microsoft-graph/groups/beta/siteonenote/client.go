package siteonenote

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenote", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteClient: %+v", err)
	}

	return &SiteOnenoteClient{
		Client: client,
	}, nil
}
