package groupsiteonenotenotebooksectiongroupparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionGroupParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionGroupParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionGroupParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectiongroupparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionGroupParentNotebookClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionGroupParentNotebookClient{
		Client: client,
	}, nil
}
