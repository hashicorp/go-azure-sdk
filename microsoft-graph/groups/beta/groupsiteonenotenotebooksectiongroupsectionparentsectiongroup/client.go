package groupsiteonenotenotebooksectiongroupsectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionGroupSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionGroupSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionGroupSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectiongroupsectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionGroupSectionParentSectionGroupClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionGroupSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
