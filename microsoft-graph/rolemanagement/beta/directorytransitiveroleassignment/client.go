package directorytransitiveroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryTransitiveRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewDirectoryTransitiveRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryTransitiveRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directorytransitiveroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryTransitiveRoleAssignmentClient: %+v", err)
	}

	return &DirectoryTransitiveRoleAssignmentClient{
		Client: client,
	}, nil
}
