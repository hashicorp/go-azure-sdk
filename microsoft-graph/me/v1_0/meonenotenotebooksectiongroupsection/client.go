package meonenotenotebooksectiongroupsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionGroupSectionClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionGroupSectionClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionGroupSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectiongroupsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionGroupSectionClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionGroupSectionClient{
		Client: client,
	}, nil
}
