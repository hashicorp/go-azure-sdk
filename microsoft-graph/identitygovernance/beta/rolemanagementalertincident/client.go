package rolemanagementalertincident

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementAlertIncidentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementAlertIncidentClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementAlertIncidentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementalertincident", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementAlertIncidentClient: %+v", err)
	}

	return &RoleManagementAlertIncidentClient{
		Client: client,
	}, nil
}
