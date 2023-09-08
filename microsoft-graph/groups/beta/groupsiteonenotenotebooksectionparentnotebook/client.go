package groupsiteonenotenotebooksectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionParentNotebookClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionParentNotebookClient{
		Client: client,
	}, nil
}
