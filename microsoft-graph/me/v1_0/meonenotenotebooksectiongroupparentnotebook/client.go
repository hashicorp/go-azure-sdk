package meonenotenotebooksectiongroupparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionGroupParentNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionGroupParentNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionGroupParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectiongroupparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionGroupParentNotebookClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionGroupParentNotebookClient{
		Client: client,
	}, nil
}
