package meonenotenotebooksectiongroupsectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionGroupSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionGroupSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionGroupSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectiongroupsectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionGroupSectionPageParentNotebookClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionGroupSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
