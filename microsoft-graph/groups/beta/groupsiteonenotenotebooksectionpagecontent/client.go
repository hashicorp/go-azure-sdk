package groupsiteonenotenotebooksectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionPageContentClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionPageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionPageContentClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionPageContentClient{
		Client: client,
	}, nil
}
