package groupsiteonenotenotebooksectiongroupsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionGroupSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionGroupSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionGroupSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectiongroupsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionGroupSectionGroupClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionGroupSectionGroupClient{
		Client: client,
	}, nil
}
