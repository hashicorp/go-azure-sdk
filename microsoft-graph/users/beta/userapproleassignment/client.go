package userapproleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAppRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewUserAppRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*UserAppRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userapproleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAppRoleAssignmentClient: %+v", err)
	}

	return &UserAppRoleAssignmentClient{
		Client: client,
	}, nil
}
