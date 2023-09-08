package groupsiteonenotenotebooksection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionClient{
		Client: client,
	}, nil
}
