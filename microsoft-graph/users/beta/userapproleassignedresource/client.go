package userapproleassignedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAppRoleAssignedResourceClient struct {
	Client *msgraph.Client
}

func NewUserAppRoleAssignedResourceClientWithBaseURI(api sdkEnv.Api) (*UserAppRoleAssignedResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userapproleassignedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAppRoleAssignedResourceClient: %+v", err)
	}

	return &UserAppRoleAssignedResourceClient{
		Client: client,
	}, nil
}
