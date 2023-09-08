package groupsiteonenotesectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionParentSectionGroupClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
