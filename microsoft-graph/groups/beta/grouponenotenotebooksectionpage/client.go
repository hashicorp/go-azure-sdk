package grouponenotenotebooksectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionPageClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionPageClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionPageClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionPageClient{
		Client: client,
	}, nil
}
