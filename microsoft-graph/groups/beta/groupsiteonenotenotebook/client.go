package groupsiteonenotenotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookClient{
		Client: client,
	}, nil
}
