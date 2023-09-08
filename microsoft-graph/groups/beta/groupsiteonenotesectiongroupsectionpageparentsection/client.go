package groupsiteonenotesectiongroupsectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionGroupSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionGroupSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionGroupSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectiongroupsectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionGroupSectionPageParentSectionClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionGroupSectionPageParentSectionClient{
		Client: client,
	}, nil
}
