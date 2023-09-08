package groupsiteonenotesectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionPageParentSectionClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionPageParentSectionClient{
		Client: client,
	}, nil
}
