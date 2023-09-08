package meonenotenotebooksectiongroupsectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionGroupSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionGroupSectionPageContentClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionGroupSectionPageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectiongroupsectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionGroupSectionPageContentClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionGroupSectionPageContentClient{
		Client: client,
	}, nil
}
