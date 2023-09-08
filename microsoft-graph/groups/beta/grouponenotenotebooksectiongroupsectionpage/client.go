package grouponenotenotebooksectiongroupsectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionGroupSectionPageClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionGroupSectionPageClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionGroupSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectiongroupsectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionGroupSectionPageClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionGroupSectionPageClient{
		Client: client,
	}, nil
}
