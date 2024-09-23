package rolemanagementalertoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementAlertOperationClient struct {
	Client *msgraph.Client
}

func NewRoleManagementAlertOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementAlertOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementalertoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementAlertOperationClient: %+v", err)
	}

	return &RoleManagementAlertOperationClient{
		Client: client,
	}, nil
}
