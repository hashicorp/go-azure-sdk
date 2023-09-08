package groupsiteonenotenotebooksectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionParentSectionGroupClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
