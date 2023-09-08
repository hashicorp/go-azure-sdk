package groupsiteanalyticlastsevenday

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteAnalyticLastSevenDayClient struct {
	Client *msgraph.Client
}

func NewGroupSiteAnalyticLastSevenDayClientWithBaseURI(api sdkEnv.Api) (*GroupSiteAnalyticLastSevenDayClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteanalyticlastsevenday", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteAnalyticLastSevenDayClient: %+v", err)
	}

	return &GroupSiteAnalyticLastSevenDayClient{
		Client: client,
	}, nil
}
