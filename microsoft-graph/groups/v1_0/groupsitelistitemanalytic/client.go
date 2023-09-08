package groupsitelistitemanalytic

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemAnalyticClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemAnalyticClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemAnalyticClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemanalytic", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemAnalyticClient: %+v", err)
	}

	return &GroupSiteListItemAnalyticClient{
		Client: client,
	}, nil
}
