package grouponenotenotebooksectiongroupsectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionGroupSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionGroupSectionPageContentClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionGroupSectionPageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectiongroupsectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionGroupSectionPageContentClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionGroupSectionPageContentClient{
		Client: client,
	}, nil
}
