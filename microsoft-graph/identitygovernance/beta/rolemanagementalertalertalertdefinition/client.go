package rolemanagementalertalertalertdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementAlertAlertAlertDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementAlertAlertAlertDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementAlertAlertAlertDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementalertalertalertdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementAlertAlertAlertDefinitionClient: %+v", err)
	}

	return &RoleManagementAlertAlertAlertDefinitionClient{
		Client: client,
	}, nil
}
