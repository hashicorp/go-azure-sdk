package rolemanagementalertalertconfigurationalertdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementAlertAlertConfigurationAlertDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementAlertAlertConfigurationAlertDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementAlertAlertConfigurationAlertDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementalertalertconfigurationalertdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementAlertAlertConfigurationAlertDefinitionClient: %+v", err)
	}

	return &RoleManagementAlertAlertConfigurationAlertDefinitionClient{
		Client: client,
	}, nil
}
