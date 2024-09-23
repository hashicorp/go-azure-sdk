package siteonenoteresourcecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteResourceContentClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteResourceContentClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteResourceContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenoteresourcecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteResourceContentClient: %+v", err)
	}

	return &SiteOnenoteResourceContentClient{
		Client: client,
	}, nil
}
