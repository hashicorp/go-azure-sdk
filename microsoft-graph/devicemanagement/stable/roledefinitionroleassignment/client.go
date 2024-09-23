package roledefinitionroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleDefinitionRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleDefinitionRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleDefinitionRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "roledefinitionroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleDefinitionRoleAssignmentClient: %+v", err)
	}

	return &RoleDefinitionRoleAssignmentClient{
		Client: client,
	}, nil
}
