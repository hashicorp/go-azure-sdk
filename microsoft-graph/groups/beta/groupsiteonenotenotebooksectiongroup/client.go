package groupsiteonenotenotebooksectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionGroupClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionGroupClient{
		Client: client,
	}, nil
}
