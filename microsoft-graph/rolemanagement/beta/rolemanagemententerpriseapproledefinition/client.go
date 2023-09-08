package rolemanagemententerpriseapproledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleDefinitionClient{
		Client: client,
	}, nil
}
