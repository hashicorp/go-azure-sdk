package sitecreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewSiteCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitecreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteCreatedByUserClient: %+v", err)
	}

	return &SiteCreatedByUserClient{
		Client: client,
	}, nil
}
