package enterpriseapproledefinitioninheritspermissionsfrom

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleDefinitionInheritsPermissionsFromClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleDefinitionInheritsPermissionsFromClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproledefinitioninheritspermissionsfrom", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleDefinitionInheritsPermissionsFromClient: %+v", err)
	}

	return &EnterpriseAppRoleDefinitionInheritsPermissionsFromClient{
		Client: client,
	}, nil
}
