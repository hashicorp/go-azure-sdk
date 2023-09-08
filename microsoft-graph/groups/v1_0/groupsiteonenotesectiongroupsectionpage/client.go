package groupsiteonenotesectiongroupsectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionGroupSectionPageClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionGroupSectionPageClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionGroupSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectiongroupsectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionGroupSectionPageClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionGroupSectionPageClient{
		Client: client,
	}, nil
}
