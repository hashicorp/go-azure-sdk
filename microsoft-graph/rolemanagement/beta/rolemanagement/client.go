package rolemanagement

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementClient struct {
	Client *msgraph.Client
}

func NewRoleManagementClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementClient: %+v", err)
	}

	return &RoleManagementClient{
		Client: client,
	}, nil
}
