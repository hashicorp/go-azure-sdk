package serviceprincipalappmanagementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalAppManagementPolicyClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalAppManagementPolicyClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalAppManagementPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalappmanagementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalAppManagementPolicyClient: %+v", err)
	}

	return &ServicePrincipalAppManagementPolicyClient{
		Client: client,
	}, nil
}
