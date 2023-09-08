package groupsiteanalyticitemactivitystatactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteAnalyticItemActivityStatActivityClient struct {
	Client *msgraph.Client
}

func NewGroupSiteAnalyticItemActivityStatActivityClientWithBaseURI(api sdkEnv.Api) (*GroupSiteAnalyticItemActivityStatActivityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteanalyticitemactivitystatactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteAnalyticItemActivityStatActivityClient: %+v", err)
	}

	return &GroupSiteAnalyticItemActivityStatActivityClient{
		Client: client,
	}, nil
}
