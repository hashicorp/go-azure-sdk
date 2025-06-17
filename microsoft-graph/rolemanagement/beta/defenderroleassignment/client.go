package defenderroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewDefenderRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defenderroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderRoleAssignmentClient: %+v", err)
	}

	return &DefenderRoleAssignmentClient{
		Client: client,
	}, nil
}
