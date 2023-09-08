package groupsiteonenotesectiongroupsectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionGroupSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionGroupSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionGroupSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectiongroupsectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionGroupSectionParentSectionGroupClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionGroupSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
