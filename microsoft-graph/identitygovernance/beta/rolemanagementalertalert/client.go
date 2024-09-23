package rolemanagementalertalert

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementAlertAlertClient struct {
	Client *msgraph.Client
}

func NewRoleManagementAlertAlertClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementAlertAlertClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementalertalert", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementAlertAlertClient: %+v", err)
	}

	return &RoleManagementAlertAlertClient{
		Client: client,
	}, nil
}
