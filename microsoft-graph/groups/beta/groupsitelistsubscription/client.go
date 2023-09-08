package groupsitelistsubscription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListSubscriptionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListSubscriptionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListSubscriptionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistsubscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListSubscriptionClient: %+v", err)
	}

	return &GroupSiteListSubscriptionClient{
		Client: client,
	}, nil
}
