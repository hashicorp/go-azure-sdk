package usergroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserGroupClient struct {
	Client *resourcemanager.Client
}

func NewUserGroupClientWithBaseURI(api environments.Api) (*UserGroupClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "usergroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserGroupClient: %+v", err)
	}

	return &UserGroupClient{
		Client: client,
	}, nil
}
