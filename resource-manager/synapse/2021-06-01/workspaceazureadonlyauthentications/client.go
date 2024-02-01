package workspaceazureadonlyauthentications

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceAzureADOnlyAuthenticationsClient struct {
	Client *resourcemanager.Client
}

func NewWorkspaceAzureADOnlyAuthenticationsClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkspaceAzureADOnlyAuthenticationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "workspaceazureadonlyauthentications", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceAzureADOnlyAuthenticationsClient: %+v", err)
	}

	return &WorkspaceAzureADOnlyAuthenticationsClient{
		Client: client,
	}, nil
}
