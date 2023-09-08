package groupsitelistcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListCreatedByUserClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListCreatedByUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListCreatedByUserClient: %+v", err)
	}

	return &GroupSiteListCreatedByUserClient{
		Client: client,
	}, nil
}
