package rolemanagemententerpriseappresourcenamespaceresourceactionresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppResourceNamespaceResourceActionResourceScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppResourceNamespaceResourceActionResourceScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppResourceNamespaceResourceActionResourceScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseappresourcenamespaceresourceactionresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppResourceNamespaceResourceActionResourceScopeClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppResourceNamespaceResourceActionResourceScopeClient{
		Client: client,
	}, nil
}
