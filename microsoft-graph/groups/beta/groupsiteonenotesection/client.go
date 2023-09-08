package groupsiteonenotesection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionClient{
		Client: client,
	}, nil
}
