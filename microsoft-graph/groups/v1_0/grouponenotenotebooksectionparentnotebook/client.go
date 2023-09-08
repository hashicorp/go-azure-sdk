package grouponenotenotebooksectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionParentNotebookClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionParentNotebookClient{
		Client: client,
	}, nil
}
