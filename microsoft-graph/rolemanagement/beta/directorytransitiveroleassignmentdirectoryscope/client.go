package directorytransitiveroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryTransitiveRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryTransitiveRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directorytransitiveroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryTransitiveRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &DirectoryTransitiveRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}
