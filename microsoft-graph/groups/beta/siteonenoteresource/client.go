package siteonenoteresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteResourceClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenoteresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteResourceClient: %+v", err)
	}

	return &SiteOnenoteResourceClient{
		Client: client,
	}, nil
}
