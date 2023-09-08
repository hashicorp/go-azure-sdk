package groupsiteonenotesectiongroupparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionGroupParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionGroupParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionGroupParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectiongroupparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionGroupParentNotebookClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionGroupParentNotebookClient{
		Client: client,
	}, nil
}
