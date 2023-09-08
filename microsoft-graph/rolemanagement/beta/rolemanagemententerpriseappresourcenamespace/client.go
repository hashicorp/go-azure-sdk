package rolemanagemententerpriseappresourcenamespace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppResourceNamespaceClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppResourceNamespaceClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppResourceNamespaceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseappresourcenamespace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppResourceNamespaceClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppResourceNamespaceClient{
		Client: client,
	}, nil
}
