package grouponenotenotebooksectiongroupsectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionGroupSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionGroupSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionGroupSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectiongroupsectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionGroupSectionPageParentNotebookClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionGroupSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
