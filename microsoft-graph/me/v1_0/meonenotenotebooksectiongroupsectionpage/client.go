package meonenotenotebooksectiongroupsectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionGroupSectionPageClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionGroupSectionPageClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionGroupSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectiongroupsectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionGroupSectionPageClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionGroupSectionPageClient{
		Client: client,
	}, nil
}
