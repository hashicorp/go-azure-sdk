package grouponenotenotebooksectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionParentSectionGroupClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
