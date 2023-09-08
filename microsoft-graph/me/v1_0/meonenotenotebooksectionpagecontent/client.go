package meonenotenotebooksectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionPageContentClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionPageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionPageContentClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionPageContentClient{
		Client: client,
	}, nil
}
