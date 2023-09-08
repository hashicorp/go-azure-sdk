package serviceprincipalapproleassignedto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalAppRoleAssignedToClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalAppRoleAssignedToClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalAppRoleAssignedToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalapproleassignedto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalAppRoleAssignedToClient: %+v", err)
	}

	return &ServicePrincipalAppRoleAssignedToClient{
		Client: client,
	}, nil
}
