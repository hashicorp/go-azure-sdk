package meonenotenotebooksectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionPageParentNotebookClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
