package groupsiteonenotenotebooksectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionPageClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionPageClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionPageClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionPageClient{
		Client: client,
	}, nil
}
