package roleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "roleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleAssignmentClient: %+v", err)
	}

	return &RoleAssignmentClient{
		Client: client,
	}, nil
}
