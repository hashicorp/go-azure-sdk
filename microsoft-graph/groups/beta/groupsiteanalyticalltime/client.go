package groupsiteanalyticalltime

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteAnalyticAllTimeClient struct {
	Client *msgraph.Client
}

func NewGroupSiteAnalyticAllTimeClientWithBaseURI(api sdkEnv.Api) (*GroupSiteAnalyticAllTimeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteanalyticalltime", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteAnalyticAllTimeClient: %+v", err)
	}

	return &GroupSiteAnalyticAllTimeClient{
		Client: client,
	}, nil
}
