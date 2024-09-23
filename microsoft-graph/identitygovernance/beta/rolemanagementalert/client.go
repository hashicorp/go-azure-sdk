package rolemanagementalert

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementAlertClient struct {
	Client *msgraph.Client
}

func NewRoleManagementAlertClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementAlertClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementalert", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementAlertClient: %+v", err)
	}

	return &RoleManagementAlertClient{
		Client: client,
	}, nil
}
