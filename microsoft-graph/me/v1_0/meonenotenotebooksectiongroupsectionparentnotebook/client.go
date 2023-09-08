package meonenotenotebooksectiongroupsectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionGroupSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionGroupSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionGroupSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectiongroupsectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionGroupSectionParentNotebookClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionGroupSectionParentNotebookClient{
		Client: client,
	}, nil
}
