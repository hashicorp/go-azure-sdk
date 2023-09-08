package meonenotenotebooksection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionClient{
		Client: client,
	}, nil
}
