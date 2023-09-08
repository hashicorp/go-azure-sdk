package groupsiteonenotesectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionGroupParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
