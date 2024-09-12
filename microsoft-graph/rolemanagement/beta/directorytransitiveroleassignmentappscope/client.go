package directorytransitiveroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryTransitiveRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryTransitiveRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directorytransitiveroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryTransitiveRoleAssignmentAppScopeClient: %+v", err)
	}

	return &DirectoryTransitiveRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
