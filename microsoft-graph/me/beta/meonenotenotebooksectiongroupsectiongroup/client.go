package meonenotenotebooksectiongroupsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionGroupSectionGroupClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionGroupSectionGroupClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionGroupSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectiongroupsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionGroupSectionGroupClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionGroupSectionGroupClient{
		Client: client,
	}, nil
}
