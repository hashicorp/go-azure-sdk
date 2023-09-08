package grouponenotenotebooksectiongroupsectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionGroupSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionGroupSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionGroupSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectiongroupsectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionGroupSectionParentSectionGroupClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionGroupSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
