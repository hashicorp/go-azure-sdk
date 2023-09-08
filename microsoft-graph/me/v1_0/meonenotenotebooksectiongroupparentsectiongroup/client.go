package meonenotenotebooksectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
