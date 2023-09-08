package groupsiteonenotenotebooksectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
