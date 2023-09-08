package rolemanagemententerpriseappresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppResourceNamespaceResourceActionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseappresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppResourceNamespaceResourceActionClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
