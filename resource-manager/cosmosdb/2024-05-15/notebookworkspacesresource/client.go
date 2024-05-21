package notebookworkspacesresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookWorkspacesResourceClient struct {
	Client *resourcemanager.Client
}

func NewNotebookWorkspacesResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*NotebookWorkspacesResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "notebookworkspacesresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NotebookWorkspacesResourceClient: %+v", err)
	}

	return &NotebookWorkspacesResourceClient{
		Client: client,
	}, nil
}
