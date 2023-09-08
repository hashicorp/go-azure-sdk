package groupsiteonenotenotebooksectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionPageParentNotebookClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
