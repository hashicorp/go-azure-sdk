package grouponenotenotebooksectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionPageParentNotebookClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
