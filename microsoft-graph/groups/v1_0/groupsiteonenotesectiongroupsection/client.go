package groupsiteonenotesectiongroupsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionGroupSectionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionGroupSectionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionGroupSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectiongroupsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionGroupSectionClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionGroupSectionClient{
		Client: client,
	}, nil
}
