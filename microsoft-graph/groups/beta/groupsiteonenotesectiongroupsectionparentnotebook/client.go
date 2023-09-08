package groupsiteonenotesectiongroupsectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionGroupSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionGroupSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectiongroupsectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionGroupSectionParentNotebookClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionGroupSectionParentNotebookClient{
		Client: client,
	}, nil
}
