package grouponenotenotebooksectiongroupparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionGroupParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionGroupParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionGroupParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectiongroupparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionGroupParentNotebookClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionGroupParentNotebookClient{
		Client: client,
	}, nil
}
