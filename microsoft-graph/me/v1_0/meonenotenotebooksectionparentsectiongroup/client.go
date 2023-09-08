package meonenotenotebooksectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionParentSectionGroupClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
