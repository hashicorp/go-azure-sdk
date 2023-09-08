package grouponenotenotebooksectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionGroupClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionGroupClient{
		Client: client,
	}, nil
}
