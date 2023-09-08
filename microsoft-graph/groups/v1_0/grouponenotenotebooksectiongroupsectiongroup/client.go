package grouponenotenotebooksectiongroupsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionGroupSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionGroupSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionGroupSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectiongroupsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionGroupSectionGroupClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionGroupSectionGroupClient{
		Client: client,
	}, nil
}
