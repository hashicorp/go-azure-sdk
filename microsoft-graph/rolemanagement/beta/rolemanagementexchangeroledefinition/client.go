package rolemanagementexchangeroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementExchangeRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementExchangeRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementExchangeRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementexchangeroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementExchangeRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementExchangeRoleDefinitionClient{
		Client: client,
	}, nil
}
