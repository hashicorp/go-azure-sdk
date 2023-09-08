package grouponenotenotebooksectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
