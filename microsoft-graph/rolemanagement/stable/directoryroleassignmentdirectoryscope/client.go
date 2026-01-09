package directoryroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &DirectoryRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}
