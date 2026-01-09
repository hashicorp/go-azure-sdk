package rolemanagementalertalertconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementAlertAlertConfigurationClient struct {
	Client *msgraph.Client
}

func NewRoleManagementAlertAlertConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementAlertAlertConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementalertalertconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementAlertAlertConfigurationClient: %+v", err)
	}

	return &RoleManagementAlertAlertConfigurationClient{
		Client: client,
	}, nil
}
