package serviceprincipalapproleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalAppRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalAppRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalAppRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalapproleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalAppRoleAssignmentClient: %+v", err)
	}

	return &ServicePrincipalAppRoleAssignmentClient{
		Client: client,
	}, nil
}
