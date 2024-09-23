package rolemanagementalertalertalertincident

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementAlertAlertAlertIncidentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementAlertAlertAlertIncidentClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementAlertAlertAlertIncidentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementalertalertalertincident", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementAlertAlertAlertIncidentClient: %+v", err)
	}

	return &RoleManagementAlertAlertAlertIncidentClient{
		Client: client,
	}, nil
}
