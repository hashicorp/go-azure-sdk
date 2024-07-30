package roleassignmentrolescopetag

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentRoleScopeTagClient struct {
	Client *msgraph.Client
}

func NewRoleAssignmentRoleScopeTagClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleAssignmentRoleScopeTagClient, error) {
	client, err := msgraph.NewClient(sdkApi, "roleassignmentrolescopetag", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleAssignmentRoleScopeTagClient: %+v", err)
	}

	return &RoleAssignmentRoleScopeTagClient{
		Client: client,
	}, nil
}
