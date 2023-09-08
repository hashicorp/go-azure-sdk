package groupsitepagecreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSitePageCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewGroupSitePageCreatedByUserClientWithBaseURI(api sdkEnv.Api) (*GroupSitePageCreatedByUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitepagecreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSitePageCreatedByUserClient: %+v", err)
	}

	return &GroupSitePageCreatedByUserClient{
		Client: client,
	}, nil
}
