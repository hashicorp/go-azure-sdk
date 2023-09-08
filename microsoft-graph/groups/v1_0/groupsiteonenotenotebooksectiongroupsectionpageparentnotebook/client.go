package groupsiteonenotenotebooksectiongroupsectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectiongroupsectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
