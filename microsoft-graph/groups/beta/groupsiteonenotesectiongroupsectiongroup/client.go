package groupsiteonenotesectiongroupsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionGroupSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionGroupSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionGroupSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectiongroupsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionGroupSectionGroupClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionGroupSectionGroupClient{
		Client: client,
	}, nil
}
