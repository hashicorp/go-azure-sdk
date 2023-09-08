package meonenotenotebooksectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionPageClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionPageClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionPageClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionPageClient{
		Client: client,
	}, nil
}
